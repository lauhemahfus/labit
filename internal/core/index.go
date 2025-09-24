package core

import (
	"encoding/json"
	"fmt"
	"os"

	"labit/internal/types"
)


func LoadIndex() (*types.Index, error) {
	data, err := os.ReadFile(types.IndexFile)
	if err != nil {
		// If file doesn't exist, return empty index
		if os.IsNotExist(err) {
			return &types.Index{Entries: []types.IndexEntry{}}, nil
		}
		return nil, fmt.Errorf("failed to read index: %v", err)
	}

	var index types.Index
	if err := json.Unmarshal(data, &index); err != nil {
		return nil, fmt.Errorf("failed to unmarshal index: %v", err)
	}

	return &index, nil
}

func SaveIndex(index *types.Index) error {
	data, err := json.Marshal(index)
	if err != nil {
		return fmt.Errorf("failed to marshal index: %v", err)
	}

	return os.WriteFile(types.IndexFile, data, 0644)
}

func AddToIndex(path string, hash string, mode int, size int64) error {
	index, err := LoadIndex()
	if err != nil {
		return err
	}

	for i, entry := range index.Entries {
		if entry.Path == path {
			index.Entries[i] = types.IndexEntry{
				Path: path,
				Hash: hash,
				Mode: mode,
				Size: size,
			}
			return SaveIndex(index)
		}
	}

	entry := types.IndexEntry{
		Path: path,
		Hash: hash,
		Mode: mode,
		Size: size,
	}
	index.Entries = append(index.Entries, entry)

	return SaveIndex(index)
}


func ClearIndex() error {
	index := &types.Index{Entries: []types.IndexEntry{}}
	return SaveIndex(index)
}
