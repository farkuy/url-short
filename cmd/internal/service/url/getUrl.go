package service

import (
	"log/slog"
	"net/http"
	"usr-short/cmd/internal/model"

	"github.com/go-chi/render"
)

type requestUrlGet struct {
	Alias string `json:"alias"`
}

type responseUrlGet struct {
	model.Response
	Url string `json:"url"`
}

type getUrl interface {
	GetUrl(alias string) (string, error)
}

func GetUrl(log *slog.Logger, uGet getUrl) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log = log.With(slog.String("handler", "get-url"))

		var req requestUrlGet
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("Error decode json", err)
			render.JSON(w, r, model.ERROR("Ошибка запроса"))
			return
		}

		if req.Alias == "" {
			log.Error("The request did not contain an alias")
			render.JSON(w, r, model.ERROR("Пустой alias"))
			return
		}

		url, err := uGet.GetUrl(req.Alias)
		if err != nil {
			log.Error(err.Error())
			render.JSON(w, r, model.ERROR("Произошла ошибка при получении данных"))
			return
		}

		render.JSON(w, r, responseUrlGet{
			Response: model.OK(),
			Url:      url,
		})
	}
}
