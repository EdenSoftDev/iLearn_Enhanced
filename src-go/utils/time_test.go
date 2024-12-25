package utils

import (
	"testing"
)

func TestTime(t *testing.T) {
	result := GetTime()
	t.Log(result)
}
