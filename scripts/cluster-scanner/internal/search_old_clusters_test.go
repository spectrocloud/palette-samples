package internal_test

import (
	"testing"
	"time"

	"github.com/spectrocloud/palette-samples/cluster-scanner/internal"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

type TestCaseSearch struct {
	input []*models.V1SpectroClusterSummary
	output []string
	expectedError error
}

func TestSearchOldClusters (t *testing.T) {
	now := time.Now()
	parseTime := func (now time.Time, input string) time.Time {
		creationDate, _ := time.ParseDuration(input)
		age := now.Add(creationDate)
		return age
	}

	createSummary := func (creationTime string, name string, cloudType string) *models.V1SpectroClusterSummary {
		return &models.V1SpectroClusterSummary {
			Metadata: &models.V1ObjectMeta{
				CreationTimestamp: models.V1Time(parseTime(now, creationTime)),
				Name: name,
			},
			SpecSummary: &models.V1SpectroClusterSummarySpecSummary{
				CloudConfig: &models.V1CloudConfigMeta{
					CloudType: cloudType,
				},
			},
		}
	}

	tc := map[string]TestCaseSearch{
		"one cluster older than 24h":{
			input: []*models.V1SpectroClusterSummary{
				createSummary("-30h", "test-cluster", "aws"),
			},
			output: []string{
				"The aws cluster named test-cluster has been running for 1 days 6 hours. Are you sure you need this cluster?",
			},
		},
		"two clusters older than 24h":{
			input: []*models.V1SpectroClusterSummary{
				createSummary("-30h", "test-cluster", "aws"),
				createSummary("-50h", "test-cluster-azure", "azure"),
			},
			output: []string{
				"The aws cluster named test-cluster has been running for 1 days 6 hours. Are you sure you need this cluster?",
				"The azure cluster named test-cluster-azure has been running for 2 days 2 hours. Are you sure you need this cluster?",
			},
		},
		"one cluster with 24h":{
			input: []*models.V1SpectroClusterSummary{
				createSummary("-24h", "test-cluster", "aws"),
			},
			output: []string{
				"The aws cluster named test-cluster has been running for 1 days . Are you sure you need this cluster?",
			},
		},
		"one cluster with less than 24h":{
			input: []*models.V1SpectroClusterSummary{
				createSummary("-20h", "test-cluster", "aws"),
			},
			output: nil,
		},
		"one cluster with negative age":{
			input: []*models.V1SpectroClusterSummary{
				createSummary("20h", "test-cluster", "aws"),
			},
			output: nil,
		},
	}

	for key, value := range tc {
		t.Run(key, func(t *testing.T) {
			clustersGot, err := internal.SearchOldClusters(value.input)

			if value.expectedError != nil && err == nil {
				t.Errorf("Expected an error, but got none")
			}
			if value.expectedError == nil && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if value.expectedError != nil && err != nil {
				if value.expectedError.Error() != err.Error() {
					t.Errorf("Errors do not match: got %v, want %v", err, value.expectedError.Error())
				}
			}
			if len(clustersGot) != len(value.output) {
				t.Errorf("Lenghts mismatch. Got %v elements, want %v elements", len(clustersGot), len(value.output))
			}
			for _, got := range clustersGot {
				found := false
				for _, want := range value.output {
					if got == want {
						found = true
						return 
					}
				}
				if !found {
					t.Errorf("Got unexpected value %v, want %v", got, value.output)
				}
			}
		})
	}
}