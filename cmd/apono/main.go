package main

import (
	"fmt"
)

var (
	commit  string
	date    string
	version string
)

func main() {
	fmt.Println("Welcome to Apono CLI")
	fmt.Println("Commit:", commit)
	fmt.Println("Build Date:", date)
	fmt.Println("Version:", version)
}
