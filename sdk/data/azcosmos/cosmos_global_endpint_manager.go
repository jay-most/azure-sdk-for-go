package azcosmos

import (
	"context"
	"log"
	"net/url"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

const DefaultBackgroundRefreshLocationTimeIntervalInMS = 5 * 60 * 1000
const MinimumIntervalForNonForceRefreshLocationInMS = "50000"

type AccountProperties struct {
	ReadRegions                  []accountRegion `json:"readRegions"`
	WriteRegions                 []accountRegion `json:"writeRegions"`
	EnableMultipleWriteLocations bool            `json:"enableMultipleWriteLocations"`
}

type GlobalEndpointManager struct {
	client           *Client
	preferredRegions []string
	locationCache    *locationCache
}

func readDatabaseAccountProperties(client *azcosmos.CosmosClient) (*azcosmos.DatabaseAccountProperties, error) {
	// Create a new CosmosDatabase instance for the "admin" database
	database := client.GetDatabase("admin")

	// Get the database account properties
	properties, err := database.Read(context.Background())
	if err != nil {
		return nil, err
	}

	return properties, nil
}

func NewGlobalEndpointManager(endpoint string, key azcore.TokenCredential, options *ClientOptions) (*GlobalEndpointManager, error) {

	client, err := NewClient(endpoint, key, options)
	if err != nil {
		return nil, err
	}

	endpoint, err = client.Endpoint(), nil

	parsedEndpoint, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	preferredRegions := options.PreferredRegions
	return &GlobalEndpointManager{
		client:           client,
		preferredRegions: preferredRegions,
		locationCache:    newLocationCache(preferredRegions, parsedEndpoint),
	}, nil
}

func (gem *GlobalEndpointManager) GetWriteEndpoints() ([]url.URL, error) {
	return gem.locationCache.writeEndpoints()
}

func (gem *GlobalEndpointManager) GetAccountProperties() ([]url.URL, error) {

}

func (gem *GlobalEndpointManager) GetReadEndpoints() ([]url.URL, error) {
	return gem.locationCache.readEndpoints()
}

func (gem *GlobalEndpointManager) GetDefaultEndpoint() (url.URL, error) {
	endpoint, err := url.Parse(gem.client.endpoint)
	if err != nil {
		return url.URL{}, err
	}
	return *endpoint, nil
}

func (gem *GlobalEndpointManager) GetLocation(endpoint url.URL) string {
	return gem.locationCache.getLocation(endpoint)
}

func (gem *GlobalEndpointManager) MarkEndpointUnavailableForRead(endpoint url.URL) error {
	return gem.locationCache.markEndpointUnavailableForRead(endpoint)
}

func (gem *GlobalEndpointManager) MarkEndpointUnavailableForWrite(endpoint url.URL) error {
	return gem.locationCache.markEndpointUnavailableForWrite(endpoint)
}

func (gem *GlobalEndpointManager) Update(accountProperties accountProperties) error {
	gem.locationCache.mapMutex.Lock()
	defer gem.locationCache.mapMutex.Unlock()

	writeRegions := make([]accountRegion, len(accountProperties.WriteRegions))
	copy(writeRegions, accountProperties.WriteRegions)
	readRegions := make([]accountRegion, len(accountProperties.ReadRegions))
	copy(readRegions, accountProperties.ReadRegions)

	// Update the location cache
	err := gem.locationCache.update(writeRegions, readRegions, gem.preferredLocs, &accountProperties.EnableMultipleWriteLocations)
	if err != nil {
		return err
	}

	// Update the client's default endpoint
	endpoints, err := gem.locationCache.writeEndpoints()
	if err != nil {
		return err
	}
	if len(endpoints) > 0 {
		gem.client.endpoint = endpoints[0].String()
	}

	return nil
}

func (gem *GlobalEndpointManager) RefreshStaleEndpoints() {
	gem.locationCache.refreshStaleEndpoints()
}

func (gem *GlobalEndpointManager) IsEndpointUnavailable(endpoint url.URL, ops requestedOperations) bool {
	return gem.locationCache.isEndpointUnavailable(endpoint, ops)
}

func (gem *GlobalEndpointManager) CanUseMultipleWriteLocations() bool {
	return gem.locationCache.canUseMultipleWriteLocs()
}

func (gem *GlobalEndpointManager) DatabaseAccountRead(dbAcct accountProperties) error {
	return gem.locationCache.databaseAccountRead(dbAcct)
}

func (gem *GlobalEndpointManager) PreferredLocations() []string {
	return gem.preferredLocs
}

func (gem *GlobalEndpointManager) SetPreferredLocations(locations []string) {
	gem.preferredLocs = locations
	gem.locationCache.update(nil, nil, locations, nil)
}

func (gem *GlobalEndpointManager) UpdateLocationCache(ctx context.Context) {
	// Create a background refresh timer
	refreshTimer := time.NewTicker(DefaultBackgroundRefreshLocationTimeIntervalInMS)
	defer refreshTimer.Stop()

	// Initial refresh of the location cache
	err := gem.RefreshLocationCache(ctx)
	if err != nil {
		log.Printf("Failed to refresh the location cache: %v", err)
	}

	for {
		select {
		case <-ctx.Done():
			// Context cancelled, stop the location cache refresh
			return
		case <-refreshTimer.C:
			// Refresh the location cache periodically
			err := gem.RefreshLocationCache(ctx)
			if err != nil {
				log.Printf("Failed to refresh the location cache: %v", err)
			}
		}
	}
}

func (gem *GlobalEndpointManager) RefreshLocationCache(ctx context.Context) error {
	// Set the minimum interval for non-force refresh
	(*locationCache).refreshStaleEndpoints(gem.locationCache)
	minimumInterval, ok := azruntime.PipelineOptions.(MinimumIntervalForNonForceRefreshLocationInMS).(time.Duration)
	if !ok {
		minimumInterval = DefaultBackgroundRefreshLocationTimeIntervalInMS
	}

	// Check if the refresh interval has elapsed
	lastRefreshTime := gem.locationCache.lastUpdateTime
	elapsed := time.Since(lastRefreshTime)
	if elapsed < minimumInterval {
		return nil
	}

	// Acquire a lock to update the location cache
	gem.locationCache.mapMutex.Lock()
	defer gem.locationCache.mapMutex.Unlock()

	// Retrieve the database account properties
	accountProperties, err := databaseAccountLocationsInfo
	if err != nil {
		return err
	}

	// Update the location cache with the account properties
	err = gem.locationCache.update(accountProperties.availWriteLocations, accountProperties.availReadLocations, gem.preferredLocs, &gem.locationCache.enableMultipleWriteLocations)
	if err != nil {
		return err
	}

	// Update the client's default endpoint
	endpoints, err := gem.locationCache.writeEndpoints()
	if err != nil {
		return err
	}
	if len(endpoints) > 0 {
		gem.client.endpoint = endpoints[0].String()
	}

	return nil
}
