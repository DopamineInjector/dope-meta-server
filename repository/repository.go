package repository

import (
	"database/sql"

	"github.com/google/uuid"
  _ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

type NftMetadataEntry struct {
  Id string `json:"id"`
  Description string `json:"description"`
  ImageId string `json:"imageId"`
}

var db *sql.DB;

func InitSqliteConnection(dbPath string) error {
  conn, err := sql.Open("sqlite3", dbPath);
  if err != nil {
    return err
  }
  db = conn
  _, err = db.Exec("create table if not exists metadata (id text not null primary key, description text, imageId text)");
  if err != nil {
    return err
  }
  return nil
}

func InsertMetadata(description string) (*NftMetadataEntry, error) {
  if db == nil {
    log.Panicln("Accessed db without initializing connection");
  }
  metaId := uuid.NewString();
  imageId := uuid.NewString();
  entry := NftMetadataEntry {
    Id: metaId,
    Description: description,
    ImageId: imageId,
  };
  stmt, err := db.Prepare("insert into metadata values (?, ?, ?)");
  if err != nil {
    return nil, err
  }
  defer stmt.Close();
  _, err = stmt.Exec(entry.Id, entry.Description, entry.ImageId);
  if err != nil {
    return nil, err
  }
  return &entry, nil
}

func GetMetadata(id string) (*NftMetadataEntry, error) {
  if db == nil {
    log.Panicln("Accessed db without initializing connection");
  }
  query, err := db.Prepare("select * from metadata where id = ?")
  if err != nil {
    return nil, err
  }
  defer query.Close();
  var dbId string;
  var description string;
  var imageId string;
  err = query.QueryRow(id).Scan(&dbId, &description, &imageId);
  if err != nil {
    return nil, err
  }
  if dbId == "" {
    return nil, nil
  }
  metadata := NftMetadataEntry{
    Id: dbId,
    Description: description,
    ImageId: imageId,
  }
  return &metadata, nil
}
