package core

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"labit/internal/types"
)


func SaveObject(obj *types.Object, hash string) error {
	objDir := filepath.Join(types.ObjectsDir, hash[:2])
	if err := os.MkdirAll(objDir, 0755); err != nil {
		return fmt.Errorf("failed to create object directory: %v", err)
	}

	objPath := filepath.Join(objDir, hash[2:])
	data, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("failed to marshal object: %v", err)
	}

	return os.WriteFile(objPath, data, 0644)
}

func LoadObject(hash string) (*types.Object, error) {
	objPath := filepath.Join(types.ObjectsDir, hash[:2], hash[2:])
	data, err := os.ReadFile(objPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read object: %v", err)
	}

	var obj types.Object
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, fmt.Errorf("failed to unmarshal object: %v", err)
	}

	return &obj, nil
}

func SaveBlob(content []byte) (string, error) {
	hash := HashContent(content)
	
	obj := &types.Object{
		Type:    types.ObjectTypeBlob,
		Content: content,
	}

	if err := SaveObject(obj, hash); err != nil {
		return "", err
	}

	return hash, nil
}

func SaveCommit(commit *types.Commit) error {
	commitData, err := json.Marshal(commit)
	if err != nil {
		return fmt.Errorf("failed to marshal commit: %v", err)
	}

	obj := &types.Object{
		Type:    types.ObjectTypeCommit,
		Content: commitData,
	}

	return SaveObject(obj, commit.Hash)
}
