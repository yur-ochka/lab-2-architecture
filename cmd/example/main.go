package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/yur-ochka/lab-2-architecture"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute (e.g., \"3 4 +\")")
	fileInput       = flag.String("f", "", "File containing the expression")    
	fileOutput      = flag.String("o", "", "File to write the result to (optional)")   
)

func main() {
	flag.Parse()

	if *inputExpression != "" && *fileInput != "" {
		fmt.Fprintln(os.Stderr, "Error: Both -e and -f options provided. Use only one.")
		os.Exit(1)
	}

	var inputReader io.Reader
	if *fileInput != "" {
		inputFile, err := os.Open(*fileInput)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
			os.Exit(1)
		}
		defer inputFile.Close()
		inputReader = inputFile
	} else if *inputExpression != "" { 
		inputReader = strings.NewReader(*inputExpression) 
	} else {
		fmt.Fprintln(os.Stderr, "Error: No expression provided. Use -e or -f option.")
		os.Exit(1)
	}

	var outputWriter io.Writer
	if *fileOutput != "" {
		outputFile, err := os.Create(*fileOutput)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer outputFile.Close()
		outputWriter = outputFile
	} else {
		outputWriter = os.Stdout
	}

	handler := lab2.ComputeHandler{
		Input:     inputReader,
		Output:    outputWriter,
		Converter: lab2.PostfixToInfix,
	}

	err := handler.Compute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing expression: %v\n", err)
		os.Exit(1)
	}
}
