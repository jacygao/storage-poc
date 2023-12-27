package cosmos

import (
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
		client.CreateItem(item, item.Version)
	}
}

func BenchmarkRead(b *testing.B) {
	client := NewClient()
	item := &storage.Item{}
	for i := 0; i < b.N; i++ {
		client.ReadItem("12816e6f-1d17-48f3-906f-b5486065dc26", "1", item)
	}
}
