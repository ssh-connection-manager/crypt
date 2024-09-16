package crypt

import (
	"crypto/rand"
	"errors"

	"github.com/ssh-connection-manager/file"
)

func GetKey(pathConf string, NameFIleKey string) ([]byte, error) {
	data, err := file.ReadFile(file.GetFullPath(pathConf, NameFIleKey))
	if err != nil {
		return []byte(data), errors.New("empty name")
	}

	return []byte(data), nil
}

func GenerateKey(pathConf string, NameFIleKey string) error {
	filePath := file.GetFullPath(pathConf, NameFIleKey)

	if !file.IsExistFile(filePath) {
		file.CreateFile(filePath)

		cryptKey, err := GetKey(pathConf, NameFIleKey)
		if err != nil {
			return errors.New("empty name")
		}

		if len(cryptKey) == 0 {
			key := make([]byte, 32)

			_, err := rand.Read(key)
			if err != nil {
				return errors.New("key generation error")
			}

			err = file.WriteFile(file.GetFullPath(pathConf, NameFIleKey), key)
			if err != nil {
				return errors.New("error writing key")
			}
		}
	}

	return nil
}
