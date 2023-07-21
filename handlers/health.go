package handlers

import (
	"net/http"

	t "github.com/mrkhay/creative-quill-backend/models"
	"github.com/mrkhay/creative-quill-backend/utility"
)

// @Summary Show the status of server.
// @Description get the status of server.
// @Tags health
// @Accept */*
// @Produce plain
// @Success 200 "OK"
// @Router /health [get]
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) *t.ApiError {

	return utility.WriteJson(w, http.StatusOK, "Server is healthy")

}
