package router

import "github.com/gorilla/mux"

// Gerar vai retonar um router com as rotas configuradas.
func Gerar() *mux.Router {
	return mux.NewRouter()
}
