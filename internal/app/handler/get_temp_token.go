package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kikemaru/duiroPlatform/internal/utils"
)

// GetTempToken получение временного токена
// @Summary Получение временного токена
// @Tags PlatformService
// @Description Получение временного токена для бота по ключу. Жизнь токена ограничена 24 часами.
// @Accept json
// @Produce json
// @Param app_key path string true "key bot application"
// @Success 200 {integer} integer
// @Failure 401 {integer} integer
// @Failure 500 {object} errors_handler.ErrorResponse
// @Router /v1/{app_key}/temp_token [get]
func (i *Implementation) GetTempToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := chi.URLParam(r, "app_key")
		i.tokenService.GetTempToken()
		if err := utils.Json(w, http.StatusOK, key); err != nil {
			i.Log.Error().Err(err)
		}
	}
}
