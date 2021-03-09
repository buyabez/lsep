package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
)

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func epParams(r io.Reader) error {
	scanner := bufio.NewScanner(bufio.NewReader(r))
	for scanner.Scan() {
		u, err := url.Parse(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		for p, v := range u.Query() {
			fmt.Printf("\n%s %s  %s=%s", u.Host, u.Path, p, v[0])
		}

	}

	return nil
}

func main() {
	// https://dev.to/napicella/linux-pipes-in-golang-2e8j
	// Read command line argumets
	var endpoint string
	flag.StringVar(&endpoint, "endpoint", "/", "Print paremters only for this endpoint")
	flag.Parse()

	if isInputFromPipe() {
		epParams(os.Stdin)
	} else {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: cat urls.txt | lsep")
		return
	}
}
