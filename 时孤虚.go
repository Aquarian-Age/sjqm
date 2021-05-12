package sjqm

import (
	"fmt"
)

//时孤虚
func GuXuH(dgz, hgz string) (gxs string) {
	guxu := HGuXu(dgz, hgz)
	if _, ok := guxu["孤"]; ok {
		gxs = fmt.Sprintf("孤:%s 虚:%s", guxu["孤"], guxu["虚"])
	}
	return
}
