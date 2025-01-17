package core

import (
	"github.com/blugelabs/bluge"
)

var ZINC_INDEX_LIST map[string]*Index

var ZINC_SYSTEM_INDEX_LIST map[string]*Index

func init() {
	ZINC_SYSTEM_INDEX_LIST, _ = LoadZincSystemIndexes()
	ZINC_INDEX_LIST, _ = LoadZincIndexesFromDisk()
	s3List, _ := LoadZincIndexesFromS3()
	// Load the indexes from disk.
	for k, v := range s3List {
		ZINC_INDEX_LIST[k] = v
	}
}

type Index struct {
	Name          string            `json:"name"`
	Writer        *bluge.Writer     `json:"blugeWriter"`
	CachedMapping map[string]string `json:"mapping"`
	IndexType     string            `json:"index_type"`   // "system" or "user"
	StorageType   string            `json:"storage_type"` // disk, memory, s3
}
