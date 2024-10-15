package main

import (
	"dope-meta-serv/repository"
	"dope-meta-serv/routing"
	"dope-meta-serv/utils"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const DEFAULT_PORT = "5138"

func main() {
	log.Infoln("Starting dopechain nft metadata server")
	routes := routing.GetRoutes()
  db_path := utils.GetEnvWithDefaults("DB_PATH", "./nft.db");
  err := repository.InitSqliteConnection(db_path);
  if err != nil {
    log.Fatalf("Aborting. Could not establish DB connection, reason: %s", err.Error());
  }
  log.Infoln("Sucessfully established DB connection");
	port := utils.GetEnvWithDefaults("PORT", DEFAULT_PORT)
	log.Infof("Starting to listen on port %s", port)
	port = fmt.Sprintf(":%s", port)
	err = http.ListenAndServe(port, routes)
	if err != nil {
		log.Fatalf("Error while running server: %s", err.Error())
	}
}
