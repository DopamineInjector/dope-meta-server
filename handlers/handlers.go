package handlers

import (
	"dope-meta-serv/repository"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleTest(w http.ResponseWriter, r* http.Request) {
  w.Write([]byte("Hello"));
}

func HandlePostMetadata(w http.ResponseWriter, r* http.Request) {
  var data PostMetadataDTO;
  err := json.NewDecoder(r.Body).Decode(&data);
  if err != nil {
    http.Error(w, "Could not parse request", 422);
    return;
  }
  res, err := repository.InsertMetadata(data.Description);
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

func HandleGetMetadata(w http.ResponseWriter, r* http.Request) {
  id := r.PathValue("id");
  if id == "" {
    http.Error(w, "no id provided", 400);
    return;
  }
  meta, err := repository.GetMetadata(id);
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
