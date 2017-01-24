package routers

import (
	"github.com/colombia9503/RESTful-SqlServer/common"
	"github.com/colombia9503/RESTful-SqlServer/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func SetCostosRouters(router *mux.Router) *mux.Router {
	costosRouter := mux.NewRouter()
	costosRouter.HandleFunc("/api/costos", controllers.Costos.Get).Methods("GET")
	costosRouter.HandleFunc("/api/costos/{id}", controllers.Costos.GetOne).Methods("GET")
	costosRouter.HandleFunc("/api/costos/exists/{id}", controllers.Costos.Exists).Methods("GET")
	costosRouter.HandleFunc("/api/costos", controllers.Costos.Insert).Methods("POST")
	router.PathPrefix("/api/costos").Handler(
		negroni.New(
			negroni.HandlerFunc(common.Authorize),
			negroni.Wrap(costosRouter),
		))
	return router
}
