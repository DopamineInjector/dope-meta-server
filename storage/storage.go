package storage

import (
	"fmt"
	"os"
	"path"
)

func resolvePath(storagePath, id string) string {
  return path.Join(storagePath, path.Clean(id)+".png");
}

func GetFile(storagePath, id string) ([]byte, error) {
  finalPath := resolvePath(storagePath, id);
  return os.ReadFile(finalPath);
}

func AddFile(storagePath, id string, content []byte) error {
  finalPath := resolvePath(storagePath, id);
  _, err := os.ReadFile(finalPath);
  if err == nil {
    return fmt.Errorf("File already exists");
  }
  return os.WriteFile(finalPath, content, 0644);
}
