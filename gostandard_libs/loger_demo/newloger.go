package main

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	FatalLogger   *log.Logger
)

func init() {

	file, err := os.OpenFile("newlog.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.LUTC|log.Lmicroseconds|log.Llongfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	FatalLogger = log.New(file, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

	InfoLogger.Println("This is info")
	WarningLogger.Println("This is warning.")
	ErrorLogger.Println("This is error")
	FatalLogger.Println("This is fetal error.")
}
