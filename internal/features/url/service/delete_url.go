package service


func (s *Service) HandleDeleteURL(alias string) (error) {
	if alias == "" {
		return ErrAliasIsNotProvided
	}
	url, err := s.repository.GetURL(alias)
	if err != nil && url == "" {
		return ErrNotFound
	}
	if err := s.repository.DeleteURL(alias); err != nil {
		return err
	}
	return nil
}