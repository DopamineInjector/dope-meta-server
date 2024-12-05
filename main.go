package main

import (
	"dope-meta-serv/repository"
	"dope-meta-serv/routing"
	"dope-meta-serv/utils"
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

const DEFAULT_PORT = "5138"
const DEFAULT_DB = "./nft.db"
const DEFAULT_STORAGE = "./data/"

func main() {
	log.Infoln("Starting dopechain nft metadata server")
  storagePath := utils.GetEnvWithDefaults("STORAGE_PATH", DEFAULT_STORAGE);
  os.Mkdir(storagePath, 0755);
  db_path := utils.GetEnvWithDefaults("DB_PATH", DEFAULT_DB);
  db, err := repository.InitSqliteConnection(db_path);
  if err != nil {
    log.Fatalf("Aborting. Could not establish DB connection, reason: %s", err.Error());
  }
  log.Infoln("Sucessfully established DB connection");
	port := utils.GetEnvWithDefaults("PORT", DEFAULT_PORT)
	log.Infof("Starting to listen on port %s", port)
	port = fmt.Sprintf(":%s", port)
	routes := routing.GetRoutes(db, storagePath);
	err = http.ListenAndServe(port, routes)
	if err != nil {
		log.Fatalf("Error while running server: %s", err.Error())
	}
}
