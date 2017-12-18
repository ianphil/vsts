package logger

import "fmt"
import "os"

// Stdout is an exported function
func Stdout(msg string) {
	fmt.Println(msg)
}

func Stderr(msg string) {
	fmt.Println("ERROR: " + msg)
	os.Exit(1)
}
