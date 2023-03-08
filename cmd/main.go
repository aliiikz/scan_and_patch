package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/aliiikz/scan_and_patch_servers/pkg/app"

	"golang.org/x/crypto/ssh"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func run() error {
	fmt.Println("loading configuration")
	config, err := app.LoadConfiguration("./configs/config.json")
	if err != nil {
		log.Fatalf("Error loading configuration %v", err)
	}

	privateKey, err := os.ReadFile(config.SSHFile)
	if err != nil {
		return fmt.Errorf("read private key: %w", err)
	}

	signer, err := ssh.ParsePrivateKey(privateKey)
	if err != nil {
		return fmt.Errorf("parse private key: %w", err)
	}

	sshConfig := &ssh.ClientConfig{
		User:            "htd",
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// read list of servers based on distribution
	servers := app.ReadServersFile()

	// loop through servers and perform ssh command
	for _, serverCentOSIP := range servers {
		// fmt.Print(serverCentOSIP)
		sshConn, err := ssh.Dial("tcp", serverCentOSIP, sshConfig)
		if err != nil {
			return fmt.Errorf("ssh dial: %w", err)
		}
		defer sshConn.Close()

		session, err := sshConn.NewSession()
		if err != nil {
			return fmt.Errorf("new ssh session: %w", err)
		}
		defer session.Close()

		sessionStdOut, err := session.StdoutPipe()
		if err != nil {
			panic(err)
		}
		go io.Copy(os.Stdout, sessionStdOut)
		sessionStderr, err := session.StderrPipe()
		if err != nil {
			panic(err)
		}
		go io.Copy(os.Stderr, sessionStderr)

		// run command on each server
		fmt.Printf("\n==========Now Working on Server %s==========\n", serverCentOSIP)
		splitcvefile := strings.Split(serverCentOSIP, ":")
		fileName := splitcvefile[0] + "_parsed.txt"

		packages := app.ReadCVEFile(fileName)

		if strings.TrimSpace(packages) == "" {
			fmt.Println("No packages needs updating")

		} else {
			cmd := "sudo yum update -y " + packages
			fmt.Println(cmd)
			if err := session.Run(cmd); err != nil {
				panic(err)
			}
		}
	}
	return err
}
