package main

import (
	"flag"
	"fmt"
	"os"

	"nginx-reports/internal/parser"
	"nginx-reports/internal/storage"
	"nginx-reports/internal/writers"
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

	reportData := parser.ParseLogs(*logDirFlag)

	fmt.Println("Generating report ...")

	switch *outputFlag {
	case "stdout":
		writers.WriteToStdout(reportData)
	case "json":
		writers.WriteJSON(reportData)
	case "html":
		writers.WriteHTML(reportData)
	default:
		fmt.Fprintf(os.Stderr, "Error: invalid output option %s\n", *outputFlag)
		os.Exit(1)
	}

	if *persistFlag != "" {
		storage.StoreDB(reportData)
	}

	fmt.Println("Report complete!")
}
