package connection

import (
	"github.com/NelsonMario/middleware"
	"github.com/gorilla/mux"
)

func NewRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.LogMiddleware)

	return r
}
