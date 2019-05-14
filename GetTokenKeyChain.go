package main

import (
	"log"
	"os/exec"
)

func getTokenKeyChain(login string, domain string) (string, error) {
	//    # I use mac keychain for storing pass  add or remove pass to keychain
	//    # security add-generic-password -a ${USER} -s cmdb -w SECRET_PASS
	//    # security delete-generic-password -a ${USER} -s cmdb
	path, err := exec.LookPath("security")
	if err != nil {
		log.Fatal("Cant find 'security' executable file")
	}
	out, err := exec.Command(path, "find-generic-password", "-a", login, "-s", domain, "-w").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out), nil
}

func main() {
	s, _ := getTokenKeyChain("a_artemyev", "cmdb")
	print(s)
}
