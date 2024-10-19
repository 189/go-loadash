package errors

import (
	"fmt"
)

func Raise(err error) {
	e := fmt.Errorf("got error: %v", err)
	fmt.Println(e)
}
