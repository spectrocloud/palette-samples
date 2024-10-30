package internal

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func SearchOldClusters (clusters []*models.V1SpectroClusterSummary, logger *slog.Logger) (bool, string) {
	// Variable to keep track of any found clusters older than 24h
	foundOldCluster := false

	// List the clusters that are running for more than 24h
	for _, cluster := range clusters {
		timeValue := time.Time(cluster.Metadata.CreationTimestamp)
		clusterAge := time.Now().Sub(timeValue)

		if clusterAge.Hours() >= 24 {
			foundOldCluster = true
			message := fmt.Sprintf("The %s cluster named %s has been running for %s. Are you sure you need this cluster?", cluster.SpecSummary.CloudConfig.CloudType, cluster.Metadata.Name, PrintFormattedAge(clusterAge))
			logger.Info(message)
		} 
	}
	
	return foundOldCluster, message
}