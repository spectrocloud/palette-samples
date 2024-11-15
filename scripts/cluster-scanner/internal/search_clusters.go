package internal

import (
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client"
)

// SearchClusters retrieves a list of cluster summaries from the Palette client.
func SearchClusters(paletteClient *client.V1Client) ([]*models.V1SpectroClusterSummary, error) {
	// Search for clusters using the provided Palette client with default filters and sorting.
	return paletteClient.SearchClusterSummaries(&models.V1SearchFilterSpec{}, []*models.V1SearchFilterSortSpec{})
}