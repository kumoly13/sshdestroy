# sshdestroy

This SSH Brute Force Tool is designed for security professionals to assess and reinforce SSH authentication security. It automates brute force attacks to identify weak credentials and potential vulnerabilities in SSH servers, ensuring robust security measures.
Features

    Multi-threaded Attacks: Conducts high-speed, multi-threaded brute force attacks for efficient password cracking.
    Customizable Parameters: Allows customization of username lists, password dictionaries, and target configurations.
    Real-time Logging: Provides real-time logging and reporting of successful login attempts.
    Compatibility: Supports various SSH server configurations with ease.
    Bypass Mechanisms: Includes built-in mechanisms to bypass basic detection measures.

Installation

    git clone https://github.com/yourusername/sshdestroy.git

  


Navigate to the project directory:


    cd sshdestroy

Build the project:


    go build 

Usage

    Setup: Prepare the target SSH server details and create files for usernames and passwords.
    Run the tool: Execute the tool with the necessary arguments:
    Monitor progress: The tool will log the progress and report any successful login attempts in real-time.

Configuration

    Hosts File: A file containing the target IP addresses and ports in the format IP:PORT. If no port is specified, the default port 22 is used.
    Usernames File: A file containing the list of usernames to be tested.
    Passwords File: A file containing the list of passwords to be tested.
