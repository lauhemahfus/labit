package types

import (
	"time"
)

// IndexEntry represents a file in the staging area
type IndexEntry struct {
	Path string `json:"path"`
	Hash string `json:"hash"`
	Mode int    `json:"mode"`
	Size int64  `json:"size"`
}

// Index represents the staging area
type Index struct {
	Entries []IndexEntry `json:"entries"`
}

// Commit represents a commit object
type Commit struct {
	Hash      string    `json:"hash"`
	Message   string    `json:"message"`
	Author    string    `json:"author"`
	Timestamp time.Time `json:"timestamp"`
	Parent    string    `json:"parent,omitempty"`
	Tree      string    `json:"tree"`
}

// ObjectType represents the type of Git object
type ObjectType int

const (
	ObjectTypeBlob ObjectType = iota
	ObjectTypeCommit
)

// Object represents a Git object
type Object struct {
	Type    ObjectType `json:"type"`
	Content []byte     `json:"content"`
}

// Repository configuration
const (
	LabitDir     = ".labit"
	ObjectsDir   = ".labit/objects"
	IndexFile    = ".labit/index"
	HeadFile     = ".labit/HEAD"
	ConfigFile   = ".labit/config"
)
