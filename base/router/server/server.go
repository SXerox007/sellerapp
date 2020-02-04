package server

import (
	"net/http"

	"sellerapp/base/router"

	"github.com/gorilla/handlers"
)

/**
*
* start the server
*
**/
func StartServer(port string) {
	headersOk := handlers.AllowedHeaders([]string{""})
	http.ListenAndServe("localhost:"+port, handlers.CORS(headersOk)(router.HeadNodeRouter))
}
