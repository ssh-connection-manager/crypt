package crypt

import "errors"

func Decrypt(str string, pathConf string, NameFIleKey string) (string, error) {
	cryptKey, err := GetKey(pathConf, NameFIleKey)
	if err != nil {
		return "", errors.New("err at get crypt key")
	}

	decrypted, err := decrypt(cryptKey, str)
	if err != nil {
		return "", errors.New("encryption error")
	}

	return decrypted, nil
}