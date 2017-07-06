package main

import (
	"./helper"
    "crypto/hmac"
    "crypto/sha1"
    // "reflect"
    "crypto/md5"
    "encoding/base64"
	"fmt"
	"github.com/liujingyu/pimple"
    "io"
	"net/url"
	"sort"
	"strconv"
	"time"
)

func main() {
	type Foo struct {
		baz int
	}
	type Bar struct {
		foo *Foo
	}
	type Buzz struct {
		string string
		number int
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

	fmt.Println(quickRandom(15))
	shortT2 := generateTimeStamp()

	fmt.Println(shortT2)

	m := make(map[string]interface{})
	m["a"] = "1"
	m["b"] = "2"

	generateSign(m, "123456")

	// fmt.Println(url.QueryUnescape(v.Encode()))
	// msg := "1"
	// encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	// fmt.Println(encoded)

	// //sha1
	// h := sha1.New()
	// io.WriteString(h, "aaaaaa")
	// fmt.Printf("%x\n", h.Sum(nil))

	// //hmac ,use sha1
	// key := []byte("123456")
	// mac := hmac.New(sha1.New, key)
	// mac.Write([]byte("YT0xJmI9Mg=="))
	// fmt.Println(string(mac.Sum(nil)))
	// // fmt.Printf("%x\n", mac.Sum(nil))

	// h_md5 := md5.New()
	// io.WriteString(h_md5, string(mac.Sum(nil)))
	// fmt.Printf("%x", h_md5.Sum(nil))
}

func generateSign(attributes map[string]interface{}, key string) {

	sorted_keys := ksort(attributes)

	v := url.Values{}
	for _, k := range sorted_keys {
		v.Add(k, attributes[k].(string))
	}

	// a := (url.QueryUnescape(v.Encode()))
	a := v.Encode()
	// fmt.Println(reflect.TypeOf(a))
    data := base64.StdEncoding.EncodeToString([]byte(a))
    fmt.Println(data)

    //sha1
    h_sha1 := sha1.New()
    io.WriteString(h_sha1, data)

    //hmac ,use sha1
    mac := hmac.New(sha1.New, []byte(key))
    mac.Write([]byte(string(h_sha1.Sum(nil))))

    h_md5 := md5.New()
    io.WriteString(h_md5, string(mac.Sum(nil)))
    fmt.Printf("%x", h_md5.Sum(nil))
}

func ksort(attributes map[string]interface{}) []string {
	keys := make([]string, 0)
	for k, _ := range attributes {
		keys = append(keys, k)
	}
	// sort 'string' key in increasing order
	sort.Strings(keys)
	return keys
	// v := url.Values{}

	// for _, k := range sorted_keys {
	// v.Add(k, m[k].(string))
	// }
}

func generateTimeStamp() string {
	t := time.Now().UnixNano()
	shortT := strconv.FormatInt(t, 10)
	shortT2 := shortT[:13]
	return shortT2
}

func quickRandom(length int) string {

	return helper.RandStringRunes(length)
}
