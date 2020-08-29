package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// idiomatic enum
const (
	quit uint8 = iota
	monitore
	Logs
)

const monitoraments = 3
const delay = 2

func displayIntro() {
	var version float32 = 0.3
	fmt.Println("app version:", version)
}

func displayMenu() {
	fmt.Println("\n1 -> monitorate")
	fmt.Println("2 -> read logs")
	fmt.Println("0 -> quit")
	fmt.Printf(": ")
}

func getUserInput() uint8 {
	var userImput uint8
	fmt.Scan(&userImput)
	return userImput
}

func initMonitoring() {
	pages := getPages()

	for i := 0; i < monitoraments; i++ {
		for _, page := range pages {
			testPage(page)
		}
		time.Sleep(delay * time.Second)
		fmt.Println()
	}
}

func testPage(page string) {
	response, err := http.Get(page)

	if err != nil {
		fmt.Println("at function 'testPage()', error:", err)
	}

	const successStatusCode = 200
	siteIsUp := true
	if response.StatusCode == successStatusCode {
		fmt.Println(page, "sucessful loaded")
	} else {
		fmt.Println(page, "has problems. Status Code:", response.StatusCode)
		siteIsUp = false
	}
	logMonitorament(page, siteIsUp)
}

func getPages() []string {
	var pages []string
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("at function getPages, error:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		pages = append(pages, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()
	return pages
}

func logMonitorament(page string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("at func logMonitorament(), error:", err)
	}
	file.WriteString(time.Now().Format("15/01/2006 15:04:05") + " - " + page + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func printLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("at func printLogs(), error:", err)
	}

	fmt.Println(string(file))
}
