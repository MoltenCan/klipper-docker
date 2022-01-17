package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"printbox/internal/printbox"
	"printbox/internal/rpi"
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

	// are we on an RPI?
	board, err := rpi.Identify()
	if err == nil {
		fmt.Println("detected", board.Model)
	}

	ports, err := printbox.GetPorts(board)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(ports) == 0 {
		fmt.Println("no ports found")
		os.Exit(1)
	}

	fmt.Println("ports found:")
	fmt.Printf(" %-20s %s\n", "alias", "port")
	for a, p := range ports {
		fmt.Printf(" %-20s %s\n", a, p)
	}
	fmt.Println("")

	cData, err := printbox.BuildComposeFile(ports)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	printbox.Logf("writing compose to %s", outfile)
	err = ioutil.WriteFile(outfile, cData, 0644)
	if err != nil {
		fmt.Println("failed to write compose:", err)
		os.Exit(1)
	}

	// if err := printbox.CheckDirs(board); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	os.Exit(0)
}
