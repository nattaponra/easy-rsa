package rsa

import (
	"encoding/pem"
	"io/ioutil"
)

func GetPublicKeyFromPEMFile(path string) error {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
 
	b pem.Decode(dat)

	if err != nil {
		return err
	}
	return nil
}
