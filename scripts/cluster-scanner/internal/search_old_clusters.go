package internal

import (
	"fmt"
	"time"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// SearchOldClusters checks for clusters that have been running for more than 24 hours.
// It takes a list of cluster summaries and returns a slice of messages with the
// information of the clusters that were found.
func SearchOldClusters(clusters []*models.V1SpectroClusterSummary) ([]string, error) {
	// Slice of strings to keep track of any found clusters older than 24 hours
	var messageArray []string

	// Iterate through the clusters to find those running for more than 24 hours
	for _, cluster := range clusters {
		timeValue := time.Time(cluster.Metadata.CreationTimestamp)
		clusterAge := time.Now().Sub(timeValue)

		if clusterAge.Hours() >= 24 {
			age, err := GetFormattedAge(clusterAge)
			if err != nil {
				return nil, err
			}
			message := fmt.Sprintf("-'%s' cluster (%s) - %s", cluster.Metadata.Name, cluster.SpecSummary.CloudConfig.CloudType, *age)
			messageArray = append(messageArray, message)
		}
	}
	return messageArray, nil
}
