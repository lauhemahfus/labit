package core

import (
	"fmt"
	"os"
	"path/filepath"

	"labit/internal/types"
)

func IsRepository() bool {
	_, err := os.Stat(types.LabitDir)
	return err == nil
}

func InitRepository() error {
	if IsRepository() {
		return fmt.Errorf("already a labit repository")
	}

	if err := os.MkdirAll(types.ObjectsDir, 0755); err != nil {
		return fmt.Errorf("failed to create objects directory: %v", err)
	}

	index := &types.Index{Entries: []types.IndexEntry{}}
	if err := SaveIndex(index); err != nil {
		return fmt.Errorf("failed to create index: %v", err)
	}

	if err := os.WriteFile(types.HeadFile, []byte(""), 0644); err != nil {
		return fmt.Errorf("failed to create HEAD file: %v", err)
	}

	config := fmt.Sprintf("[core]\n\trepositoryformatversion = 0\n\tfilemode = true\n")
	if err := os.WriteFile(types.ConfigFile, []byte(config), 0644); err != nil {
		return fmt.Errorf("failed to create config file: %v", err)
	}

	return nil
}

func GetRepositoryRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		labitPath := filepath.Join(dir, types.LabitDir)
		if _, err := os.Stat(labitPath); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return "", fmt.Errorf("not a labit repository")
}
