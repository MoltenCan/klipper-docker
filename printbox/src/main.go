package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"printbox/internal/printbox"
)

func main() {

	fmt.Println("printbox", printbox.VERSION, "looking for ports...")

	var debug bool
	var outfile string
	flag.BoolVar(&debug, "d", false, "enable debug")
	flag.StringVar(&outfile, "o", "docker-compose.yml", "output file")
	flag.Parse()

	if debug {
		printbox.Debug = true
		printbox.Debugf("debug enabled")
	}

	// // make sure there is a shared directory
	// if _, err := os.Stat(printbox.SharedPath); errors.Is(err, os.ErrNotExist) {
	// 	fmt.Println("no shared dir", printbox.SharedPath)
	// 	os.Exit(1)
	// }

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

	cData, err := printbox.BuildComposeFile(board)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	printbox.Logf("writing compose to %s", outfile)
	if err := ioutil.WriteFile(outfile, cData, 0644); err != nil {
		printbox.Logf(err.Error())
	}

	// if err := printbox.CheckDirs(board); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	os.Exit(0)
}
