package client

import (
	"testing"
)

func TestGetCasLt(t *testing.T) {
	lt, execution, err := GetILearnLt()
	if err != nil {
		t.Error(err)
	}
	t.Log(lt, execution)
}
