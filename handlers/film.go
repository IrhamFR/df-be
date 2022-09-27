package handlers

import (
	filmdto "dumbflix/dto/film"
	dto "dumbflix/dto/result"
	"dumbflix/models"
	"dumbflix/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerFilm struct {
	FilmRepository repositories.FilmRepository
}

var path_file = "PATH_FILE"

func HandlerFilm(FilmRepository repositories.FilmRepository) *handlerFilm {
	return &handlerFilm{FilmRepository}
}

func (h *handlerFilm) FindFilms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	films, err := h.FilmRepository.FindFilms()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// for i, p := range films {
	// 	films[i].ThumbnailFilm = os.Getenv("PATH_FILE") + p.ThumbnailFilm
	// }

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: films}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFilm) GetFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var film models.Film
	film, err := h.FilmRepository.GetFilm(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// film.ThumbnailFilm = os.Getenv("PATH_FILE") + film.ThumbnailFilm

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseFilm(film)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFilm) CreateFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("image")
	filepath := dataContex.(string)

	category_id, _ := strconv.Atoi(r.FormValue("category_id"))
	request := filmdto.CreateFilmRequest{
		Title:         r.FormValue("title"),
		ThumbnailFilm: filepath,
		Year:          r.FormValue("year"),
		CategoryID:    category_id,
		Description:   r.FormValue("description"),
		TitleEpisode:  r.FormValue("titleEpisode"),
		// ThumbnailEpisode: filename,
		LinkFilm: r.FormValue("linkfilm"),
	}

	// var categoriesId []int
	// for _, r := range r.FormValue("categoryId") {
	// 	if int(r-'0') >= 0 {
	// 		categoriesId = append(categoriesId, int(r-'0'))
	// 	}
	// }

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "dumbflix"})
	if err != nil {
		fmt.Println(err.Error())
	}

	film := models.Film{
		Title:         request.Title,
		ThumbnailFilm: resp.SecureURL,
		Year:          request.Year,
		CategoryID:    request.CategoryID,
		Description:   request.Description,
		TitleEpisode:  request.TitleEpisode,
		// ThumbnailEpisode: filename,
		LinkFilm: request.LinkFilm,
	}

	film, err = h.FilmRepository.CreateFilm(film)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	film, _ = h.FilmRepository.GetFilm(film.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: film}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerFilm) UpdateFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(filmdto.UpdateFilmRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	film, err := h.FilmRepository.GetFilm(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Title != "" {
		film.Title = request.Title
	}

	if request.ThumbnailFilm != "" {
		film.ThumbnailFilm = request.ThumbnailFilm
	}

	if request.Year != "" {
		film.Year = request.Year
	}

	if request.TitleEpisode != "" {
		film.TitleEpisode = request.TitleEpisode
	}

	// if request.ThumbnailEpisode != "" {
	// 	film.ThumbnailEpisode = request.ThumbnailEpisode
	// }

	if request.Description != "" {
		film.Description = request.Description

		data, err := h.FilmRepository.UpdateFilm(film)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseFilm(data)}
		json.NewEncoder(w).Encode(response)
	}

}

func (h *handlerFilm) DeleteFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ContentType", "application/json")

	// Get Product ID
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	film, err := h.FilmRepository.GetFilm(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.FilmRepository.DeleteFilm(film)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseFilm(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseFilm(u models.Film) models.FilmResponse {
	return models.FilmResponse{
		ID:            u.ID,
		Title:         u.Title,
		ThumbnailFilm: u.ThumbnailFilm,
		Year:          u.Year,
		Category:      u.Category,
		Description:   u.Description,
		TitleEpisode:  u.TitleEpisode,
		LinkFilm:      u.LinkFilm,
		// ThumbnailEpisode: r.ThumbnailEpisode,
	}
}
