package http

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/heptiolabs/healthcheck"
	"go.uber.org/zap"
)

type SynonymsManager interface {
	GetSynonyms(ctx context.Context, word string) ([]string, error)
	AddSynonyms(ctx context.Context, word, synonym string) error
}

func Routes(logger *zap.Logger, synonymsManager SynonymsManager) *chi.Mux {
	// handlers
	r := chi.NewRouter()

	// healthcheck & readiness
	health := healthcheck.NewHandler()
	healtcheckLogger := logger.Named("healthcheck")
	health.AddLivenessCheck("alive", func() error {
		healtcheckLogger.Info("alive")
		return nil
	})
	health.AddReadinessCheck("ready", func() error {
		healtcheckLogger.Info("ready")
		return nil
	})

	r.Get("/live", health.LiveEndpoint)
	r.Get("/ready", health.ReadyEndpoint)

	// api
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/word", GetSynonymsHandler(logger, synonymsManager))
		r.Post("/synonyms", PostSynonymsHandler(logger, synonymsManager))
	})

	return r
}
