package service

import (
	"log/slog"
	"net/http"
	"usr-short/cmd/internal/model"

	"github.com/go-chi/render"
)

type requestDeleteUrl struct {
	Alias string `json: "alias"`
}

type responseDeleteUrl struct {
	model.Response
	Url string `json: "url"`
}

type delteUrl interface {
	DeleteUrl(alias string) (string, error)
}

func DeleteUrl(log *slog.Logger, dUrl delteUrl) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log = log.With(slog.String("handler", "delte-url"))

		var req requestDeleteUrl
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

		url, err := dUrl.DeleteUrl(req.Alias)
		if err != nil {
			log.Error("Error deleting the url ", err)
			render.JSON(w, r, model.ERROR("Ошибка при удалении url"))
			return
		}

		render.JSON(w, r, responseDeleteUrl{
			Response: model.OK(),
			Url:      url,
		})
	}
}
