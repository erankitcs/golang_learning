package main

import (
	"log"
	"os"
)

type messageType int

const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
	FATAL
)

/*func main() {
	writeLog(INFO, "this is an information message!")
	writeLog(WARNING, "this is a warning")
	writeLog(ERROR, "this is an error")
	writeLog(FATAL, "we crashed")
	writeLog(INFO, "you'll never see this message!")

}*/

func writeLog(messagetype messageType, message string) {

	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	switch messagetype {
	case INFO:
		logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case WARNING:
		logger := log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case ERROR:
		logger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case FATAL:
		logger := log.New(file, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Fatal(message)
	}
}
