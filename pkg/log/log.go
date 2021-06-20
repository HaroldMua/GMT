package log

import (
	"fmt"
	"time"
)

func ErrPrint(err error) {
	fmt.Printf("[GMT]-[%s] Error: %s\n", time.Now(), err)
}

func Print(log string) {
	fmt.Printf("[GMT]-[%s] %s\n", time.Now(), log)
}