package service

import (
	"fmt"
	"log"
	"mime/multipart"

	amqp "github.com/rabbitmq/amqp091-go"
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
	for key, val := range fileMaps {
		msg := &amqp.Publishing{
			ContentType: "application/text",
			Body:        []byte(fmt.Sprintf("%s:%s", key.Hex(), val)),
		}
		err = s.broker.Produce(*msg)
		if err != nil {
			log.Printf("error: %v", err)
		}
	}
	return nil
}
