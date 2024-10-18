package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var logFormat = `(?P<ip>\d+\.\d+\.\d+\.\d+) - - \[(?P<timestamp>[^\]]+)\] "(?P<method>[A-Z]+) (?P<url>[^\s]+) HTTP/[0-9.]+" (?P<status>\d+) (?P<size>\d+) "(?P<referrer>[^"]*)" "(?P<user_agent>[^"]+)"`

type LogData struct {
	IP        string
	Timestamp string
	Method    string
	URL       string
	Status    string
	Size      string
	UserAgent string
}

func ParseLogs(logDir string) []LogData {
	var allLogs []LogData

	logFiles, err := os.ReadDir(logDir)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range logFiles {
		filePath := filepath.Join(logDir, file.Name())
		logs, _ := parseFile(filePath)
		allLogs = append(allLogs, logs...)
	}
	return allLogs
}

func parseFile(filePath string) ([]LogData, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	i := 0
	scanner := bufio.NewScanner(f)
	var logData []LogData

	for scanner.Scan() {
		i += 1
		line := scanner.Text()
		data, err := parseLogLine(line)
		if err != nil {
			fmt.Printf("Error parsing line %d: %v", i, err)
		} else {
			logData = append(logData, data)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return logData, err
}

func parseLogLine(line string) (LogData, error) {
	regex := regexp.MustCompile(logFormat)
	match:= regex.FindStringSubmatch(line)

    if match == nil {
    	return LogData{}, fmt.Errorf("Line didn't match expected format")
    }

    result := make(map[string]string)
    for i, name := range regex.SubexpNames() {
        if i != 0 && name != "" {
            result[name] = match[i]
        }
    }

    logData := LogData{
        IP: result["ip"],
        Timestamp: result["timestamp"],
        Method: result["method"],
        URL: result["url"],
        UserAgent: result["user_agent"],
        Status: result["status"],
        Size: result["size"],
    }
    return logData, nil
}
