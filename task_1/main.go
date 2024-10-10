package main

import (
	"encoding/json"
	"fmt"
	"nested-map/util"
	"os"
	"strings"
)

func main() {
	var jsonInput strings.Builder
	var stdIn string

	// check if input has been passed in through stdin
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Fprintf(os.Stderr, "no json file provided to parse\n")
		os.Exit(1)
	}

	// Read json input from stdin
	for {
		_, err := fmt.Scan(&stdIn)
		if err != nil {
			break
		}
		jsonInput.WriteString(stdIn)
	}
	jsonArray := []map[string]interface{}{}

	err := json.Unmarshal([]byte(jsonInput.String()), &jsonArray)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couln't read json input\n")
		os.Exit(1)
	}

	// recieve argv from the cli
	args := os.Args[1:]
	jsonData, err := json.MarshalIndent(
		util.BuildNestedMap(jsonArray, args),
		"", "  ",
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error formating json output: %v", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, string(jsonData))
}
