package logger

import "fmt"

// Stdout is an exported function
func Stdout(msg string) {
	fmt.Println(msg)
}

func Stderr(msg string) {
	fmt.Println("ERROR: " + msg)
}
