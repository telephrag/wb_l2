package develop

import (
	"errors"
	"testing"
)

func TestUnpack(t *testing.T) {

	input, expected := "\"10\\14\\\\abc2", "\"\"\"\"\"\"\"\"\"\"1111\\abcc"
	res, err := (&Unpacker{}).Init(input).Unpack()
	if err != nil {
		t.Errorf("ERROR: expected no error; received: %v\n", err)
	}

	if res != expected {
		t.Errorf("ERROR: expected: \"%s\"; received: \"%s\"\n", expected, res)
	} else {
		t.Logf("OK:    %-12s -> %s", input, res)
	}

	input = "\\"
	_, err = (&Unpacker{}).Init(input).Unpack()
	if err == nil {
		t.Errorf("ERROR: expected error; received: nil")
	} else {
		inner := errors.Unwrap(err)
		if inner != ErrInescapableChar {
			t.Errorf(
				"ERROR: expected inner error to be: %v; received: %v\n",
				ErrInescapableChar,
				inner,
			)
		} else {
			t.Logf("OK:    %-12s -> %v", input, err)
		}
	}

	input = "45"
	_, err = (&Unpacker{}).Init(input).Unpack()
	if err == nil {
		t.Errorf("ERROR: expected error; received: nil")
	} else {
		inner := errors.Unwrap(err)
		if inner != ErrNoCharBeforeNum {
			t.Errorf(
				"ERROR: expected inner error to be: %v; received: %v\n",
				ErrNoCharBeforeNum,
				inner,
			)
		} else {
			t.Logf("OK:    %-12s -> %v", input, err)
		}
	}
}
