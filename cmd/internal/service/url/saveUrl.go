package service

import (
	"log/slog"
	"net/http"
	"usr-short/cmd/internal/model"
	"usr-short/cmd/internal/utils"

	"github.com/go-chi/render"
)

type requestSave struct {
	Url   string `json:"url"`
	Alias string `json:"alias"`
}

type responseSave struct {
	model.Response
	Alias string `json:"alias"`
}

type urlSaver interface {
	SaveUrl(alias, longUrl string) error
}

// TODO: написать простой валидатор url
func SaveUrl(log *slog.Logger, uSaver urlSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log = log.With(slog.String("handler", "save-url"))

		var req requestSave
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("Error decode json", err)
			render.JSON(w, r, model.ERROR("Ошибка запроса"))
			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if req.Alias == "" {
			log.Info("The request did not contain an alias")
			req.Alias = utils.RandomAlias()
		}

		if req.Url == "" {
			log.Error("The request did not contain an url")
			render.JSON(w, r, model.ERROR("Пустой url"))
			return
		}

		if err = uSaver.SaveUrl(req.Alias, req.Url); err != nil {
			log.Error(err.Error())
			render.JSON(w, r, model.ERROR("Произошла ошибка при добавлении"))
			return
		}

		render.JSON(w, r, responseSave{
			Response: model.OK(),
			Alias:    req.Alias,
		})
	}
}
