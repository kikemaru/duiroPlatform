package handler

import (
	"github.com/kikemaru/duiroPlatform/internal/service"
	"github.com/rs/zerolog"
)

type Implementation struct {
	Log          *zerolog.Logger
	tokenService service.TokenService
	HandleInterface
}

func NewPlatformService(log *zerolog.Logger, tokenService service.TokenService) *Implementation {
	return &Implementation{
		Log:          log,
		tokenService: tokenService,
	}
}
