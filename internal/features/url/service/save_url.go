package service

import "fmt"


func (s *Service) HandleSaveURL(url string, alias string) (string, string, string, error) {
	if alias != "" {
		isAliasAvailable, err := s.repository.GetURL(alias)
		if err == nil && isAliasAvailable != "" {
			return "", "", "", fmt.Errorf("alias is not available")
		}
	}
	if alias == "" {
		alias = randomizeAlias(6)
	}
	surl, err := s.repository.SaveURL(url, alias)
	if err != nil {
		return "", "", "", err
	}
	return surl, alias, returnShortedUrl(alias), nil
}