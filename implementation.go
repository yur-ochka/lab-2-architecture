package lab2

import (
  "fmt"
  "strconv"
  "strings"
)

// PostfixToInfix перетворює постфіксний вираз у інфіксний
func PostfixToInfix(input string) (string, error) {
  if len(input) == 0 {
    return "", fmt.Errorf("input string is empty")
  }

  operators := "+-*/^"
  stack := []string{}
  tokens := strings.Fields(input)

  for _, token := range tokens {
    if strings.Contains(operators, token) && len(token) == 1 {
      if len(stack) < 2 {
        return "", fmt.Errorf("invalid expression: not enough operands")
      }

      op2 := stack[len(stack)-1]
      op1 := stack[len(stack)-2]
      stack = stack[:len(stack)-2]

      infix := fmt.Sprintf("(%s %s %s)", op1, token, op2)
      stack = append(stack, infix)
    } else {

      if _, err := strconv.Atoi(token); err != nil {
        return "", fmt.Errorf("invalid token: %s", token)
      }
      stack = append(stack, token)
    }
  }

  if len(stack) != 1 {
    return "", fmt.Errorf("invalid expression: leftover operands")
  }

  return stack[0], nil
}

//breacking the build
tralalelo tralala 
bombardiro crocodilo