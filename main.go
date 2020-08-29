package main

import (
	"fmt"
	"os"
)

func main() {
	displayIntro()

	for {
		displayMenu()
		input := getUserInput()

		switch input {
		case monitore:
			initMonitoring()
		case Logs:
			printLogs()
		case quit:
			fmt.Println("leaving...")
			os.Exit(0)
		default:
			fmt.Println("invalid input. Try again...")
		}
	}
}
