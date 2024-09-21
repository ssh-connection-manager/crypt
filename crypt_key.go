package crypt

import (
	"crypto/rand"
	"errors"

	"github.com/ssh-connection-manager/file"
)

func GetKey(pathConf string, fileNameKey string) ([]byte, error) {
	data, err := file.ReadFile(file.GetFullPath(pathConf, fileNameKey))
	if err != nil {
		return []byte(data), errors.New("empty name")
	}

	return []byte(data), nil
}

func GenerateKey(pathConf string, fileNameKey string) error {
	filePath := file.GetFullPath(pathConf, fileNameKey)

	if !file.IsExistFile(filePath) {
		file.CreateFile(filePath)

		cryptKey, err := GetKey(pathConf, fileNameKey)
		if err != nil {
			return errors.New("empty name")
		}

		if len(cryptKey) == 0 {
			key := make([]byte, 32)

			_, err := rand.Read(key)
			if err != nil {
				return errors.New("key generation error")
			}

			err = file.WriteFile(file.GetFullPath(pathConf, fileNameKey), key)
			if err != nil {
				return errors.New("error writing key")
			}
		}
	}

	return nil
}
