package app

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var path = "/home/ali/scan_and_patch/configs/cves/"

func ReadCVEFile(cveFile string) string {

	var serverPackages []string

	myfile, err := os.Open(path + cveFile)
	if err != nil {
		log.Fatalf("failed to open myfile: %v", err)
	}
	scanner := bufio.NewScanner(myfile)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		serverPackages = append(serverPackages, scanner.Text())
	}
	myfile.Close()
	return strings.Join(serverPackages, " ")
}
