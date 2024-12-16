package repository

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

func Connection() *aztables.ServiceClient {
	connectionString := "DefaultEndpointsProtocol=http;AccountName=devstoreaccount1;AccountKey=Eby8vdM02xW5frD4+T3ONKi7J1LPttwEELEdmAzx4nwAZv1r2cuX6kftoG3nUg==;TableEndpoint=http://127.0.0.1:10002/devstoreaccount1;"

	serviceClient, err := aztables.NewServiceClientFromConnectionString(connectionString, nil)
	tableClient := serviceClient.NewTableClient(tableName)

	if err != nil {
		log.Fatalf("Failed to create service client: %v", err)
	}
	return tableClient
}

func Write(tableName string, pk string, rk string, data map[string]interface{}) {
	serviceClient := Connection()
	tableClient := serviceClient.NewTableClient(tableName)

	entity := aztables.EDMEntity{
		PartitionKey: pk,
		RowKey:       rk,
		Properties:   data,
	}

	_, err := tableClient.AddEntity(context.Background(), entity, nil)
	if err != nil {
		log.Fatalf("Failed to add entity: %v", err)
	}
}

func Create(tableName string, data map[string]interface{}) {
	serviceClient := Connection()
	tableClient := serviceClient.NewTableClient(tableName)

	entity := aztables.EDMEntity{
		PartitionKey: data["PartitionKey"].(string),
		RowKey:       data["RowKey"].(string),
		Properties:   data,
	}

	_, err := tableClient.AddEntity(context.Background(), entity, nil)
	if err != nil {
		log.Fatalf("Failed to add entity: %v", err)
	}
}

func ReadSingle(tableName string, pk string, rk string) {
	serviceClient := Connection()
	tableClient := serviceClient.NewTableClient(tableName)

	readEntity, err := tableClient.GetEntity(context.Background(), pk, rk, nil)
	if err != nil {
		log.Fatalf("Failed to read entity: %v", err)
	}

	return readEntity
}

func ReadFilter(tableName string, filter string, value string) {
	serviceClient := Connection()
	tableClient := serviceClient.NewTableClient(tableName)

	query := aztables.Query{
		Filter: aztables.Filter{
			FilterString: filter + " eq " + value,
		},
	}

	readEntities, err := tableClient.QueryEntities(context.Background(), &query, nil)
	if err != nil {
		log.Fatalf("Failed to read entities: %v", err)
	}

	return readEntities
}

func Delete(tableName string, pk string, rk string) {
	serviceClient := Connection()
	tableClient := serviceClient.NewTableClient(tableName)

	_, err := tableClient.DeleteEntity(context.Background(), pk, rk, nil)
	if err != nil {
		log.Fatalf("Failed to delete entity: %v", err)
	}
}

func Update(tableName string, pk string, rk string, data map[string]interface{}) {
	serviceClient := Connection()
	tableClient := serviceClient.NewTableClient(tableName)

	entity := aztables.EDMEntity{
		PartitionKey: pk,
		RowKey:       rk,
		Properties:   data,
	}

	_, err := tableClient.UpdateEntity(context.Background(), entity, nil)
	if err != nil {
		log.Fatalf("Failed to update entity: %v", err)
	}
}
