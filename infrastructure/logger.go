package infrastructure

import (
	"fmt"
)

type Logger struct{}

func (l Logger) Log(message string) {
	fmt.Println(message)
}
