package main

import (
	"./api"
	"./curl"
	"./helper"
	"./models"
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/interactiv/pimple"
	"net/url"
)

func main() {
	type Foo struct {
		baz int
	}
	type Bar struct {
		foo *Foo
	}
	p := pimple.New(map[string]func(*pimple.Pimple) interface{}{
		"foo": func(p *pimple.Pimple) interface{} {
			return &Foo{baz: 1}
		},
		"bar": func(p *pimple.Pimple) interface{} {
			return &Bar{foo: p.Get("foo").(*Foo)}
		},
	})

	conf, err := goconfig.LoadConfigFile(".env")

	if err != nil {
		panic(err)
	}

	m := make(map[string]string)

	m["client_id"] = conf.MustValue("xmly", "app_id")
	m["secret"] = conf.MustValue("xmly", "app_secret")
	m["grant_type"] = "client_credentials"
	m["device_id"] = conf.MustValue("xmly", "device_id")
	m["timestamp"] = helper.GenerateTimeStamp()
	m["nonce"] = helper.QuickRandom(16)
	m["sig"] = helper.GenerateSign(m, conf.MustValue("xmly", "app_secret"))

	values := url.Values{}
	for k, v := range m {
		values.Set(k, v)
	}

	access_token_url := "http://api.ximalaya.com/oauth2/secure_access_token"

	accessToken := &models.AccessToken{}
	var accessToken2 models.AccessToken
	var errorInstance models.Error
	body := curl.HttpPost(access_token_url, values)
	albumIds := [...]string{"216726", "2950984", "259590", "3018431", "2878121", "359012", "271639", "235015", "352108", "289695", "2859342", "270353", "260744", "291274", "365909", "231461", "299880", "318324", "286034", "2814226"}
	if accessToken2 = accessToken.JsonDecode(body); accessToken2.AccessToken != "" {
		accessToken2.Params = m
		p.Value("access_token", accessToken2)
		album := api.Album{}
		album.NewAPI(p)

		for _, albumId := range albumIds {
			go fmt.Println(string(album.BrowseAPI(albumId)))
		}

	} else {
		fmt.Println(errorInstance.JsonDecode(body).ErrorCode)
	}

}
