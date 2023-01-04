package random

import "testing"

func TestRandomSample(t *testing.T) {
	bytes := []byte("0123abcd")
	randomBytes := Sample(bytes, 10)
	t.Logf("randomBytes %s", randomBytes)

	runes := []rune("0123abcd这里有中文한국어")
	randomRunes := Sample(runes, 10)
	t.Logf("randomRunes %s", string(randomRunes))
}

func TestRandomChoice(t *testing.T) {
	bytes := []byte("0123abcd")
	randomByte := Choice(bytes)
	t.Logf("randomByte %s", string(randomByte))

	runes := []rune("0123abcd这里有中文한국어")
	randomRune := Choice(runes)
	t.Logf("randomRune %s", string(randomRune))
}

func TestRandomNum(t *testing.T) {
	result := NumberSplit(15, 4)
	t.Logf("%v", result)
}
