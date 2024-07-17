package handler

import (
	"net/http"

	"github.com/kikemaru/duiroPlatform/internal/utils"
)

// Test
// @Summary Test handler
// @Tags PlatformService
// @Description
// @Accept json
// @Produce json
// @Success 200 {integer} integer
// @Failure 401 {integer} integer
// @Failure 500 {object} errors_handler.ErrorResponse
// @Router /v1/test1 [get]
func (i *Implementation) Test() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := utils.Json(w, http.StatusOK, nil); err != nil {
			i.Log.Error().Err(err)
		}
	}
}
