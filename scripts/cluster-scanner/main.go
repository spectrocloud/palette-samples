package main

import (
	"log/slog"
	"os"

	"github.com/spectrocloud/palette-samples/cluster-scanner/internal"
	"github.com/spectrocloud/palette-sdk-go/client"
)

func main() {

    // Read environment variables
    host := os.Getenv("PALETTE_HOST")
    apiKey := os.Getenv("PALETTE_API_KEY")
    projectUid := os.Getenv("PALETTE_PROJECT_UID")

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))


    if host == "" || apiKey == "" {
        logger.Error("You must specify the PALETTE_HOST and PALETTE_API_KEY environment variables.")
        os.Exit(1)
    }

    // Initialize a Palette client
    paletteClient := client.New(
        client.WithPaletteURI(host),
        client.WithAPIKey(apiKey),
    )

    if projectUid != "" {
        client.WithScopeProject(projectUid)(paletteClient)
        logger.Info("Setting scope to project.")
	} else {
        client.WithScopeTenant()(paletteClient)
        logger.Info("Setting scope to tenant.")
    }

    // Search for clusters
    clusters, err := internal.SearchClusters(paletteClient, logger)
    if err != nil {
        logger.Error("Failed to search cluster summaries", "error", err)
		os.Exit(2)
	}

    // Check active clusters
	if len(clusters) == 0 {
		logger.Warn("There are no clusters running.")
		return
	}
    
    foundOldCluster := internal.SearchOldClusters(clusters, logger)
    
    if !foundOldCluster {
        logger.Info("There are no clusters running for more than 24 hours.")
    }        
}