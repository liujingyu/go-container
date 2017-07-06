package main

import (
	"./helper"
	// "reflect"
	"fmt"
	"github.com/interactiv/pimple"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	type Foo struct {
		baz int
	}
	type Bar struct {
		foo *Foo
	}
	type AccessToken struct {
		access_token string
		expires_in   int
	}
	p := pimple.New(map[string]func(*pimple.Pimple) interface{}{
		"foo": func(p *pimple.Pimple) interface{} {
			return &Foo{baz: 1}
		},
		"bar": func(p *pimple.Pimple) interface{} {
			return &Bar{foo: p.Get("foo").(*Foo)}
		},
	})
	bar := p.Get("bar").(*Bar)
	p.Value("biz", "a")
	p.Set("buzz", func(p *pimple.Pimple) interface{} {
		return &Buzz{string: p.Get("biz").(string)}
	})
	p.Extend("buzz", func(buzz interface{}, p *pimple.Pimple) interface{} {
		buzz.(*Buzz).number = 23
		return buzz
	})

	fmt.Println(p.Get("biz").(string))
	fmt.Println(bar.foo.baz)
	fmt.Println(p.Get("biz").(string))

	app_id := "818d02450a10602d98634ab181b350ac"
	m := make(map[string]interface{})

	m["client_id"] = app_id
	m["secret"] = secret
	m["grant_type"] = "client_credentials"
	m["device_id"] = "sparkle-code-cardon-denim.local"
	m["timestamp"] = helper.GenerateTimeStamp()
	m["nonce"] = helper.QuickRandom(16)
	m["sig"] = helper.GenerateSign(m, secret)

	fmt.Println(m)

	values := url.Values{}
	for k, v := range m {
		values.Set(k, v.(string))
	}
	fmt.Println(values.Encode())

	access_token_url := "http://api.ximalaya.com/oauth2/secure_access_token"
	resp, err := http.PostForm(access_token_url, values)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("post:\n", string(body))

	// p.Value("access_token", func(p *Pimple.Pimple) interface{} {
	// return &AccessToken{a}
	// })

}
