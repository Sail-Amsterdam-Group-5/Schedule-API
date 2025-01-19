package tests

import (
	"context"
	"os"
	"testing"

	"schedule-api/database"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConnection(t *testing.T) {
	os.Setenv("AZURE_CONNECTION_STRING", "DefaultEndpointsProtocol=https;AccountName=devstoreaccount1;AccountKey=Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;TableEndpoint=http://127.0.0.1:10002/devstoreaccount1;")
	client, err := database.Connection("testTable")
	require.NoError(t, err)
	assert.NotNil(t, client)
}

func TestWriteAndReadSingle(t *testing.T) {
	os.Setenv("AZURE_CONNECTION_STRING", "DefaultEndpointsProtocol=https;AccountName=devstoreaccount1;AccountKey=Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;TableEndpoint=http://127.0.0.1:10002/devstoreaccount1;")
	ctx := context.Background()
	tableName := "testTable"
	pk := "partitionKey"
	rk := "rowKey"
	data := map[string]interface{}{
		"Name": "Test",
		"Age":  30,
	}

	err := database.Write(ctx, tableName, pk, rk, data)
	require.NoError(t, err)

	entity, err := database.ReadSingle(ctx, tableName, pk, rk)
	require.NoError(t, err)
	assert.Equal(t, pk, entity.PartitionKey)
	assert.Equal(t, rk, entity.RowKey)
	assert.Equal(t, data["Name"], entity.Properties["Name"])
	assert.Equal(t, data["Age"], entity.Properties["Age"])
}

func TestReadFilter(t *testing.T) {
	os.Setenv("AZURE_CONNECTION_STRING", "DefaultEndpointsProtocol=https;AccountName=devstoreaccount1;AccountKey=Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;TableEndpoint=http://127.0.0.1:10002/devstoreaccount1;")
	ctx := context.Background()
	tableName := "testTable"
	filter := database.BuildFilter("Name", "Test")

	entities, err := database.ReadFilter(ctx, tableName, filter)
	require.NoError(t, err)
	assert.NotEmpty(t, entities)
}

func TestReadAll(t *testing.T) {
	os.Setenv("AZURE_CONNECTION_STRING", "DefaultEndpointsProtocol=https;AccountName=devstoreaccount1;AccountKey=Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;TableEndpoint=http://127.0.0.1:10002/devstoreaccount1;")
	ctx := context.Background()
	tableName := "testTable"

	entities, err := database.ReadAll(ctx, tableName)
	require.NoError(t, err)
	assert.NotEmpty(t, entities)
}

func TestDelete(t *testing.T) {
	os.Setenv("AZURE_CONNECTION_STRING", "DefaultEndpointsProtocol=https;AccountName=devstoreaccount1;AccountKey=Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;TableEndpoint=http://127.0.0.1:10002/devstoreaccount1;")
	ctx := context.Background()
	tableName := "testTable"
	pk := "partitionKey"
	rk := "rowKey"

	err := database.Delete(ctx, tableName, pk, rk)
	require.NoError(t, err)

	_, err = database.ReadSingle(ctx, tableName, pk, rk)
	assert.Error(t, err)
}

func TestUpdate(t *testing.T) {
	os.Setenv("AZURE_CONNECTION_STRING", "DefaultEndpointsProtocol=https;AccountName=devstoreaccount1;AccountKey=Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;TableEndpoint=http://127.0.0.1:10002/devstoreaccount1;")
	ctx := context.Background()
	tableName := "testTable"
	pk := "partitionKey"
	rk := "rowKey"
	data := map[string]interface{}{
		"Name": "UpdatedTest",
		"Age":  35,
	}

	err := database.Update(ctx, tableName, pk, rk, data)
	require.NoError(t, err)

	entity, err := database.ReadSingle(ctx, tableName, pk, rk)
	require.NoError(t, err)
	assert.Equal(t, data["Name"], entity.Properties["Name"])
	assert.Equal(t, data["Age"], entity.Properties["Age"])
}
