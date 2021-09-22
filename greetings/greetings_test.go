package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Krzys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Krzys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Krzys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
func TestHelloEmpty(t *testing.T) {
	var s string = ""
	msg, err := Hello(s)
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}

}
