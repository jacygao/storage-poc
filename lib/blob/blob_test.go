package blob

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/uuid"
	storage "github.com/storage-poc/lib"
)

func BenchmarkCreate(b *testing.B) {
	// open input file
	pwd, _ := os.Getwd()
	fi, err := ioutil.ReadFile(pwd + "/../../testdata/test.kql")
	if err != nil {
		panic(err)
	}

	item := &storage.Item{
		ID:       uuid.NewString(),
		FileName: uuid.NewString(),
		Path:     "/a/b",
		Query:    string(fi),
		Version:  "1",
	}

	client := NewClient()
	for i := 0; i < b.N; i++ {
		client.CreateItem(context.Background(), item)
	}
}

func BenchmarkRead(b *testing.B) {
	client := NewClient()
	item := &storage.Item{}
	for i := 0; i < b.N; i++ {
		client.ReadItem(context.Background(), "sample-blob", item)
	}
}
