package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

// Connection creates a new Azure Table client for the given table name.
func Connection(tableName string) (*aztables.Client, error) {
	// build connection string from environment variables
	connectionString := os.Getenv("AZURE_CONNECTION_STRING")
	if connectionString == "" {
		return nil, fmt.Errorf("AZURE_STORAGE_CONNECTION_STRING is not set")
	}

	// create a new service client
	serviceClient, err := aztables.NewServiceClientFromConnectionString(connectionString, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create service client: %w", err)
	}

	// Ensure the table exists, creating it if necessary
	if err := ensureTableExists(serviceClient, tableName); err != nil {
		return nil, err
	}

	// Create and return the table client
	return serviceClient.NewClient(tableName), nil
}

func ensureTableExists(serviceClient *aztables.ServiceClient, tableName string) error {
	const tableAlreadyExistsCode = "TableAlreadyExists"

	_, err := serviceClient.CreateTable(context.Background(), tableName, nil)
	if err != nil {
		// Check if the error is due to the table already existing
		var responseError *azcore.ResponseError
		if errors.As(err, &responseError) && responseError.ErrorCode == tableAlreadyExistsCode {
			return nil // Table already exists, nothing to do
		}
		// Other errors should be propagated
		return fmt.Errorf("failed to create table %q: %w", tableName, err)
	}

	return nil // Table created successfully
}

// Write adds a new entity to the table.
func Write(ctx context.Context, tableName string, pk string, rk string, data map[string]interface{}) error {
	client, err := Connection(tableName)
	if err != nil {
		return err
	}

	entity := aztables.Entity{
		PartitionKey: pk,
		RowKey:       rk,
		Timestamp:    aztables.EDMDateTime(time.Now()),
	}

	EDMEntity := aztables.EDMEntity{
		Entity:     entity,
		Properties: data,
	}

	jsonEntity, err := json.Marshal(EDMEntity)
	if err != nil {
		return err
	}

	_, err = client.AddEntity(ctx, jsonEntity, nil)
	if err != nil {
		return err
	}

	return nil
}

// ReadSingle fetches a single entity by PartitionKey and RowKey.
func ReadSingle(ctx context.Context, tableName string, pk string, rk string) (*aztables.EDMEntity, error) {
	client, err := Connection(tableName)
	if err != nil {
		return nil, err
	}

	resp, err := client.GetEntity(ctx, pk, rk, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to read entity: %w", err)
	}

	//convert from entity to EDMEntity
	var entity aztables.EDMEntity
	err = json.Unmarshal(resp.Value, &entity)
	if err != nil {
		return nil, fmt.Errorf("failed to parse entity: %w", err)
	}

	return &entity, nil
}

// ReadFilter fetches entities based on a single filter condition.
func ReadFilter(ctx context.Context, tableName string, filter string) ([]aztables.EDMEntity, error) {
	// Establish connection to the table
	client, err := Connection(tableName)
	if err != nil {
		return nil, fmt.Errorf("connection error: %w", err)
	}

	// Setup query options
	queryOptions := &aztables.ListEntitiesOptions{
		Filter: &filter,
	}

	// Initialize pager
	pager := client.NewListEntitiesPager(queryOptions)
	var entities []aztables.EDMEntity

	// Iterate through pages
	for pager.More() {
		resp, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to query entities: %w", err)
		}

		// Parse entities in the response
		for _, entity := range resp.Entities {
			var edmEntity aztables.EDMEntity
			if err := json.Unmarshal(entity, &edmEntity); err != nil {
				return nil, fmt.Errorf("failed to parse entity: %w", err)
			}
			entities = append(entities, edmEntity)
		}
	}

	return entities, nil
}

func ReadAll(ctx context.Context, tableName string) ([]aztables.EDMEntity, error) {
	client, err := Connection(tableName)
	if err != nil {
		return nil, err
	}

	pager := client.NewListEntitiesPager(nil)
	var entities []aztables.EDMEntity

	for pager.More() {
		resp, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to query entities: %w", err)
		}

		for _, entity := range resp.Entities {
			var edmEntity aztables.EDMEntity
			err = json.Unmarshal(entity, &edmEntity)
			if err != nil {
				return nil, fmt.Errorf("failed to parse entity: %w", err)
			}

			entities = append(entities, edmEntity)
		}
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
		return errors.New("failed to delete entity: " + err.Error())
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
		Properties: data,
		Entity:     aztables.Entity{PartitionKey: pk, RowKey: rk},
	}

	jsonEntity, err := json.Marshal(entity)
	if err != nil {
		return fmt.Errorf("failed to marshal entity: %w", err)
	}

	_, err = client.UpdateEntity(ctx, jsonEntity, nil)
	if err != nil {
		return fmt.Errorf("failed to update entity: %w", err)
	}
	return nil
}

// Helper function to construct filter strings for Azure Tables.
func BuildFilter(field, value string) string {
	return fmt.Sprintf("%s eq '%s'", field, value)
}

// Helper function to construct filter strings for Azure Tables.
func BuildDuoFilter(field, value, field2, value2 string) string {
	return (field + " eq '" + value + "' and " + field2 + " eq '" + value2 + "'")
	// return fmt.Sprintf("%s eq '%s' and %s eq '%s'", field, value, field2, value2)
}
