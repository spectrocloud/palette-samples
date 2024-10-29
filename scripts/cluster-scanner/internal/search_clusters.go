package internal

import (
	"log/slog"

	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client"
)

func SearchClusters (paletteClient *client.V1Client, logger *slog.Logger) ([]*models.V1SpectroClusterSummary, error) {
	// Search for clusters
	logger.Info("Searching for clusters...")
	clusters, err := paletteClient.SearchClusterSummaries(&models.V1SearchFilterSpec{}, []*models.V1SearchFilterSortSpec{})

	return clusters, err
}