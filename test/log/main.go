package main

import (
	"os"
	"sample/common/kit"
)

func main() {
	l := kit.NewJSONLogger(os.Stdout)
	l = kit.LoggerWith(l, "service", "test")
	err := l.Log("qwe", "lkj")
	if err != nil {
		panic(err)
	}
}
