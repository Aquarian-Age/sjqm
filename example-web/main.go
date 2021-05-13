package main

import (
	"flag"
	"fmt"
	"github.com/Aquarian-Age/sjqm/api"
)

func main() {
	flag.Usage = func() {
		fmt.Println("sjqm -l port")
		flag.PrintDefaults()
	}
	p := flag.Int("l", 8111, "")
	flag.Parse()
	port := fmt.Sprintf(":%d", *p)
	fmt.Println("Default Server Port:", 8111)
	api.Main(port)
}
