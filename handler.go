package lab2

import (
	"fmt"
	"io"
	"strings"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input     io.Reader
	Output    io.Writer
	Converter func(string) (string, error)
}

func (ch *ComputeHandler) Compute() error {
	expressionBytes, err := io.ReadAll(ch.Input)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}
	expression := string(expressionBytes)
	expression = strings.TrimSpace(expression)

	if expression == "" {
		return fmt.Errorf("empty input expression")
	}

	infixExpression, err := ch.Converter(expression)
	if err != nil {
		return fmt.Errorf("expression conversion error: %w", err)
	}

	_, err = ch.Output.Write([]byte(infixExpression))
	if err != nil {
		return fmt.Errorf("failed to write output: %w", err)
	}
	return nil
}
