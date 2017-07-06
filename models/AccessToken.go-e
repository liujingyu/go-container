package models

import (
	"encoding/json"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`

    Params map[string]string
}

func (t AccessToken) JsonDecode(token []byte) AccessToken {

	err := json.Unmarshal(token, &t)

	if err != nil {
		panic(err)
	}

	return t
}

