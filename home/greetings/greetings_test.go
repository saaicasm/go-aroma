package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Aegon"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello("Aegon")

	if !want.MatchString(msg) || err != nil {
		t.Fatal("Either string not match or some error")
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatal("Either the returned string is not empty or no error")
	}

}
