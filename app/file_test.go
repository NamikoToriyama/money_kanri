package app

import (
	"strings"
	"testing"
)

func TestParseLog(t *testing.T) {
	testdata := "number,product,company,value,stock\n1,game,DeNA,1000,8000000\n2,chromebook,google,40000,200000\n3,book,amazon,3000,50000000"

	r := strings.NewReader(testdata)
	parseLog(r)

}
