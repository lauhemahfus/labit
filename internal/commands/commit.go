package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"labit/internal/core"
	"labit/internal/types"
)

func Commit(message string) error {
	if !core.IsRepository() {
		return fmt.Errorf("not a labit repository")
	}

	index, err := core.LoadIndex()
	if err != nil {
		return fmt.Errorf("failed to load index: %v", err)
	}

	if len(index.Entries) == 0 {
		return fmt.Errorf("no changes staged for commit")
	}

	headData, err := os.ReadFile(types.HeadFile)
	if err != nil {
		return fmt.Errorf("failed to read HEAD: %v", err)
	}
	parentHash := string(headData)

	currentUser, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to get current user: %v", err)
	}
	author := fmt.Sprintf("%s <%s@%s>", currentUser.Username, currentUser.Username, currentUser.Username)

	treeHash := createTreeHash(index)

	commit := &types.Commit{
		Message:   message,
		Author:    author,
		Timestamp: time.Now(),
		Parent:    parentHash,
		Tree:      treeHash,
	}

	commitContent := fmt.Sprintf("%s\n%s\n%s\n%s\n%s",
		commit.Message, commit.Author, commit.Timestamp.Format(time.RFC3339),
		commit.Parent, commit.Tree)
	commit.Hash = core.HashString(commitContent)

	if err := saveCommitSnapshot(commit.Hash, index); err != nil {
		return fmt.Errorf("failed to save commit snapshot: %v", err)
	}

	if err := core.SaveCommit(commit); err != nil {
		return fmt.Errorf("failed to save commit: %v", err)
	}

	if err := os.WriteFile(types.HeadFile, []byte(commit.Hash), 0644); err != nil {
		return fmt.Errorf("failed to update HEAD: %v", err)
	}

	if err := core.ClearIndex(); err != nil {
		return fmt.Errorf("failed to clear index: %v", err)
	}

	fmt.Printf("Committed %s\n", commit.Hash[:8])
	fmt.Printf("Author: %s\n", author)
	fmt.Printf("Message: %s\n", message)

	return nil
}

func createTreeHash(index *types.Index) string {
	var allHashes string
	for _, entry := range index.Entries {
		allHashes += fmt.Sprintf("%s:%s\n", entry.Path, entry.Hash)
	}
	return core.HashString(allHashes)
}

func saveCommitSnapshot(commitHash string, index *types.Index) error {
	snapshotDir := filepath.Join(types.LabitDir, "snapshots")
	if err := os.MkdirAll(snapshotDir, 0755); err != nil {
		return err
	}
	snapshot := make(map[string]string)
	for _, entry := range index.Entries {
		snapshot[entry.Path] = entry.Hash
	}

	data, err := json.Marshal(snapshot)
	if err != nil {
		return err
	}

	snapshotPath := filepath.Join(snapshotDir, commitHash)
	return os.WriteFile(snapshotPath, data, 0644)
}
