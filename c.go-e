package main

import (
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
	fmt.Println(p)

	conf, err := goconfig.LoadConfigFile(".env")

	if err != nil {
		panic(err)
	}

	m := make(map[string]interface{})

	m["client_id"] = conf.MustValue("xmly", "app_id")
	m["secret"] = conf.MustValue("xmly", "app_secret")
	m["grant_type"] = "client_credentials"
	m["device_id"] = conf.MustValue("xmly", "device_id")
	m["timestamp"] = helper.GenerateTimeStamp()
	m["nonce"] = helper.QuickRandom(16)
	m["sig"] = helper.GenerateSign(m, m["secret"].(string))
	fmt.Println(m)

	values := url.Values{}
	for k, v := range m {
		values.Set(k, v.(string))
	}

	access_token_url := "http://api.ximalaya.com/oauth2/secure_access_token"

	var accessToken models.AccessToken
	var errorInstance models.Error
	body := curl.HttpPost(access_token_url, values)
	if accessToken.JsonDecode(body).AccessToken != "" {
	} else {
		fmt.Println(errorInstance.JsonDecode(body).ErrorCode)
	}

	fmt.Println(string(curl.HttpGet("http://www.baidu.com/s", url.Values{"ie": {"UTF-8"}, "wd": {"asdf"}})))

	// p.Value("access_token", func(p *Pimple.Pimple) interface{} {
	// return &AccessToken{a}
	// })

}
