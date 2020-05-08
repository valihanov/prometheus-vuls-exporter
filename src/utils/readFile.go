package utils

import (
	"io/ioutil"
	"log"
)

func ReadFile(path string) []byte {
	var data, err = ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
