package internal

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// SearchOldClusters checks for clusters that have been running for more than 24 hours.
// It returns a boolean indicating if any old clusters were found and a message
// with the clusters' information.
func SearchOldClusters (clusters []*models.V1SpectroClusterSummary, logger *slog.Logger) (bool, string) {
	// Variable to keep track of any found clusters older than 24 hours
	foundOldCluster := false
	var message string

	// Iterate through the clusters to find those running for more than 24 hours
	for _, cluster := range clusters {
		timeValue := time.Time(cluster.Metadata.CreationTimestamp)
		clusterAge := time.Now().Sub(timeValue)

		if clusterAge.Hours() >= 24 {
			foundOldCluster = true
			message = fmt.Sprintf("The %s cluster named %s has been running for %s. Are you sure you need this cluster?", cluster.SpecSummary.CloudConfig.CloudType, cluster.Metadata.Name, PrintFormattedAge(clusterAge))
		} 
	}
	
	return foundOldCluster, message
}