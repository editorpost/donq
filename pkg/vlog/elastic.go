package vlog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/avast/retry-go"
	es "github.com/elastic/go-elasticsearch/v8"
	"log/slog"
	"time"
)

type ElasticIngester struct {
	endpoint string
	mapper   MapperFn
	client   *es.Client
}

func NewElasticIngest(endpoint string, mapper MapperFn) (*ElasticIngester, error) {

	var (
		client *es.Client
		err    error
	)

	connect := func() error {
		client, err = es.NewClient(es.Config{
			Addresses: []string{endpoint},
		})
		return err
	}

	err = retry.Do(connect, retry.Attempts(3), retry.Delay(1*time.Second))
	if err != nil {
		return nil, err
	}

	return &ElasticIngester{endpoint: endpoint, mapper: mapper, client: client}, nil
}

func (e ElasticIngester) Sender() SenderFn {

	return func(logs []slog.Record) error {

		// marshal records to lines
		buf := e.Buffer(logs)

		// try to send the logs
		err := retry.Do(
			func() error {
				res, err := e.client.Bulk(bytes.NewReader(buf.Bytes()))
				if err != nil {
					return err
				}

				res.Body.Close()
				if res.IsError() {
					return fmt.Errorf("Failed to index data: %s", res.Status())
				}

				return err
			},
			retry.Attempts(10), retry.Delay(15*time.Second),
		)

		return err
	}
}

// Buffer from mapped records
func (e ElasticIngester) Buffer(logs []slog.Record) bytes.Buffer {

	var buf bytes.Buffer

	buf.WriteString(`{"create":{}}` + "\n")

	for _, record := range logs {

		data, err := json.Marshal(e.mapper(record))
		if err != nil {
			fmt.Printf("error marshaling document: %v", err)
			continue
		}

		// append record as a line
		data = append(data, "\n"...)
		buf.Grow(len(data))
		buf.Write(data)
	}

	return buf
}

// StdoutSender sender for logs that failed to be sent to the primary endpoint
func StdoutSender(mapper MapperFn) func([]slog.Record) error {

	return func(logs []slog.Record) error {

		for _, record := range logs {

			data, err := json.Marshal(mapper(record))
			if err != nil {
				fmt.Printf("error marshaling document: %v", err)
				continue
			}

			fmt.Println(string(data))
		}

		return nil
	}
}
