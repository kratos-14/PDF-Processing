package service

func (s *service) DownloadFile() error {
	err := s.repo.DownloadFile()
	return nil	
}
