package api

import (
	"flag"
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Usage = func() {
		fmt.Println("sjqm -l port")
		flag.PrintDefaults()
	}
	p := flag.Int("l", 8111, "")
	flag.Parse()
	port := fmt.Sprintf(":%d", *p)
	fmt.Println("Default Server Port:", 8111)
	Main(port)
}
