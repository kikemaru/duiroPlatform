package route

import (
	"github.com/go-chi/chi"
	"github.com/kikemaru/duiroPlatform/internal/app/handler"
	"github.com/rs/zerolog"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Routes struct {
	log *zerolog.Logger
	m   *chi.Mux
	i   *handler.Implementation
}

func NewRoutes(l *zerolog.Logger, m *chi.Mux, i *handler.Implementation) *Routes {
	return &Routes{
		log: l,
		m:   m,
		i:   i,
	}
}

func (route *Routes) Setup() {

	route.m.Group(func(r chi.Router) {
		{
			route.log.Info().Msg("setting up documentation routes")
			r.Mount("/swagger", httpSwagger.WrapHandler)
		}

		{
			route.log.Info().Msg("setting up handlers routes")
			r.Get("/test1", route.i.Test())
		}
	})

}
