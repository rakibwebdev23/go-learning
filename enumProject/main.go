package main

import "fmt"

type LogLevel int


const (	
	LogLevelTrace LogLevel = iota
	LogLevelDebug
	LogLevelInfo
	LogLevelWarning
	LogLevelError
)

var logLevelNames = []string{
	"Trace",
	"Debug",
	"Info",
	"Warning",
	"Error",
}

func (l LogLevel) String() string {
	if l < LogLevelTrace || l > LogLevelError {
		return "Unknown"
	}
	return logLevelNames[l]
}

func main() {
	LogLevel	:= LogLevelInfo
	fmt.Printf("Current log level: %s\n", LogLevel)

	// Example of using the log level
	switch LogLevel {
	case LogLevelTrace:
		fmt.Println("This is a trace message.")
	case LogLevelDebug:
		fmt.Println("This is a debug message.")
	case LogLevelInfo:
		fmt.Println("This is an info message.")
	case LogLevelWarning:
		fmt.Println("This is a warning message.")
	case LogLevelError:
		fmt.Println("This is an error message.")
		default:
			fmt.Println("Unknown log level.")									
		}
	}