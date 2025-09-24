package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"labit/internal/core"
	"labit/internal/types"
)

func Status() error {
	if !core.IsRepository() {
		return fmt.Errorf("not a labit repository")
	}

	index, err := core.LoadIndex()
	if err != nil {
		return fmt.Errorf("failed to load index: %v", err)
	}

	committedFiles, err := getLastCommitFiles()
	if err != nil {
		return fmt.Errorf("failed to get committed files: %v", err)
	}

	workingFiles, err := getWorkingFiles()
	if err != nil {
		return fmt.Errorf("failed to scan working directory: %v", err)
	}

	stagedFiles := make(map[string]types.IndexEntry)
	for _, entry := range index.Entries {
		stagedFiles[entry.Path] = entry
	}

	var staged []string
	var modified []string  
	var untracked []string

	for filePath := range workingFiles {
		if _, isStaged := stagedFiles[filePath]; isStaged {
			staged = append(staged, filePath)
		} else if committedHash, wasCommitted := committedFiles[filePath]; wasCommitted {
			if isModifiedSinceCommit(filePath, committedHash) {
				modified = append(modified, filePath)
			}
		} else {
			untracked = append(untracked, filePath)
		}
	}

	hasOutput := false

	if len(staged) > 0 {
		fmt.Println("Changes to be committed:")
		for _, file := range staged {
			if _, exists := committedFiles[file]; exists {
				fmt.Printf("  modified:   %s\n", file)
			} else {
				fmt.Printf("  new file:   %s\n", file)
			}
		}
		fmt.Println()
		hasOutput = true
	}

	if len(modified) > 0 {
		fmt.Println("Changes not staged for commit:")
		for _, file := range modified {
			fmt.Printf("  modified:   %s\n", file)
		}
		fmt.Println()
		fmt.Println("Use \"labit add <file>...\" to update what will be committed")
		hasOutput = true
	}

	if len(untracked) > 0 {
		fmt.Println("Untracked files:")
		for _, file := range untracked {
			fmt.Printf("  %s\n", file)
		}
		fmt.Println()
		fmt.Println("Use \"labit add <file>...\" to include in what will be committed")
		hasOutput = true
	}

	if !hasOutput {
		fmt.Println("nothing to commit, working tree clean")
	}

	return nil
}

func getLastCommitFiles() (map[string]string, error) {
	files := make(map[string]string)

	headData, err := os.ReadFile(types.HeadFile)
	if err != nil || len(strings.TrimSpace(string(headData))) == 0 {
		return files, nil
	}

	commitHash := strings.TrimSpace(string(headData))
	return loadCommitSnapshot(commitHash)
}

func loadCommitSnapshot(commitHash string) (map[string]string, error) {
	snapshotPath := filepath.Join(types.LabitDir, "snapshots", commitHash)
	
	data, err := os.ReadFile(snapshotPath)
	if err != nil {
		return make(map[string]string), nil
	}

	var snapshot map[string]string
	if err := json.Unmarshal(data, &snapshot); err != nil {
		return make(map[string]string), nil
	}

	return snapshot, nil
}

func isModifiedSinceCommit(filePath, committedHash string) bool {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return true 
	}

	currentHash := core.HashContent(content)
	return currentHash != committedHash
}

func getWorkingFiles() (map[string]bool, error) {
	files := make(map[string]bool)

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.Contains(path, ".labit") || filepath.Base(path) == "labit" {
			return nil
		}
		path = filepath.ToSlash(path)
		if strings.HasPrefix(path, "./") {
			path = path[2:]
		}

		files[path] = true
		return nil
	})

	return files, err
}