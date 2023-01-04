package h2o

import (
	"testing"
)

func TestParseTime(t *testing.T) {
	// flag := "12d1h2m3s"
	flag := "2m"

	seconds, err := TimeParse(flag)
	t.Logf("value %d %s", seconds, err)

	value := TimeFormatStyle(seconds)
	t.Logf("value %s %v", value, value == flag)
}
