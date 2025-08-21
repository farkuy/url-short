package service

import (
	"log/slog"
	"net/http"
	"usr-short/cmd/internal/model"

	"github.com/go-chi/render"
)

type requestUpdateUrl struct {
	NewUrl string `json:"newUrl"`
	Alias  string `json:"alias"`
}

type responseUpdateUrl struct {
	model.Response
	NewUrl string `json: "newUrl"`
}

type updateUrl interface {
	UpdateUrl(alias, newUrl string) error
}

// TODO: логика если заменяют на тот же самый url
// TODO: логика проверки существования alias который обновляют
func UpdateUrl(log *slog.Logger, uUrl updateUrl) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log = log.With(slog.String("handler", "update-url"))

		var req requestUpdateUrl
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("Error decode json", err)
			render.JSON(w, r, model.ERROR("Ошибка при декодировании данных"))
			return
		}

		if req.Alias == "" {
			log.Error("Empty alias", err)
			render.JSON(w, r, model.ERROR("Передано пустое поле alias"))
			return
		}

		if req.NewUrl == "" {
			log.Error("Empty newUrl", err)
			render.JSON(w, r, model.ERROR("Передано пустое поле newUrl"))
			return
		}

		err = uUrl.UpdateUrl(req.Alias, req.NewUrl)
		if err != nil {
			log.Error("Error updating the url ", err)
			render.JSON(w, r, model.ERROR("Ошибка при обновлении url"))
			return
		}

		render.JSON(w, r, responseUpdateUrl{
			Response: model.OK(),
			NewUrl:   req.NewUrl,
		})
	}
}
