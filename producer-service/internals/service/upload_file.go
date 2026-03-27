package service

import (
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
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
	// topic := "my-topic"
	for key, val := range fileMaps {
		data := map[string]string{
			val: key.Hex(),
		}
		msg, err := json.Marshal(data)
		if err != nil {
			return err
		}
		err = s.broker.Produce(s.key, msg)
		if err != nil {
			log.Printf("error: %v", err)
		}
	}
	return nil
}
