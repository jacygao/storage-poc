package blob

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"

	storage "github.com/storage-poc/lib"
)

const (
	ENDPOINT      = "DefaultEndpointsProtocol=https;AccountName=jgtestdatastore;AccountKey=xRB84tm1QqfQzODs7T48rGiWBC4kzGo0jKjzWW/TtwlbHGq+URDFdOB7XehyVuXnOxFDGEVmotGC+ASts3hTTQ==;EndpointSuffix=core.windows.net"
	CONTAINERNAME = "testdata"
)

type Client struct {
	BlobClient *azblob.Client
}

func NewClient() *Client {
	client, err := azblob.NewClientFromConnectionString(ENDPOINT, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	// ctx := context.Background()

	// fmt.Printf("Creating a container named %s\n", CONTAINERNAME)
	// if _, err = client.CreateContainer(ctx, CONTAINERNAME, nil); err != nil {
	// 	log.Fatal(err.Error())
	// }

	return &Client{
		BlobClient: client,
	}
}

func (c *Client) CreateItem(ctx context.Context, item *storage.Item) error {
	blobName := "sample-blob"
	data, err := json.Marshal(item)
	if err != nil {
		return err
	}

	// Upload to data to blob storage
	fmt.Printf("Uploading a blob named %s\n", blobName)
	if _, err := c.BlobClient.UploadBuffer(ctx, CONTAINERNAME, blobName, data, &azblob.UploadBufferOptions{}); err != nil {
		return err
	}
	return nil
}

func (c *Client) ReadItem(ctx context.Context, path string, item *storage.Item) error {
	get, err := c.BlobClient.DownloadStream(ctx, CONTAINERNAME, path, nil)
	if err != nil {
		return err
	}

	downloadedData := bytes.Buffer{}
	retryReader := get.NewRetryReader(ctx, &azblob.RetryReaderOptions{})
	if _, err = downloadedData.ReadFrom(retryReader); err != nil {
		return err
	}

	if err := retryReader.Close(); err != nil {
		return err
	}

	if err := json.Unmarshal(downloadedData.Bytes(), item); err != nil {
		return err
	}
	return nil
}
