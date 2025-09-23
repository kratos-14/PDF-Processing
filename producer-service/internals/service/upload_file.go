package service

import (
	"fmt"
	"mime/multipart"
)

func (s *service) UploadFile(form *multipart.Form) error {
	files := form.File["files"]
	if len(files) == 0 {
		return fmt.Errorf("no files Uploaded")
	}
	err := s.repo.UploadFile(files)
	if err != nil {
		return fmt.Errorf("error adding files to bucket")
	}
	return nil
}
