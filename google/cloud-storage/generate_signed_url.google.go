package google

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

// GenerateSignedURL generates a signed URL for a GCS object.
func GenerateSignedURL(ctx context.Context, serviceAccountPath, bucketName, objectFullPath string) (string, error) {
	// Initialize the client with the service account
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(serviceAccountPath))
	if err != nil {
		return "", fmt.Errorf("storage.NewClient: %w", err)
	}
	defer client.Close()

	// Generate the signed URL
	objectFullPath = strings.TrimPrefix(objectFullPath, bucketName+"/")
	url, err := client.Bucket(bucketName).SignedURL(objectFullPath, &storage.SignedURLOptions{
		Expires: time.Now().Add(6 * time.Hour), // URL valid for 1 hour
		Method:  "GET",
	})
	if err != nil {
		return "", fmt.Errorf("Bucket(%q).SignedURL: %w", bucketName, err)
	}

	return url, nil
}
