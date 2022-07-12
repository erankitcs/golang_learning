package main

import (
	"alog"
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	out := flag.String("out", "stdout", "File name to use for log output. If stdout is provided, then output is written directly to the console.")
	async := flag.Bool("async", false, "This flag determines if the logger should write asynchronously.")
	flag.Parse()

	var w io.Writer
	var err error
	if strings.ToLower(*out) == "stdout" {
		w = os.Stdout
	} else {
		w, err = os.Create(*out)
		if err != nil {
			log.Fatal("Unable to open log file", err)
		}
	}
	l := alog.New(w)
	go l.Start()

	messageChan := l.MessageChannel()
	errChan := l.ErrorChannel()

	if errChan != nil {
		go func(errChan <-chan error) {
			err := <-errChan
			l.Stop()
			log.Fatalf("Error received from logger: %v\n", err)
		}(errChan)
	}

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Please enter message to write to log or 'q' to quit.")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Unable to read input from command line, please try again.", err)
			continue
		}

		if strings.ToLower(input) == "q\n" || strings.ToLower(input) == "q\r\n" {
			if wc, ok := w.(io.Closer); ok {
				err := wc.Close()
				if err != nil {
					fmt.Println("Failed to close log file:", err)
				}
			}
			l.Stop()
			break
		}
		if *async {
			if messageChan != nil {
				messageChan <- input
			}
		} else {
			_, err = l.Write(input)
			if err != nil {
				fmt.Println("Unable to write message out to log")
			}
		}
	}
}
