package main

import (
	"log/slog"
	"os"

	"github.com/spectrocloud/palette-samples/cluster-scanner/internal"
	"github.com/spectrocloud/palette-sdk-go/client"
)

func main() {

	// Read environment variables required to initialize the Palette client.
	host := os.Getenv("PALETTE_HOST")
	apiKey := os.Getenv("PALETTE_API_KEY")
	projectUid := os.Getenv("PALETTE_PROJECT_UID")

	// Initialize a logger for structured output.
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Ensure the required environment variables are provided.
	if host == "" || apiKey == "" {
		logger.Error("You must specify the PALETTE_HOST and PALETTE_API_KEY environment variables.")
		os.Exit(1)
	}

	// Initialize a Palette client with the provided host and API key.
	paletteClient := client.New(
		client.WithPaletteURI(host),
		client.WithAPIKey(apiKey),
	)

	// Set the scope for the client based on wether the project UID is provided.
	if projectUid != "" {
		client.WithScopeProject(projectUid)(paletteClient)
		logger.Info("Setting scope to project.")
	} else {
		client.WithScopeTenant()(paletteClient)
		logger.Info("Setting scope to tenant.")
	}

	// Search for clusters using the Palette client and the SearchClusters function.
	logger.Info("Searching for clusters...")
	clusters, err := internal.SearchClusters(paletteClient)
	if err != nil {
		logger.Error("Failed to search cluster summaries", "error", err)
		os.Exit(2)
	}

	// Check if any clusters were found.
	if len(clusters) == 0 {
		logger.Warn("There are no clusters running.")
		return
	}

	// Check for clusters running for more than 24 hours using the SearchOldClusters function.
	messageArray, err := internal.SearchOldClusters(clusters)
	if err != nil {
		logger.Error("Failed to search for old clusters", "error", err)
		os.Exit(2)
	}
	if len(messageArray) != 0 {
		for _, message := range messageArray {
			logger.Info(message)
		}
        return
	}
	logger.Info("There are no clusters running for more than 24 hours.")
}
