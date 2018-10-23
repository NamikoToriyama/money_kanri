package app

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"

	"cloud.google.com/go/storage"
	"google.golang.org/appengine/log"
)

// MoneyItem is the items of csv.
type MoneyItem struct {
	numberID int
	product  string
	company  string
	value    int
	stock    int
}

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

	info, err := parseLog(reader, ctx)
	if err != nil {
		log.Errorf(ctx, "failed to parse ninjalog: %v", err)
		return err
	}
	log.Debugf(ctx, "info: %v", info)
	return nil
}

func parseLog(info io.Reader, ctx context.Context) (*[]MoneyItem, error) {
	scanner := bufio.NewScanner(info)
	items := []MoneyItem{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "number") {
			continue
		} else {
			log.Debugf(ctx, "line: %v", line)
			item, err := logItems(line)
			if err != nil {
				return nil, err
			}
			items = append(items, *item)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &items, nil
}

func logItems(line string) (*MoneyItem, error) {

	lineitems := strings.Split(line, ",")
	numberIDItem, err := strconv.Atoi(lineitems[0])
	if err != nil {
		return nil, fmt.Errorf("failed to prefix number to int: %v", err)
	}
	valueItem, err := strconv.Atoi(lineitems[3])
	if err != nil {
		return nil, fmt.Errorf("failed to prefix value to int: %v", err)
	}
	stockItem, err := strconv.Atoi(lineitems[4])
	if err != nil {
		return nil, fmt.Errorf("failed to prefix stock to int: %v", err)
	}
	items := &MoneyItem{
		numberID: numberIDItem,
		product:  lineitems[1],
		company:  lineitems[2],
		value:    valueItem,
		stock:    stockItem,
	}
	return items, nil
}
