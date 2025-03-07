package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ") 
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)

		if input == "exit" {
            fmt.Println("Exiting shell :D")
			break
		}

		processCommand(input)
	}
}

func processCommand(input string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}

	switch parts[0] {
	case "echo":
		fmt.Println(strings.Join(parts[1:], " "))
	default:
		fmt.Println("Unknown command:", parts[0])
	}
}
