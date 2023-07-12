package azcosmos

import (
	"net/url"
	"testing"
)

func TestNewClientFromConnStrSuccess(t *testing.T) {
	connStr := "AccountEndpoint=http://127.0.0.1:80;AccountKey=dG9fYmFzZV82NA==;"
	client, err := NewClientFromConnectionString(connStr, nil)
	if err != nil {
		t.Fatal(err)
	}

	actualEnpoint := client.endpoint
	expectedEndpoint := "http://127.0.0.1:80"
	if actualEnpoint != expectedEndpoint {
		t.Errorf("Expected %v, but got %v", expectedEndpoint, actualEnpoint)
	}
}

func TestMarkEndpointUnavailableForRead(t *testing.T) {
	// Create a new Cosmos DB client
	client := azcosmos.NewClient("<your-connection-string>")

	// Create a new GlobalEndpointManager with preferred locations
	preferredLocations := []string{"East US", "West US"}
	gem := azcosmos.NewGlobalEndpointManager(client, preferredLocations)

	// Mark an endpoint as unavailable for read
	endpoint := url.URL{
		Scheme: "https",
		Host:   "example.com",
	}
	err := gem.MarkEndpointUnavailableForRead(endpoint)
	if err != nil {
		t.Fatalf("Failed to mark endpoint unavailable for read: %v", err)
	}

	// Verify the endpoint is marked as unavailable for read
	if !gem.IsEndpointUnavailable(endpoint, azcosmos.Read) {
		t.Errorf("Expected endpoint to be marked as unavailable for read")
	}
}

func TestUpdate(t *testing.T) {
	// Create a new Cosmos DB client
	client := azcosmos.NewClient("<your-connection-string>")

	// Create a new GlobalEndpointManager with preferred locations
	preferredLocations := []string{"East US", "West US"}
	gem := azcosmos.NewGlobalEndpointManager(client, preferredLocations)

	// Create account properties with updated regions
	accountProps := azcosmos.AccountProperties{
		WriteRegions: []azcosmos.AccountRegion{
			{
				Name:     "East US",
				Endpoint: "https://eastus.documents.azure.com/",
			},
			{
				Name:     "West US",
				Endpoint: "https://westus.documents.azure.com/",
			},
		},
		ReadRegions: []azcosmos.AccountRegion{
			{
				Name:     "East US",
				Endpoint: "https://eastus.documents.azure.com/",
			},
		},
		EnableMultipleWriteLocations: true,
	}

	// Update the location cache
	err := gem.Update(accountProps)
	if err != nil {
		t.Fatalf("Failed to update location cache: %v", err)
	}

	// Verify the location cache is updated
	writeEndpoints, err := gem.GetWriteEndpoints()
	if err != nil {
		t.Fatalf("Failed to get write endpoints: %v", err)
	}
	if len(writeEndpoints) != 2 {
		t.Errorf("Expected 2 write endpoints, got %d", len(writeEndpoints))
	}

	readEndpoints, err := gem.GetReadEndpoints()
	if err != nil {
		t.Fatalf("Failed to get read endpoints: %v", err)
	}
	if len(readEndpoints) != 1 {
		t.Errorf("Expected 1 read endpoint, got %d", len(readEndpoints))
	}

	defaultEndpoint := gem.GetDefaultEndpoint()
	if defaultEndpoint.String() != "https://eastus.documents.azure.com/" {
		t.Errorf("Expected default endpoint to be 'https://eastus.documents.azure.com/', got '%s'", defaultEndpoint.String())
	}
}

func TestRefreshStaleEndpoints(t *testing.T) {
	// Create a new Cosmos DB client
	client := azcosmos.NewClient("<your-connection-string>")

	// Create a new GlobalEndpointManager with preferred locations
	preferredLocations := []string{"East US", "West US"}
	gem := azcosmos.NewGlobalEndpointManager(client, preferredLocations)

	// Refresh the stale endpoints
	gem.RefreshStaleEndpoints()

	// TODO: Add assertions to verify the stale endpoints are refreshed
}

func TestCanUseMultipleWriteLocations(t *testing.T) {
	// Create a new Cosmos DB client
	client := azcosmos.NewClient("<your-connection-string>")

	// Create a new GlobalEndpointManager with preferred locations
	preferredLocations := []string{"East US", "West US"}
	gem := azcosmos.NewGlobalEndpointManager(client, preferredLocations)

	// Verify the ability to use multiple write locations
	canUseMultipleWrite := gem.CanUseMultipleWriteLocations()
	if !canUseMultipleWrite {
		t.Errorf("Expected to be able to use multiple write locations")
	}
}

func TestDatabaseAccountRead(t *testing.T) {
	// Create a new Cosmos DB client
	client := azcosmos.NewClient("<your-connection-string>")

	// Create a new GlobalEndpointManager with preferred locations
	preferredLocations := []string{"East US", "West US"}
	gem := azcosmos.NewGlobalEndpointManager(client, preferredLocations)

	// Create account properties with updated regions
	accountProps := azcosmos.AccountProperties{
		WriteRegions: []azcosmos.AccountRegion{
			{
				Name:     "East US",
				Endpoint: "https://eastus.documents.azure.com/",
			},
			{
				Name:     "West US",
				Endpoint: "https://westus.documents.azure.com/",
			},
		},
		ReadRegions: []azcosmos.AccountRegion{
			{
				Name:     "East US",
				Endpoint: "https://eastus.documents.azure.com/",
			},
			{
				Name:     "West US",
				Endpoint: "https://westus.documents.azure.com/",
			},
		},
		EnableMultipleWriteLocations: true,
	}

	// Update the location cache with account properties
	err := gem.DatabaseAccountRead(accountProps)
	if err != nil {
		t.Fatalf("Failed to update location cache: %v", err)
	}

	// Verify the location cache is updated
	writeEndpoints, err := gem.GetWriteEndpoints()
	if err != nil {
		t.Fatalf("Failed to get write endpoints: %v", err)
	}
	if len(writeEndpoints) != 2 {
		t.Errorf("Expected 2 write endpoints, got %d", len(writeEndpoints))
	}

	readEndpoints, err := gem.GetReadEndpoints()
	if err != nil {
		t.Fatalf("Failed to get read endpoints: %v", err)
	}
	if len(readEndpoints) != 2 {
		t.Errorf("Expected 2 read endpoints, got %d", len(readEndpoints))
	}
}

// ... continue with more unit tests to exhaust other methods ...
