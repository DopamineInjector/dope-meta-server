package routing

import (
	"dope-meta-serv/handlers"
	"net/http"
)

func GetRoutes() *http.ServeMux {
  mux := http.DefaultServeMux;
  mux.Handle("GET /", http.HandlerFunc(handlers.HandleTest));
  mux.Handle("POST /metadata", http.HandlerFunc(handlers.HandlePostMetadata))
  mux.Handle("GET /metadata/{id}", http.HandlerFunc(handlers.HandleGetMetadata))
  return mux
}
