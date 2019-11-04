package routes

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/kudejwiktor/go-api-example/src/User/domain"
	//"github.com/pkg/errors"
	"go-api-example/app/http/middleware"

	//"go-api-example/source/User/infrastructure/persistence"
	"net/http"
	"strconv"
)

type UserRouter struct {
	router     *chi.Mux
	repository domain.UserRepository
}

func NewRouter(router *chi.Mux, repository domain.UserRepository) *UserRouter {
	return &UserRouter{
		router:     router,
		repository: repository,
	}
}

func (router *UserRouter) Routes() {
	router.router.Get("/rest/banks/{id:[0-9]+}", middleware.CommonHeaders(router.User()))
}

func (router UserRouter) User() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			fmt.Println(err)
			//handleErrors(w, errors.Wrap(err, http.StatusText(http.StatusBadRequest)))
			return
		}
		b, err := router.repository.GetUserOfId(id)
		if err != nil {
			fmt.Println(err)
			//handleErrors(w, err)
			return
		}
		if err := json.NewEncoder(w).Encode(b); err != nil {
			fmt.Println(err)
			//handleErrors(w, err)
			return
		}
	}
}
