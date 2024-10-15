package utils

import "os"

func GetEnvWithDefaults(key, defaultValue string) string {
	res, ok := os.LookupEnv(key);
  if !ok {
    res = defaultValue
  }
  return res
}
