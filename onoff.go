package main

import (
	"fmt"
	"os"
)

func main() {
	var err error

	if len(os.Args) < 3 {
		usage()
	}

	LoadConfig()

	switch os.Args[1] {
	case "start":
		err = start(os.Args[2])
		break
	case "stop":
		err = stop(os.Args[2])
		break
	default:
		usage()
	}

	if err != nil {
		fmt.Println(err)
	}
}

func start(instance string) error {
	var err error

	if instance == "all" {
		for key := range ConfigFile.Instance {
			if err = powerOn(key); err != nil {
				return err
			}
		}
	} else {
		if err = powerOn(instance); err != nil {
			return err
		}
	}

	return nil
}

func stop(instance string) error {
	var err error

	if instance == "all" {
		for key := range ConfigFile.Instance {
			if err = powerOff(key); err != nil {
				return err
			}
		}
	} else {
		if err = powerOff(instance); err != nil {
			return err
		}
	}

	return nil
}

func usage() {
	fmt.Println("Usage: onoff <start/stop> <instance/all>")
	os.Exit(1)
}
