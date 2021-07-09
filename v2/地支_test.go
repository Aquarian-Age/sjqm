package sjqm

import (
	"fmt"
	"testing"
)

func TestNewZHI(t *testing.T) {
	zhi := "子"
	obj := NewZHI(zhi)
	fmt.Println(obj)
}
