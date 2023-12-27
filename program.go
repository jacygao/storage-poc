package main

import (
	"context"

	"github.com/google/uuid"
	storage "github.com/storage-poc/lib"
	"github.com/storage-poc/lib/blob"
	"github.com/storage-poc/lib/cosmos"
)

func main() {
	cosmosClient := cosmos.NewClient()
	item := &storage.Item{
		ID:       uuid.NewString(),
		FileName: uuid.NewString(),
		Path:     "/a/b",
		Query:    "abcde",
		Version:  "1",
	}

	cosmosClient.CreateItem(item, "1")

	blobCient := blob.NewClient()
	blobCient.CreateItem(context.Background(), item)

}
