package main // import "ssh-client"

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type HostInfo struct {
	Name           string
	Description    string
	Address        string
	Port           int
	Username       string
	Password       string
	PrivateKeyFile string
	PrivateKeyText string
}

var (
	hostsFile   = flag.String("f", "", "host data file (required)")
	hostFileKey = flag.String("k", "", "host data file key which is base64 encoded (required)")
	hostsIDX    = flag.Int("i", 0, "index of host data (required)")

	hosts []HostInfo
	host  HostInfo
	key   []byte
)

func main() {
	var err error

	flag.Parse()
	if flag.NArg() > 0 || *hostsFile == "" || *hostFileKey == "" || *hostsIDX == 0 {
		binaryName := filepath.Base(os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", binaryName)
		flag.PrintDefaults()
		return
	}

	// key = []byte("0123456789!#$%^&*()abcdefghijklm")
	// key, err = generateKey(*hostFileKey)
	// if err != nil {
	// 	fmt.Println("error key generation")
	// 	return
	// }
	key, _ = base64.URLEncoding.DecodeString(*hostFileKey)

	err = loadHostData(*hostsFile, key, &hosts)
	if err != nil {
		fmt.Println("error loading host data file")
		return
	}

	if *hostsIDX > len(hosts) {
		fmt.Printf("index not exist. max index is %d\n", len(hosts))
		return
	}

	*hostsIDX--
	host = hosts[*hostsIDX]
	fmt.Printf("Connecting %s/%s\n", host.Name, host.Address)

	err = openSession()
	if err != nil {
		fmt.Println(err)
	}
}
