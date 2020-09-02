package main

import (
	"fmt"
	"net"
)

func main() {
	/*
		Code that is needed to be tested on kubernetes goes here.
	*/

	// Example:
	// This code will find the CNAME of `kubernetes.default` service

	addr, err := net.LookupCNAME("kubernetes.default")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(addr)
}
