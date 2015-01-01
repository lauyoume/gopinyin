package main

import (
	"fmt"
	"github.com/lauyoume/gopinyin"
)

func main() {
	fmt.Println(gopinyin.Convert("Hello,四节穇穈穉", false))
	fmt.Println(gopinyin.Convert("Hel穇穈穉", false))
}
