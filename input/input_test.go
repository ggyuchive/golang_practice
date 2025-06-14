package input_test

import (
	"ggyuchive/input"
	"testing"
)

func TestNextIntStr(t *testing.T) {
	// defer executes after panic() executes
	defer func() {
		// if recover() executes in defer, recover panic
		// if not, program crashes
		r := recover()
		if r == nil {
			t.Errorf("expected panic but did not panic")
		} else {
			msg := r.(string)
			if msg != "NextInt: failed to scan token (eof or input error)" {
				t.Errorf("unexpected panic message: %s", msg)
			}
		}
	}()

	testInput := " -10\n20yr \n tbd \t abc 1001\n-1024\n"
	input.InitForTesting(testInput)

	if x := input.NextInt(); x != -10 {
		t.Errorf("expected -10, got %d", x)
	}
	if x := input.NextStr(); x != "20yr" {
		t.Errorf("expected '20yr', got %s", x)
	}
	if x := input.NextStr(); x != "tbd" {
		t.Errorf("expected 'tbd', got '%s'", x)
	}
	if x := input.NextStr(); x != "abc" {
		t.Errorf("expected 'abc', got '%s'", x)
	}
	if x := input.NextInt(); x != 1001 {
		t.Errorf("expected 1001, got '%d'", x)
	}
	if x := input.NextInt(); x != -1024 {
		t.Errorf("expected -1024, got '%d'", x)
	}
	input.NextInt()
}

func TestNextLongLongInt(t *testing.T) {
	testInput := " 9223372036854775807 -9223372036854775808\n"
	input.InitForTesting(testInput)

	if x := input.NextInt(); x != 9223372036854775807 {
		t.Errorf("expected 9223372036854775807, got %d", x)
	}
	if x := input.NextInt(); x != -9223372036854775808 {
		t.Errorf("expected -9223372036854775808, got %d", x)
	}
}
