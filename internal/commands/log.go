package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"labit/internal/core"
	"labit/internal/types"
)


func Log() error {
	if !core.IsRepository() {
		return fmt.Errorf("not a labit repository")
	}

	headData, err := os.ReadFile(types.HeadFile)
	if err != nil {
		return fmt.Errorf("failed to read HEAD: %v", err)
	}

	currentHash := strings.TrimSpace(string(headData))
	if currentHash == "" {
		fmt.Println("No commits yet.")
		return nil
	}

	for currentHash != "" {
		obj, err := core.LoadObject(currentHash)
		if err != nil {
			return fmt.Errorf("failed to load commit %s: %v", currentHash, err)
		}

		if obj.Type != types.ObjectTypeCommit {
			return fmt.Errorf("object %s is not a commit", currentHash)
		}

		var commit types.Commit
		if err := json.Unmarshal(obj.Content, &commit); err != nil {
			return fmt.Errorf("failed to parse commit %s: %v", currentHash, err)
		}

		fmt.Printf("commit %s\n", commit.Hash)
		fmt.Printf("Author: %s\n", commit.Author)
		fmt.Printf("Date: %s\n", commit.Timestamp.Format("Mon Jan 2 15:04:05 2006 -0700"))
		fmt.Printf("\n    %s\n\n", commit.Message)

		currentHash = strings.TrimSpace(commit.Parent)
	}

	return nil
}
