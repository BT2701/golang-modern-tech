package logger

import (
	"log"
	"os"
)

// Error function
func Error(message string) {
	log.SetOutput(os.Stdout)
	log.Println(message)
}

// Info function
func Info(message string) {
	log.SetOutput(os.Stdout)
	log.Println(message)
}

// Debug function
func Debug(message string) {
	log.SetOutput(os.Stdout)
	log.Println(message)
}

// Warn function
func Warn(message string) {
	log.SetOutput(os.Stdout)
	log.Println(message)
}

// Fatal function
func Fatal(message string) {
	log.SetOutput(os.Stdout)
	log.Fatal(message)
}

// Panic function
func Panic(message string) {
	log.SetOutput(os.Stdout)
	log.Panic(message)
}

// Errorf function
func Errorf(format string, v ...interface{}) {
	log.SetOutput(os.Stdout)
	log.Printf(format, v...)
}

// Infof function
func Infof(format string, v ...interface{}) {
	log.SetOutput(os.Stdout)
	log.Printf(format, v...)
}
