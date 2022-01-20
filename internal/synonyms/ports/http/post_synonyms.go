package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"synonyms/internal/synonyms"

	"go.uber.org/zap"
)

type PostSynonymsRequest struct {
	Word    string `json:"word"`
	Synonym string `json:"synonym"`
}

func PostSynonymsHandler(logger *zap.Logger, manager SynonymsManager,
) func(w http.ResponseWriter, r *http.Request) {

	logger = logger.Named("post_synonyms")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger.Info("request received", zap.String("method", r.Method), zap.String("url", r.URL.String()))

		var request PostSynonymsRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			logger.Error("failed to decode request", zap.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err := manager.AddSynonyms(ctx, request.Word, request.Synonym)
		switch {
		case err == nil:
		// ok
		case errors.Is(err, context.Canceled):
			logger.Info("request cancelled", zap.Error(err))
			return
		case errors.Is(err, context.DeadlineExceeded):
			logger.Warn("request timed out", zap.Error(err))
			return
		case errors.Is(err, synonyms.ErrValiadation):
			logger.Info("request body is not valid", zap.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			logger.Error("unexpected error", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		logger.Info("request processed", zap.String("method", r.Method), zap.String("url", r.URL.String()))
	}
}
