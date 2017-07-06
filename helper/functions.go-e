package helper

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"net/url"
	"sort"
	"strconv"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[r.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenerateSign(attributes map[string]interface{}, key string) string {

	sorted_keys := Ksort(attributes)

	v := url.Values{}
	for _, k := range sorted_keys {
		v.Add(k, attributes[k].(string))
	}

	// a := (url.QueryUnescape(v.Encode()))
	a := v.Encode()
	data := base64.StdEncoding.EncodeToString([]byte(a))

	//hmac ,use sha1
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))

	h_md5 := md5.New()
	io.WriteString(h_md5, string(mac.Sum(nil)))

	return fmt.Sprintf("%x", h_md5.Sum(nil))
}

func Ksort(attributes map[string]interface{}) []string {
	keys := make([]string, 0)
	for k, _ := range attributes {
		keys = append(keys, k)
	}

	// sort 'string' key in increasing order
	sort.Strings(keys)
	return keys
}

func GenerateTimeStamp() string {
	t := time.Now().UnixNano()
	shortT := strconv.FormatInt(t, 10)
	return shortT[:13]
}

func QuickRandom(length int) string {
	return RandStringRunes(length)
}
