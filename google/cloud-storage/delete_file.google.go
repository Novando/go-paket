package google

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func DeleteFile(ctx context.Context, serviceAccountPath, bucket, objectFullPath string) error {
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(serviceAccountPath))
	if err != nil {
		return fmt.Errorf("storage.NewClient: %w", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	o := client.Bucket(bucket).Object(objectFullPath)

	// Optional: set a generation-match precondition to avoid potential race
	// conditions and data corruptions. The request to delete the file.go is aborted
	// if the object's generation number does not match your precondition.
	attrs, err := o.Attrs(ctx)
	if err != nil {
		return fmt.Errorf("object.Attrs: %w", err)
	}
	o = o.If(storage.Conditions{GenerationMatch: attrs.Generation})

	if err = o.Delete(ctx); err != nil {
		return fmt.Errorf("Object(%q).Delete: %w", objectFullPath, err)
	}
	//fmt.Fprintf("Blob %v deleted.\n", object)
	return nil
}
