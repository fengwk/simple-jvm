package main

import "github.com/fengwk/simple-jvm/cmd"

func main() {
	cmd := cmd.Parse()
	cmd.Execute()
}
