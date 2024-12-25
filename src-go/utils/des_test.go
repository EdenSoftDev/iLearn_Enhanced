package utils

import (
	"testing"
)

func TestRSA(t *testing.T) {
	username := "test"
	password := "test"
	lt := "LT-360142-T4enUcmCQWFAbejsrcPrfgYFFeMHvF-tpass"
	result := StrEnc(username+password+lt, "1", "2", "3")
	t.Log(result)
}
