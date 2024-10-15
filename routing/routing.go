package routing

import (
	"database/sql"
	"dope-meta-serv/handlers"
	"net/http"
)

func GetRoutes(db *sql.DB, storagePath string) *http.ServeMux {
  mux := http.DefaultServeMux;
  mux.Handle("GET /", http.HandlerFunc(handlers.HandleTest));
  mux.Handle("POST /metadata", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    handlers.HandlePostMetadata(w, r, db);
  }))
  mux.Handle("GET /metadata/{id}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    handlers.HandleGetMetadata(w, r, db);
  }))
  mux.Handle("GET /avatars/{id}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    handlers.HandleGetImage(w, r, storagePath);
  }))
  mux.Handle("POST /avatars/{id}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    handlers.HandlePostImage(w, r, storagePath);
  }))
  return mux
}
