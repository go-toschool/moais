package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getHTTP(url string, data interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return errors.New("invalid status code")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Printf("%s", body)

	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	return nil
}
