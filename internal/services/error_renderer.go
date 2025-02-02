package services

import (
	"fmt"
	"net/http"

	"github.com/RHEnVision/provisioning-backend/internal/ctxval"
	"github.com/go-chi/render"
)

// writeBasicError is used when rendering of the error fails so at least something is written
func writeBasicError(w http.ResponseWriter, r *http.Request, err error) {
	if logger := ctxval.GetLogger(r.Context()); logger != nil {
		logger.Error().Msgf("unable to render error %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(508)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"msg": "%s"}`, err.Error())))
}

func renderError(w http.ResponseWriter, r *http.Request, renderer render.Renderer) {
	errRender := render.Render(w, r, renderer)
	if errRender != nil {
		writeBasicError(w, r, errRender)
	}
}
