package main

import (
	"fmt"
	"os"

	"labit/internal/commands"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "init":
		if err := commands.Init(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "add":
		if len(args) == 0 {
			fmt.Fprintf(os.Stderr, "Error: no files specified\n")
			os.Exit(1)
		}
		if err := commands.Add(args); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "commit":
		message := ""
		if len(args) >= 2 && args[0] == "-m" {
			message = args[1]
		} else {
			fmt.Fprintf(os.Stderr, "Error: commit message required. Use: labit commit -m \"message\"\n")
			os.Exit(1)
		}
		if err := commands.Commit(message); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "log":
		if err := commands.Log(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "status":
		if err := commands.Status(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Labit - Simple Version Control System")
	fmt.Println("\nUsage:")
	fmt.Println("  labit init                    Initialize a new repository")
	fmt.Println("  labit add <file>...           Add files to staging area")
	fmt.Println("  labit commit -m \"message\"     Commit staged changes")
	fmt.Println("  labit log                     Show commit history")
	fmt.Println("  labit status                  Show repository status")
}
