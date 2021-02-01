package service

import (
	"github.com/drrrMikado/shorten/internal/service/shorturl"
	"github.com/google/wire"
)

var Set = wire.NewSet(New)

type Service struct {
	ShortUrl shorturl.Service
}

func New(repo shorturl.Repository) *Service {
	return &Service{
		ShortUrl: shorturl.NewService(repo),
	}
}
