package database

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

// Connection creates a new Azure Table client for the given table name.
func Connection(tableName string) (*aztables.Client, error) {
	connectionString := "UseDevelopmentStorage=true"
	if connectionString == "" {
		return nil, fmt.Errorf("AZURE_STORAGE_CONNECTION_STRING is not set")
	}

	serviceClient, err := aztables.NewServiceClientFromConnectionString(connectionString, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create service client: %w", err)
	}

	client := serviceClient.NewClient(tableName)
	return client, nil
}

// Write adds a new entity to the table.
func Write(ctx context.Context, tableName string, pk string, rk string, data map[string]interface{}) error {
	client, err := Connection(tableName)
	if err != nil {
		return err
	}

	entity := aztables.EDMEntity{
		Properties: data,
		Entity:     aztables.Entity{PartitionKey: pk, RowKey: rk},
	}

	json, err := json.Marshal(entity)

	_, err = client.AddEntity(ctx, json, nil)
	if err != nil {
		return fmt.Errorf("failed to add entity: %w", err)
	}
	return nil
}

// ReadSingle fetches a single entity by PartitionKey and RowKey.
func ReadSingle(ctx context.Context, tableName string, pk string, rk string) (*aztables.EDMEntity, error) {
	client, err := Connection(tableName)
	if err != nil {
		return nil, err
	}

	entity, err := client.GetEntity(ctx, pk, rk, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to read entity: %w", err)
	}
	return entity, nil
}

// ReadFilter fetches entities based on a single filter condition.
func ReadFilter(ctx context.Context, tableName string, filter string) ([]aztables.EDMEntity, error) {
	client, err := Connection(tableName)
	if err != nil {
		return nil, err
	}

	pager := client.ListEntities(nil)
	var entities []aztables.EDMEntity

	for pager.NextPage(ctx) {
		for _, entity := range pager.PageResponse().Entities {
			if entity.MatchesFilter(filter) {
				entities = append(entities, entity)
			}
		}
	}

	if err = pager.Err(); err != nil {
		return nil, fmt.Errorf("failed to query entities: %w", err)
	}
	return entities, nil
}

func ReadAll(ctx context.Context, tableName string) ([]aztables.EDMEntity, error) {
	client, err := Connection(tableName)
	if err != nil {
		return nil, err
	}

	pager := client.ListEntities(nil)
	var entities []aztables.EDMEntity

	for pager.NextPage(ctx) {
		for _, entity := range pager.PageResponse().Entities {
			entities = append(entities, entity)
		}
	}

	if err = pager.Err(); err != nil {
		return nil, fmt.Errorf("failed to query entities: %w", err)
	}
	return entities, nil
}

// Delete removes an entity by PartitionKey and RowKey.
func Delete(ctx context.Context, tableName string, pk string, rk string) error {
	client, err := Connection(tableName)
	if err != nil {
		return err
	}

	_, err = client.DeleteEntity(ctx, pk, rk, nil)
	if err != nil {
		return fmt.Errorf("failed to delete entity: %w", err)
	}
	return nil
}

// Update modifies an existing entity in the table.
func Update(ctx context.Context, tableName string, pk string, rk string, data map[string]interface{}) error {
	client, err := Connection(tableName)
	if err != nil {
		return err
	}

	entity := aztables.EDMEntity{
		PartitionKey: pk,
		RowKey:       rk,
		Properties:   data,
	}

	_, err = client.UpdateEntity(ctx, entity, nil)
	if err != nil {
		return fmt.Errorf("failed to update entity: %w", err)
	}
	return nil
}

// Helper function to construct filter strings for Azure Tables.
func BuildFilter(field, value string) string {
	return fmt.Sprintf("%s eq '%s'", field, value)
}
