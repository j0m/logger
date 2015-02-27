package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

var logLevel int = 0

func SetLogLevel(l int) {
	logLevel = l
}

// SetOutput sets the output destination for the standard logger.
func SetOutput(w io.Writer) {
	log.SetOutput(w)
}

// Flags returns the output flags for the standard logger.
func Flags() int {
	return log.Flags()
}

// SetFlags sets the output flags for the standard logger.
func SetFlags(flag int) {
	log.SetFlags(flag)
}

// Prefix returns the output prefix for the standard logger.
func Prefix() string {
	return log.Prefix()
}

// SetPrefix sets the output prefix for the standard logger.
func SetPrefix(prefix string) {
	log.SetPrefix(prefix)
}

// These functions write to the standard logger.

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	if logLevel > 0 {
		log.Print(v)
	}
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	if logLevel > 0 {
		log.Printf(format, v)
	}
}

// Println calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
	if logLevel > 0 {
		log.Println(v)
	}
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	if logLevel > 0 {
		log.Fatal(v)
	}
	
	os.Exit(1)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	if logLevel > 0 {
		log.Fatal(format, v)
	}

	os.Exit(1)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(v ...interface{}) {
	if logLevel > 0 {
		log.Fatalln(v)
	}

	os.Exit(1)
}

// Panic is equivalent to Print() followed by a call to panic().
func Panic(v ...interface{}) {
	if logLevel > 0 {
		log.Panic(v)
	}

	s := fmt.Sprintln(v...)
	panic(s)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
	if logLevel > 0 {
		log.Panicf(format, v)
	}

	s := fmt.Sprintln(v...)
	panic(s)
}

// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(v ...interface{}) {
	if logLevel > 0 {
		log.Panicln(v)
	}

	s := fmt.Sprintln(v...)
	panic(s)
}
