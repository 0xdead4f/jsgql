package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	operationRegex = regexp.MustCompile(`(mutation|query)\s+(\w+)\s*\([^)]*\)\s*{(?:[^{}]*|{[^{}]*})*}`)
)

type Operation struct {
	Name string `json:"name"`
	File string `json:"file"`
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 && !isInputFromPipe() {
		showUsage()
		os.Exit(1)
	}

	var input io.Reader
	if flag.NArg() > 0 {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		input = os.Stdin
	}

	content, err := io.ReadAll(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	text := string(content)
	operations := operationRegex.FindAllStringSubmatch(text, -1)

	if flag.NArg() > 0 {
		for _, op := range operations {
			result := Operation{
				Name: op[0],
				File: flag.Arg(0),
			}
			jsonResult, _ := json.Marshal(result)
			fmt.Println(string(jsonResult))
		}
	} else {
		for _, op := range operations {
			opStr := regexp.MustCompile(`\s+`).ReplaceAllString(op[0], " ")
			fmt.Printf("%s\n", strings.TrimSpace(opStr))
		}
	}
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func showUsage() {
	fmt.Fprintf(os.Stderr, "Usage: jsgql [file]\n")
	fmt.Fprintf(os.Stderr, "Usage: cat [file] | jsgql\n")

}
