package main

import (
	"errors"
	"fmt"
	"os"
	"start/lib"
)

func main() {

	fmt.Println("printbox", lib.VERSION, "looking for ports...")
	if len(os.Args) > 1 {
		if os.Args[1] == "-d" {
			lib.Debug = true
			lib.Debugf("debug enabled")
		}
	}

	// make sure there is a shared directory
	if _, err := os.Stat(lib.SharedPath); errors.Is(err, os.ErrNotExist) {
		fmt.Println("no shared dir", lib.SharedPath)
		os.Exit(1)
	}

	board, err := lib.GetBoard()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := lib.CheckPorts(board); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if board.USBCount == 0 {
		fmt.Println("no ports connected")
		os.Exit(1)
	}
	fmt.Println(board.USBCount, "port(s) found")

	if err := lib.BuildComposeFile(board); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := lib.CheckDirs(board); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
