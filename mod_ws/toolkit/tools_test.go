package toolkit

import "testing"

func TestTools_RandomString(t *testing.T) {
	var testTools Tools
	s := testTools.GenerateRandomString(10)
	if len(s) != 10 {
		t.Error("Wrong length of generated string")
	}
}
