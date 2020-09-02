package main

import (
	"fmt"
	"kmodules.xyz/client-go/meta"
)

func main() {
	/*
		Code that is needed to be executed on kubernetes cluster goes here.
	*/

	// Example:
	// This code will find the cluster domain

	fmt.Println(meta.ClusterDomain())
}
