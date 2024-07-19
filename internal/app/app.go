package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/kikemaru/duiroPlatform/config"
	_ "github.com/kikemaru/duiroPlatform/docs"
	"github.com/kikemaru/duiroPlatform/internal/app/handler"
	"github.com/kikemaru/duiroPlatform/internal/repository"
	"github.com/kikemaru/duiroPlatform/internal/route"
	"github.com/kikemaru/duiroPlatform/internal/service"
	"github.com/kikemaru/duiroPlatform/internal/utils"
	chi_router "github.com/kikemaru/duiroPlatform/pkg/chi"
	"github.com/kikemaru/duiroPlatform/pkg/httpserver"
	"github.com/kikemaru/duiroPlatform/pkg/logger"
)

func Run() {
	logger, err := logger.NewZerologLogger()
	if err != nil {
		log.Fatalf("error init logger: %v", err)
	}

	cfg, err := config.Parse()
	if err != nil {
		panic(any("cant parse variables from config: " + err.Error()))
	}

	pg, err := repository.New(cfg.Db)
	if err != nil {
		logger.Error().Err(err).Msg("Postgres start error")
		return
	}
	defer func() {
		if err = pg.Close(); err != nil {
			logger.Error().Err(err).Msg("error close database")
		}
	}()

	tokenService := service.NewTokenService(logger)

	platformService := handler.NewPlatformService(
		logger,
		tokenService,
	)

	router := chi.NewRouter()
	handlers := route.NewRoutes(logger, router, platformService)
	utils.NewRoutes(
		handlers,
	).Setup()
	chiMux := chi_router.NewChiMux(logger)
	chiMux.Mount("/api/v1", router)

	httpServer := httpserver.New(
		chiMux,
		httpserver.Addr(cfg.MainBackendConfig.Host, cfg.MainBackendConfig.Port),
		httpserver.ReadTimeout(time.Duration(cfg.Timeout)*time.Millisecond),
		httpserver.WriteTimeout(time.Duration(cfg.Timeout)*time.Millisecond),
	)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info().Msg("app - Run - signal: " + s.String())
		logger.Error().Msg(fmt.Sprintf("app - Run - httpServer.Notify: %v", s))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		logger.Error().Msg(fmt.Sprintf("app - Run - httpServer.Shutdown: %w", err))
	}
}
