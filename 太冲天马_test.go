package sjqm

import (
	"strings"
	"testing"
)

//太冲天马
func TestTianMa(t *testing.T) {

	yjobj := []struct{ yj, hgz string }{
		{yj: "亥", hgz: "丙寅"},
		{yj: "午", hgz: "甲午"},
		{yj: "子", hgz: "丁卯"},
		{yj: "丑", hgz: "戊辰"},
	}
	want := []string{"午", "卯", "午", "酉"}
	for i := 0; i < len(yjobj); i++ {
		tms := TianMa(yjobj[i].yj, yjobj[i].hgz)
		if !strings.EqualFold(tms, want[i]) {

		}
	}
}
