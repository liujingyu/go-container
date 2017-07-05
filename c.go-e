package main

import "fmt"
// import "github.com/interactiv/expect"
import "github.com/liujingyu/pimple"
import "./helper"
import "time"

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

    fmt.Println(helper.RandStringRunes(4))
    t := time.Now().UnixNano()
    fmt.Println(t)
}
