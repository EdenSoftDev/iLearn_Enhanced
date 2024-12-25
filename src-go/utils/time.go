package utils

import (
	"fmt"
	"time"
)

func GetTime() string {
	return fmt.Sprint(time.Now().UnixMilli())
}
