package service

import "github.com/rs/zerolog"

type TokenService interface {
	GetTempToken()
}

type tokenService struct {
	log *zerolog.Logger
}

func NewTokenService(log *zerolog.Logger) TokenService {
	return &tokenService{
		log: log,
	}
}

func (s *tokenService) GetTempToken() {
}
