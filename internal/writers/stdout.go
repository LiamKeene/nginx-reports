package writers

import (
	"fmt"

	"nginx-reports/internal/parser"
)

func WriteToStdout(logs []parser.LogData) {
	for _, log := range logs {
		fmt.Println(log)
	}
}
