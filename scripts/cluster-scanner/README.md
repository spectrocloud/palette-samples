# Cluster Scanner

The **Cluster Scanner** tool uses the Palette Go SDK to scan your Palette environment and identify clusters that have been active for more than 24 hours.

## Prerequisites

- Go version 1.22.5 or later
- Git
- The `palette-samples` repository available locally
- A Palette acount
- A Palette API key

## Usage

1. Open a terminal window and export your Palette URL. Replace `<your-palette-url>` with your Palette URL, for example, `console.spectrocloud.com`.

   ```shell
   export PALETTE_HOST=<your-palette-url>
   ```

2. Export your Palette API key. Replace `<your-palette-api-key>` with your Palette API key.

   ```shell
   export PALETTE_API_KEY=<your-palette-api-key>
   ```

3. To scan a specific project, export the project's UID. Replace `<your-palette-project>` with the Palette project UID. If no project is provided, the tool assumes a tenant scope and scans clusters across all projects.

   ```shell
   export PALETTE_PROJECT_UID=<your-palette-project>
   ```

4. Navigate to the `cluster-scanner` folder.

   ```shell
   cd cluster-scanner
   ```

5. Issue the command below to download the required Palette SDK modules.

   ```shell
   go get
   ```

6. Execute the `cluster-scanner` application.

   ```shell
   go run .
   ```

   The application will print the clusters that have been active in your Palette environment for more than 24 hours.

   ```text hideClipboard
   time=2024-10-28T21:21:47.516-04:00 level=INFO msg="Setting scope to tenant."
   time=2024-10-28T21:21:47.516-04:00 level=INFO msg="Searching for clusters..."
   time=2024-10-28T21:21:48.297-04:00 level=INFO msg="The aws cluster named aws-test has been running for 2 weeks 6 days 2 hours. Are you sure you need this cluster?"
   ```
