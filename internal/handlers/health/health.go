package health

import (
	"errors"
	"net/http"

	"github.com/larry-a4/nftbento/internal/handlers"
	"github.com/larry-a4/nftbento/internal/repository/adapter"
	HttpStatus "github.com/larry-a4/nftbento/utils/http"
)

type Handler struct {
	handlers.Interface
	Repository adapter.Interface
}

func NewHandler(repository adapter.Interface) handlers.Interface {
	return &Handler{
		Repository: repository,
	}

}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {

	if !h.Repository.Health() {
		HttpStatus.StatusInternalServerError(w, r, errors.New("database not alive"))
	}
	HttpStatus.StatusOK(w, r, "service OK")
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r, nil)
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r, nil)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r, nil)
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusNoContent(w, r)
}
