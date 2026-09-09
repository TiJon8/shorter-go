package service


func (s *Service) HandleRedirectURL(alias string) (string, error) {
	if alias == "" {
		return "", ErrAliasIsNotProvided
	}
	url, err := s.repository.GetURL(alias)
	if err != nil && url == "" {
		return "", ErrNotFound
	}
	return url, nil
}