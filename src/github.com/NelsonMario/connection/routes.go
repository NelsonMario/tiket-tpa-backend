package connection

import (
	middleware2 "github.com/NelsonMario/middleware"
	"github.com/gorilla/mux"
)

func newRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware2.LogMiddleware)

	return r
}
