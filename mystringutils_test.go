package mystringutils_test

import (
	"strings"
	"testing"

	"github.com/michaelsobota/mystringutils"
)

func TestUpper(t *testing.T) {
	ts := mystringutils.Upper("i should be uppercase")
	if strings.Compare(ts, "I SHOULD BE UPPERCASE") != 0 {
		t.Error("expected string to be uppercase")
	}
}

func TestLower(t *testing.T) {
	ts := mystringutils.Lower("I SHOULD BE LOWERCASE")
	if strings.Compare(ts, "i should be lowercase") != 0 {
		t.Error("expected string to be lowercase")
	}
}
