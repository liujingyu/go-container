package curl

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func HttpPost(httpurl string, values url.Values) []byte {

	resp, err := http.PostForm(httpurl, values)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return body
}

func HttpGet(httpurl string, values url.Values) []byte {
	resp, err := http.Get(httpurl +"?"+ values.Encode())
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body
}
