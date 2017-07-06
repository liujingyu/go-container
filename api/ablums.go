package api

import "../curl"
import "net/url"
import "../helper"

const API_BROWSE = "http://api.ximalaya.com/albums/browse"

type Album struct {
	AbstractAPI
}

func (album Album) BrowseAPI(album_id string) []byte {
	values := url.Values{}

	params := map[string]string{"client_os_type": "3", "app_key": album.AccessToken.Params["client_id"], "device_id": album.AccessToken.Params["device_id"]}

	params["album_id"] = album_id

	params["access_token"] = album.AccessToken.AccessToken

	for k, v := range params {
        values.Add(k, v)
	}

	values.Add("sig", helper.GenerateSign(params, album.AccessToken.Params["secret"]))

    return curl.HttpGet(API_BROWSE, values)
}
