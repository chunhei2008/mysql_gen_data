package main

import (
	"bytes"
	"math/rand"
	"strings"
	"time"
)

type Gen struct {
	seed chan int64 //统一的随机种子生成器
}

func NewGen() *Gen {
	return &Gen{func() chan int64 {
		ch := make(chan int64, 1024)
		go func() {
			var i int64 = 0
			for ; ; i++ {
				ch <- time.Now().UnixNano() + i
			}
		}()
		return ch
	}()}
}

//区间随机数生成器
func (g *Gen) gen_varint(min, max int) chan interface{} {
	ch := make(chan interface{}, 100)
	go func() {
		if max < min || max <= 0 || min < 0 {
			ch <- 0
		} else {
			length := max - min
			for i := 0; ; i++ {
				rand.Seed(<-g.seed)
				rands := min + rand.Intn(length)
				ch <- rands
			}
		}
	}()
	return ch
}

//区间随机数生成器
func (g *Gen) gen_int(max int) chan interface{} {
	ch := make(chan interface{}, 100)
	go func() {
		if max <= 0 {
			for i := 0; ; i++ {
				ch <- 0
			}
		} else {
			for i := 0; ; i++ {
				rand.Seed(<-g.seed)
				rands := rand.Intn(max)
				ch <- rands
			}
		}
	}()
	return ch
}

//定长字符生成器
func (g *Gen) gen_char(length int) chan interface{} {
	ch := make(chan interface{})
	go func() {
		if length <= 0 {
			for i := 0; ; i++ {
				ch <- ""
			}
		} else {
			str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
			if re := length / len(str); re > 0 {
				str = strings.Repeat(str, re+1)
			}
			for i := 0; ; i++ {
				buffer := bytes.NewBuffer([]byte{})
				rand.Seed(<-g.seed)
				rands := rand.Perm(len(str))
				for i := 0; i < length; i++ {
					buffer.WriteString(str[rands[i] : rands[i]+1])
				}
				ch <- buffer.String()
			}
		}
	}()
	return ch
}

//变长字符生成器
func (g *Gen) gen_varchar(min, max int) chan interface{} {
	ch := make(chan interface{})
	go func() {
		if max < min || min <= 0 || max <= 0 {
			for i := 0; ; i++ {
				ch <- ""
			}
		} else {
			str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
			if re := max / len(str); re > 0 {
				str = strings.Repeat(str, re+1)
			}
			span := max - min
			for i := 0; ; i++ {
				randstr := bytes.NewBuffer([]byte{})
				rand.Seed(<-g.seed)
				seed := rand.Intn(span)
				rands := rand.Perm(len(str))
				for i := 0; i < min+seed; i++ {
					randstr.WriteString(str[rands[i] : rands[i]+1])
				}
				ch <- randstr.String()
			}
		}
	}()
	return ch
}

//以基数自增
func (g *Gen) gen_autoincr(base int) chan interface{} {
	ch := make(chan interface{}, 1024)
	go func() {
		if base < 0 {
			base = 0
		}
		i := base
		for ; ; i++ {
			ch <- i
		}
	}()
	return ch
}
