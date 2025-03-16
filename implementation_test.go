package lab2

import (
  "fmt"
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestPostfixToInfix(t *testing.T) {
  tests := []struct {
    input    string
    expected string
    hasError bool
  }{
    {"3 4 +", "(3 + 4)", false},
    {"5 1 2 + 4 * + 3 -", "((5 + ((1 + 2) * 4)) - 3)", false},
    {"2 3 4 * +", "(2 + (3 * 4))", false},
    {"2 3 4 ^ *", "(2 * (3 ^ 4))", false},
    {"3 2 ^ 4 + 5 *", "(((3 ^ 2) + 4) * 5)", false},
    {"10 6 9 3 / * +", "(10 + (6 * (9 / 3)))", false},
    {"", "", true},                   
    {"3 +", "", true},                
    {"2 3 4", "", true},               
    {"3 4 5 +", "", true},             
    {"2 a +", "", true},               
  }

  for _, test := range tests {
    result, err := PostfixToInfix(test.input)
    if test.hasError {
      assert.Error(t, err, "expected an error for input: %s", test.input)
    } else {
      assert.NoError(t, err, "unexpected error for input: %s", test.input)
      assert.Equal(t, test.expected, result, "unexpected result for input: %s", test.input)
    }
  }
}

// ExamplePostfixToInfix ілюструє використання функції PostfixToInfix.
func ExamplePostfixToInfix() {
  result, err := PostfixToInfix("4 2 - 3 2 ^ * 5 +")
  if err == nil {
    fmt.Println(result)
  }
  // Output:
  // (((4 - 2) * (3 ^ 2)) + 5)
}