package app

import (
	"bytes"
	"testing"
)

func TestParseLog(t *testing.T) {
	testdata := "number,product,company,value,stock\n1,game,DeNA,1000,8000000\n2,chromebook,google,40000,200000\n3,book,amazon,3000,50000000"
	r := bytes.NewReader([]byte(testdata))
	itemdatas, err := parseLog(r)
	if err != nil {
		t.Errorf("parseLog(%v)=%v, got=%v", testdata, err, itemdatas)
	}
	t.Logf("itemdatas=%v", itemdatas)

}
func TestLogItems(t *testing.T) {
	testdata := "1,game,DeNA,1000,8000000"
	itemdata, err := logItems(testdata)
	if err != nil {
		t.Errorf("logItems(%v)=%v: got=%v", testdata, err, itemdata)
	}
	t.Logf("itemdata=%v", itemdata)

}
