package app

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadServersFile() []string {

	// read the list of servers
	fmt.Println("loading configuration")
	config, err := LoadConfiguration("./configs/config.json")
	if err != nil {
		log.Fatalf("Error loading configuration %v", err)
	}

	file, err := os.Open(config.HostsPath)

	if err != nil {
		log.Fatalf("failed to open hostscentos.txt: %v", err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var serverIP []string

	for scanner.Scan() {
		serverIP = append(serverIP, scanner.Text())
	}
	file.Close()
	return serverIP
}
