package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"synonyms/internal/synonyms"

	"go.uber.org/zap"
)

type GetSynonymsResponse struct {
	Word     string   `json:"word"`
	Synonyms []string `json:"synonyms"`
}

func GetSynonymsHandler(logger *zap.Logger, manager SynonymsManager,
) func(w http.ResponseWriter, r *http.Request) {

	logger = logger.Named("get_synonyms")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger.Info("request received", zap.String("method", r.Method), zap.String("url", r.URL.String()))

		word := r.URL.Query().Get("word")

		ss, err := manager.GetSynonyms(ctx, word)
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
			logger.Info("request input is not valid", zap.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, synonyms.ErrNotFound):
			logger.Info("request data is not valid", zap.Error(err))
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			logger.Error("unexpected error", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := GetSynonymsResponse{
			Word:     word,
			Synonyms: ss,
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(&resp)
		if err != nil {
			logger.Error("failed to encode response", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logger.Info("request processed", zap.String("method", r.Method), zap.String("url", r.URL.String()))
	}
}
