package main

import "6251/localservice/cmd"

// @title LocalService back-end server
// @version 0.0.1
// @description
func main() {
	defer cmd.Clean()
	cmd.Start()
}
