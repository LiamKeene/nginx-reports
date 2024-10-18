package writers

import (
	"encoding/json"
	"log"
	"os"

	"nginx-reports/internal/parser"
)

func WriteJSON(logs []parser.LogData) {
	f, err := os.Create("output/data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	err = encoder.Encode(logs)
	if err != nil {
		log.Fatal(err)
	}
}
