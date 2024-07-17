package handler

import "github.com/rs/zerolog"

type Implementation struct {
	Log *zerolog.Logger
	HandleInterface
}

func NewPlatformService(log *zerolog.Logger) *Implementation {
	return &Implementation{
		Log: log,
	}
}
