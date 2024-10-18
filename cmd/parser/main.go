package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	outputFlag := flag.String("output", "stdout", "Output destination: stdout, json, or html")
	logDirFlag := flag.String("logdir", "./logs", "Directory containing the nginx logs to parse")
	persistFlag := flag.String("persist", "", "Persist parsed logs to DB file (SQLite)")

	flag.Parse()

	if _, err := os.Stat(*logDirFlag); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: log directory %s does not exist\n", *logDirFlag)
		os.Exit(1)
	}

	fmt.Println("Generating report ...")

	switch *outputFlag {
	case "stdout":
        fmt.Println("Write to stdout")
	case "json":
        fmt.Println("Write to JSON file")
	case "html":
        fmt.Println("Write to HTML file")
	default:
		fmt.Fprintf(os.Stderr, "Error: invalid output option %s\n", *outputFlag)
		os.Exit(1)
	}

	if *persistFlag != "" {
        fmt.Println("Save data to SQLite")
	}

	fmt.Println("Report complete!")
}
