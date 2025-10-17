package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PromptUntilValid(prompt string, validate func(string) error) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print(prompt)
		if !scanner.Scan() {
			return "", scanner.Err()
		}
		
		value := strings.TrimSpace(scanner.Text())
		if err := validate(value); err != nil {
			fmt.Println("Error:", err)
			continue
		}
		
		return value, nil
	}
}
