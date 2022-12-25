package main

import (
	"io"
	"log"
	"os"
	"time"
)

func main() {
	// defaultLog()
	// supportedFlags()
	// modifyFormat()
	// addLogLevelWithSetPrefix()
	// addLogLevelWithNew()
	// writeToLogfile()
	writeToConsoleAndLogfile()
}

func defaultLog() {
	log.Println("Hello, world!")
}

const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

func supportedFlags() {
	// Log to the console
	log.Println("Default lof message")

	// Set Date to the log: use log.SetFlag()
	log.SetFlags(Ldate)
	// Log to the console with date prepended
	log.Println("Log to the console with date prepended")

	// Set Time to the log
	log.SetFlags(Ltime)
	// Log to the console with time prepended
	log.Println("Log to the console with time prepended")

	// Set Date and Time to the log
	log.SetFlags(Ldate | Ltime)
	// Log to the console with date and time prepended
	log.Println("Log to the console with date and time prepended")

	// Set filename to the log
	log.SetFlags(Lshortfile)
	// Log to the console with filename prepended
	log.Println("Log to the console with filename prepended")

	// Set date, time and filename to the log
	log.SetFlags(LstdFlags | Lshortfile)
	// Log to the console with date, time and filename prepended
	log.Println("Log to the console with date, time and filename prepended")
}

const (
	// YYYY-MM-DD: 2022-03-23
	YYYYMMDD = "2006-01-02"
	// 24h hh:mm:ss: 14:23:20
	HHMMSS24h = "15:04:05"
	// 12h hh:mm:ss: 2:23:20 PM
	HHMMSS12h = "3:04:05 PM"
	// text date: March 23, 2022
	TextDate = "January 2, 2006"
	// text date with weekday: Wednesday, March 23, 2022
	TextDateWithWeekday = "Monday, January 2, 2006"
	// abbreviated text date: Mar 23 Wed
	AbbrTextDate = "Jan 2 Mon"
)

func modifyFormat() {
	// Add flags: prepend filename
	log.SetFlags(Lshortfile)
	// Add custome date
	log.SetPrefix(time.Now().UTC().Format(YYYYMMDD) + ": ")
	log.Println("Log message with date formatted")
	// Add custom time
	log.SetPrefix(time.Now().UTC().Format(HHMMSS12h) + ": ")
	log.Println("Log message with time formatted")
	// Add custome date and time
	log.SetPrefix(time.Now().UTC().Format(YYYYMMDD+" "+HHMMSS12h) + ": ")
	log.Println("Log message with date and time formatted")
}

func addLogLevelWithSetPrefix() {
	// Define flags to be used
	flags := Lshortfile
	// Log without log level
	datetime := time.Now().UTC().Format(YYYYMMDD + " " + HHMMSS12h + ": ")
	log.Println("Without log levels")
	// Set log level SetPrefix
	log.SetFlags(flags)
	// INFO log level
	log.SetPrefix("INFO: " + datetime)
	log.Println("INFO level")
	// WARNING log level
	log.SetPrefix("WARN: " + datetime)
	log.Println("WARN level")
	// DEBUG log level
	log.SetPrefix("DEBUG: " + datetime)
	log.Println("DEBUG level")
	// ERROR log level
	log.SetPrefix("ERROR: " + datetime)
	log.Println("ERROR level")
	// FATAL log level
	log.SetPrefix("FATAL: " + datetime)
	log.Println("FATAL level")
}

func addLogLevelWithNew() {
	// Defines flags
	flags := Lshortfile
	// Create a new logger
	logger := log.New(os.Stdout, "", flags)
	// Define date and time format
	datetime := time.Now().UTC().Format(YYYYMMDD+" "+HHMMSS12h) + ": "
	// Set INFO prefix
	logger.SetPrefix("INFO: " + datetime)
	logger.Println("INFO level")
	// Set WARNING prefix
	logger.SetPrefix("WARN: " + datetime)
	logger.Println("WARN level")
	// Set ERROR prefix
	logger.SetPrefix("ERROR: " + datetime)
	logger.Println("ERROR level")
	// Set DEBUG prefix
	logger.SetPrefix("DEBUG: " + datetime)
	logger.Println("DEBUG level")
	// Set FATAL prefix
	logger.SetPrefix("FATAL: " + datetime)
	logger.Println("FATAL level")
}

func writeToLogfile() {
	// Define flags
	flags := Lshortfile
	// Define date and time format
	datetime := time.Now().UTC().Format(YYYYMMDD+" "+HHMMSS12h) + ": "

	// Open a file if it exist or create one if file does not exist
	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// Set file as the destination to log messages with log level "FATAL"
		logger := log.New(file, "", flags)
		logger.SetPrefix("FATAL: " + datetime)
		// Log a fatal message into the log file
		logger.Println(err)
	}
	defer file.Close()

	logger := log.New(file, "", flags)
	logger.SetPrefix("INFO: " + datetime)
	// Log a fatal message into the log file
	logger.Println("Log custom date and time in a logs.log file")
}

func writeToConsoleAndLogfile() {
	// Define falgs
	flags := Lshortfile
	// Define date and time format
	datetime := time.Now().UTC().Format(YYYYMMDD+" "+HHMMSS12h) + ": "

	// Open a file if it exist or create one if file does not exist
	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// Set file as the destination to log messages with log level "FATAL"
		logger := log.New(file, "", flags)
		logger.SetPrefix("FATAL: " + datetime)
		// Log a fatal message into the log file
		logger.Println(err)
	}
	defer file.Close()

	logger := log.New(file, "", flags)
	logger.SetPrefix("INFO: " + datetime)
	// Log a fatal message into the log file
	mw := io.MultiWriter(os.Stdout, file)
	logger.SetOutput(mw)
	logger.Println("Log to console and file")
}
