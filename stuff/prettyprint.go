package stuff

import (
	"fmt"
)

func PrettyPrint(s todo) {
	for i, v := range s {
		fmt.Println(i+1, ">", v)
	}
}
