package router

import (
	"github.com/gorilla/mux"
)

var HeadNodeRouter *mux.Router

// func InitRouter() *mux.Router {
// 	return mux.NewRouter()
// }

/**
*
*  initilize teh router
*
**/
func InitRouter() {
	HeadNodeRouter = mux.NewRouter()
}

/**
*
* create the subRouter
*
**/
func SubRouter(subRouterPath string) *mux.Router {
	return HeadNodeRouter.PathPrefix(subRouterPath).Subrouter()
}
