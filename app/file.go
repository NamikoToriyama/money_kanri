package app

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
	"google.golang.org/appengine/log"
)

// GetFile fetches GCS upload file.
func GetFile(ctx context.Context, filename string, bucketID string) error {
	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "failed to create client: %v", err)
		return err
	}
	defer func() {
		if err := client.Close(); err != nil {
			log.Warningf(ctx, "failed to close client: %v", err)
		}
	}()

	var reader io.ReadCloser
	// Creates a Bucket instance.
	reader, err = client.Bucket(bucketID).Object(filename).NewReader(ctx)
	if err != nil {
		log.Errorf(ctx, "failed to get reader: %v", err)
		return err
	}
	closer := func(reader io.ReadCloser) {
		if err := reader.Close(); err != nil {
			log.Warningf(ctx, "failed to close client: %v", err)
		}
	}
	defer closer(reader)
	parseLog(reader)

	// info, err := ninjalog.Parse(filename, reader)
	// if err != nil {
	// 	log.Errorf(ctx, "failed to parse ninjalog: %v", err)
	// 	return nil, err
	// }
	return nil
}

func parseLog(info io.Reader) {

}
