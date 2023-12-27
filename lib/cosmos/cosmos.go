package cosmos

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	storage "github.com/storage-poc/lib"
)

const (
	ENDPOINT      = "https://shadowcatfileapi.documents.azure.com:443/"
	KEY           = "3WRywbvH8qy15WZ2Di2bqdSusgkZRfOXsQiQKHW3wZtZErTGfcxsLRB0Tx1QWZZ3Kh3LPR5yrgQpACDbB7StHQ=="
	DBNAME        = "correlations"
	CONTAINERNAME = "testdata"
)

type Client struct {
	Container *azcosmos.ContainerClient
}

func NewClient() *Client {
	var endpoint = ENDPOINT
	var key = KEY

	cred, err := azcosmos.NewKeyCredential(key)
	if err != nil {
		log.Fatal("Failed to create a credential: ", err)
	}

	// Create a CosmosDB client
	client, err := azcosmos.NewClientWithKey(endpoint, cred, nil)
	if err != nil {
		log.Fatal("Failed to create Azure Cosmos DB client: ", err)
	}

	// Create database client
	if _, err := client.NewDatabase(DBNAME); err != nil {
		log.Fatal("Failed to create database client:", err)
	}

	// Create container client
	containerClient, err := client.NewContainer(DBNAME, CONTAINERNAME)
	if err != nil {
		log.Fatal("Failed to create a container client:", err)
	}

	return &Client{
		Container: containerClient,
	}
}

func (c *Client) ReadItem(itemId string, partitionKey string, item *storage.Item) error {
	// Specifies the value of the partiton key
	pk := azcosmos.NewPartitionKeyString(partitionKey)

	// Read an item
	ctx := context.TODO()
	itemResponse, err := c.Container.ReadItem(ctx, pk, itemId, nil)
	if err != nil {
		return err
	}

	itemResponseBody := &storage.Item{}

	err = json.Unmarshal(itemResponse.Value, itemResponseBody)
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(itemResponseBody, "", "    ")
	if err != nil {
		return err
	}
	//fmt.Printf("%s\n", b)

	if err := json.Unmarshal(b, item); err != nil {
		return err
	}

	//log.Printf("Status %d. Item %v read. ActivityId %s. Consuming %v Request Units.\n", itemResponse.RawResponse.StatusCode, pk, itemResponse.ActivityID, itemResponse.RequestCharge)

	return nil
}

func (c *Client) CreateItem(item *storage.Item, partitionKey string) error {
	// Specifies the value of the partiton key
	pk := azcosmos.NewPartitionKeyString(partitionKey)

	b, err := json.Marshal(item)
	if err != nil {
		return err
	}
	// setting item options upon creating ie. consistency level
	itemOptions := azcosmos.ItemOptions{
		ConsistencyLevel: azcosmos.ConsistencyLevelSession.ToPtr(),
	}
	ctx := context.TODO()
	if _, err := c.Container.CreateItem(ctx, pk, b, &itemOptions); err != nil {
		return err
	}

	//log.Printf("Status %d. Item %v created. ActivityId %s. Consuming %v Request Units.\n", itemResponse.RawResponse.StatusCode, pk, itemResponse.ActivityID, itemResponse.RequestCharge)

	return nil
}
