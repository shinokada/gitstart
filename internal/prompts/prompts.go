package prompts

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// DryRunPrompt displays a message when dry run mode is enabled.
func DryRunPrompt(action string) {
	fmt.Printf("[DRY RUN] Would perform: %s\n", action)
}

// PromptUser displays a prompt and returns user input.
func PromptUser(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// PromptSelect displays a selection prompt and returns the chosen option.
func PromptSelect(prompt string, options []string) string {
	fmt.Println(prompt)
	for i, opt := range options {
		fmt.Printf("%d) %s\n", i+1, opt)
	}
	fmt.Print("Select option: ")
	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil {
		return ""
	}
	if choice > 0 && choice <= len(options) {
		return options[choice-1]
	}
	return ""
}
