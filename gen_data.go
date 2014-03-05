package main

import (
	"flag"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

var rows = flag.Int64("n", 10000, "row number")
var format = flag.String("f", "", "row format %m[-n]d,%m[-n]s,%[m]n")

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
	gen := NewGen()
	if len(*format) == 0 {
		fmt.Println("data format error")
		return
	}
	strs := strings.Split(*format, ",")
	contain := make([]chan interface{}, len(strs))
	fmtFormat := make([]string, len(strs))
	for idx, str := range strs {
		strlen := len(str)
		strtype := str[strlen-1 : strlen]
		limit := str[1 : strlen-1]
		if str[0:1] != "%" || (strtype != "n" && strtype != "d" && strtype != "s") {
			fmt.Println("data format error")
			return
		}
		switch strtype {
		case "s":
			if strings.Contains(limit, "-") {
				limits := strings.Split(limit, "-")
				min, _ := strconv.Atoi(limits[0])
				max, _ := strconv.Atoi(limits[1])
				if max < min || max <= 0 || min <= 0 {
					fmt.Println("data format error")
					return
				}
				contain[idx] = gen.gen_varchar(min, max)
			} else {
				length, _ := strconv.Atoi(limit)
				if length <= 0 {
					fmt.Println("data format error")
					return
				}
				contain[idx] = gen.gen_char(length)
			}
			fmtFormat[idx] = "%s"
		case "d":
			if strings.Contains(limit, "-") {
				limits := strings.Split(limit, "-")
				min, _ := strconv.Atoi(limits[0])
				max, _ := strconv.Atoi(limits[1])
				if max < min || max <= 0 || min < 0 {
					fmt.Println("data format error")
					return
				}
				contain[idx] = gen.gen_varint(min, max)
			} else {
				length, _ := strconv.Atoi(limit)
				if length <= 0 {
					fmt.Println("data format error")
					return
				}
				contain[idx] = gen.gen_int(length)
			}
			fmtFormat[idx] = "%d"
		case "n":
			base := 0
			if limit != "" {
				base, _ = strconv.Atoi(limit)
			}
			contain[idx] = gen.gen_autoincr(base)
			fmtFormat[idx] = "%d"
		}
	}

	var i int64
	for i = 0; i < *rows; i++ {
		tmp := make([]string, len(strs))
		for k, v := range contain {
			tmp[k] = fmt.Sprintf(fmtFormat[k], <-v)
		}
		fmt.Println(strings.Join(tmp, ","))
	}
}
