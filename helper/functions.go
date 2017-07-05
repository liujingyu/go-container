package helper

import "time"
import "math/rand"

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
