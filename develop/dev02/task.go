package develop

import (
	"errors"
	"strconv"
	"strings"
)

const numeric = "1234567890"
const escapable = numeric + "\\"

var (
	ErrNoCharBeforeNum = errors.New("no char before num to unpack")
	ErrInescapableChar = errors.New("char is inescapable")
)

type Unpacker struct {
	str  string
	cart int
	err  error
}

func (u *Unpacker) Init(str string) (self *Unpacker) {
	u.str = str
	u.cart = 0
	u.err = nil
	return u
}

func (u *Unpacker) Err() error { return u.err }

// Returns char that `cart` is currentyl pointing at.
// If `cart` is out of bounds for underlying `string` returns empty `string`.
func (u *Unpacker) currentChar() string {
	if u.cart >= len(u.str) || u.cart < 0 {
		return ""
	}
	return u.str[u.cart : u.cart+1]
}

func (u *Unpacker) currentRune() rune {
	return rune(u.currentChar()[0])
}

func (u *Unpacker) parseNext() string {
	char := u.currentChar()
	if u.currentChar() == "\\" { // check for escape sequence and it's validity if present...
		u.cart++
		if u.cart >= len(u.str) || !isEscapable(u.currentRune()) {
			u.err = UnpackError{ErrInescapableChar, u.cart}
			return ""
		} else {
			char = u.currentChar()
		}
	} else { // if not check what kind of normal symbol we've got
		if isNumeric(rune(char[0])) {
			// I think that standalone numbers without symbols in front should be treated
			// as if we are unpacking "". Result of this will be empty string.
			// Comment this block out and uncomment line bellow if you think the same.
			// (some tests will fail)
			// char = ""

			u.err = UnpackError{ErrNoCharBeforeNum, u.cart}
			return ""
		}
	}

	// Parse number after symbol that will be unpacked.
	// Lack of number implies 1 repetition.
	mul, times := parseNum(u.str[u.cart+1:]), 1
	if mul != "" {
		times, _ = strconv.Atoi(mul)
	}

	u.cart += len(mul) + 1 // advance cart

	return strings.Repeat(char, times)
}

func (u *Unpacker) Unpack() (string, error) {
	var res string
	for u.str[u.cart:] != "" {
		res += u.parseNext()
		if u.err != nil {
			return "", u.err
		}
	}

	return res, nil
}
