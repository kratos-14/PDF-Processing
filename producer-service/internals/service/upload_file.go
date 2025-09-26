package service

import (
	"fmt"
	"log"
	"mime/multipart"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (s *service) UploadFile(form *multipart.Form) error {
	files := form.File["files"]
	if len(files) == 0 {
		return fmt.Errorf("no files Uploaded")
	}
	fileMaps, err := s.repo.UploadFile(files)
	if err != nil {
		return fmt.Errorf("error adding files to bucket")
	}
	topic := "my-topic"
	for key, val := range fileMaps {
		msg := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(val),
			Value:          []byte(key.Hex()),
		}
		err = s.broker.Produce(*msg)
		if err != nil {
			log.Printf("error: %v", err)
		}
	}
	return nil
}
