package main

import (
	"errors"
	"fmt"
	"os"
	"printbox/internal/printbox"
)

func main() {

	fmt.Println("printbox", printbox.VERSION, "looking for ports...")
	if len(os.Args) > 1 {
		if os.Args[1] == "-d" {
			printbox.Debug = true
			printbox.Debugf("debug enabled")
		}
	}

	// make sure there is a shared directory
	if _, err := os.Stat(printbox.SharedPath); errors.Is(err, os.ErrNotExist) {
		fmt.Println("no shared dir", printbox.SharedPath)
		os.Exit(1)
	}

	board, err := printbox.GetBoard()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := printbox.CheckPorts(board); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if board.USBCount == 0 {
		fmt.Println("no ports connected")
		os.Exit(1)
	}
	fmt.Println(board.USBCount, "port(s) found")

	if err := printbox.BuildComposeFile(board); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := printbox.CheckDirs(board); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
