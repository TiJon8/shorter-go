package service

import "fmt"


func (s *Service) HandlePatchURL(newUrl string, alias string) (string, string, error) {
	if alias == "" {
		return "", "", ErrAliasIsNotProvided
	}
	url, err := s.repository.GetURL(alias)
	if err != nil && url == "" {
		return "", "", ErrNotFound
	}
	if url == newUrl {
		return newUrl, returnShortedUrl(alias), nil
	}

	patched, err := s.repository.PatchURL(newUrl, alias)
	fmt.Println(patched)
	if err != nil {
		return "", "", err
	}
	return patched, returnShortedUrl(alias), nil
}