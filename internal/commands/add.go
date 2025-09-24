package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"labit/internal/core"
)

func Add(files []string) error {
	if !core.IsRepository() {
		return fmt.Errorf("not a labit repository")
	}

	for _, pattern := range files {
		if err := addPattern(pattern); err != nil {
			return err
		}
	}

	return nil
}

func addPattern(pattern string) error {
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return fmt.Errorf("invalid pattern '%s': %v", pattern, err)
	}

	if len(matches) == 0 {
		if err := addFile(pattern); err != nil {
			return err
		}
		return nil
	}

	for _, match := range matches {
		if err := addFile(match); err != nil {
			return err
		}
	}

	return nil
}

func addFile(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("cannot stat '%s': %v", path, err)
	}

	// Skip directories
	if info.IsDir() {
		return nil
	}

	// Skip .labit directory
	if strings.Contains(path, ".labit") {
		return nil
	}

	// Read file content
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("cannot read '%s': %v", path, err)
	}

	// Save as blob and get hash
	hash, err := core.SaveBlob(content)
	if err != nil {
		return fmt.Errorf("failed to save blob for '%s': %v", path, err)
	}

	// Add to index
	if err := core.AddToIndex(path, hash, int(info.Mode()), info.Size()); err != nil {
		return fmt.Errorf("failed to add '%s' to index: %v", path, err)
	}

	fmt.Printf("Added '%s'\n", path)
	return nil
}
