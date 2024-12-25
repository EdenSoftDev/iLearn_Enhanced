package client

import (
	"testing"
)

func TestGetTpassLt(t *testing.T) {
	lt, execution, err := GetTpassLt()
	if err != nil {
		t.Error(err)
	}
	t.Log(lt, execution)
}
