package handlers

import (
	"database/sql"
	"dope-meta-serv/repository"
	"dope-meta-serv/storage"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func HandleTest(w http.ResponseWriter, r* http.Request) {
  w.Write([]byte("Hello"));
}

func HandlePostMetadata(w http.ResponseWriter, r* http.Request, db *sql.DB) {
  var data PostMetadataDTO;
  err := json.NewDecoder(r.Body).Decode(&data);
  if err != nil {
    http.Error(w, "Could not parse request", 422);
    return;
  }
  res, err := repository.InsertMetadata(db, data.Description);
  if err != nil {
    http.Error(w, "Error while operating on db", 500);
    return
  }
  serialized, err := json.Marshal(res);
  if err != nil {
    http.Error(w, "Error parsind db response", 500);
    return
  }
  w.Header().Add("content-type", "application/json");
  w.WriteHeader(http.StatusCreated);
  w.Write(serialized);
}

func HandleGetMetadata(w http.ResponseWriter, r* http.Request, db *sql.DB) {
  id := r.PathValue("id");
  if id == "" {
    http.Error(w, "no id provided", 400);
    return;
  }
  if _, err := uuid.Parse(id); err != nil {
    http.Error(w, "invalid uuid format", 400);
    return;
  }
  meta, err := repository.GetMetadata(db, id);
  if err != nil {
    http.Error(w, fmt.Sprintf("error while operating on db, %s", err.Error()), 500);
    return;
  }
  serialized, err := json.Marshal(meta);
  if err != nil {
    http.Error(w, "Error parsing db response", 500);
    return
  }
  w.Header().Add("content-type", "application/json");
  w.WriteHeader(http.StatusOK);
  w.Write(serialized);
}

func HandlePostImage(w http.ResponseWriter, r* http.Request, storagePath string) {
  id := r.PathValue("id");
  if id == "" {
    http.Error(w, "no id provided", 400);
    return;
  }
  if _, err := uuid.Parse(id); err != nil {
    http.Error(w, "invalid uuid format", 400);
    return;
  }
  log.Println(id)
  body, err := io.ReadAll(r.Body);
  if err != nil {
    http.Error(w, "invalid request body", 400);
    return;
  }
  err = storage.AddFile(storagePath, id, body);
  if err != nil {
    http.Error(w, "file with this id already exists", 409);
    return;
  }
  w.WriteHeader(http.StatusCreated);
}

func HandleGetImage(w http.ResponseWriter, r* http.Request, storagePath string) {
  id := r.PathValue("id");
  if id == "" {
    http.Error(w, "no id provided", 400);
    return;
  }
  if _, err := uuid.Parse(id); err != nil {
    http.Error(w, "invalid uuid format", 400);
    return;
  }
  content, err := storage.GetFile(storagePath, id);
  if err != nil {
    http.Error(w, "file does not exist", 404);
    return;
  }
  w.Header().Add("content-type", "image/png");
  w.WriteHeader(http.StatusOK);
  w.Write(content);
}
