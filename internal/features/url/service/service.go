package service

import (
	"errors"
	"github.com/TiJon8/shorter-go/internal/storage"
	"math/rand"
	"time"
)


var (
	ErrAliasIsNotProvided = errors.New("alias is not provided")
	ErrAliasIsNotAvailable = errors.New("alias is not available")
	ErrNotFound = errors.New("not found")
)

type repository interface {
	SaveURL(url string, alias string) (string, error)
	GetURL(alias string) (string, error)
	PatchURL(newUrl string, alias string) (string, error)
	DeleteURL(alias string) (error)
}

type Service struct {
	repository repository
}

func InitService(storage *storage.Storage) *Service {
	return &Service{repository: storage}
}


func randomizeAlias(size int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + "0123456789")

	c := make([]rune, size)
	for i := range c {
		c[i] = chars[rnd.Intn(len(chars))]
	}

	return string(c)
}

func returnShortedUrl(alias string) string {
	return "/"+alias
}