package pattern

import (
	"errors"
	"fmt"
	"strings"
)

type APICall struct {
	call string
}

func (c *APICall) Call() {
	fmt.Printf("making an api call to %s\n", c.call)
}

type Validator struct {
	c   *APICall
	err error
}

func (v *Validator) Err() error { return v.err }

func (v *Validator) Init() (self *Validator) {
	v.c = &APICall{}
	return v
}

func (v *Validator) ValidateURL(url string) (self *Validator) {
	if v.err != nil {
		return v
	}

	// example....com will be split into "example" and "com" which will result in correct output
	parts := strings.Split(url, ".")
	switch {
	case len(parts) == 1:
		v.err = errors.New("invalid url")
	case len(parts) == 2:
		if parts[1] != "com" && parts[1] != "org" && parts[1] != "ru" {
			v.err = errors.New("invalid domain type")
		}
		if strings.ContainsAny(parts[0], "./,'\":;[]{}!@#$%") {
			v.err = errors.New("invalid domain name")
		}
		v.c.call = fmt.Sprintf("%s%s", url, v.c.call)
	default:
		v.err = errors.New("invalid url")
	}

	return v
}

func (v *Validator) ValidateID(id uint) (self *Validator) {
	if v.err != nil {
		return v
	}

	if id > 100 {
		v.err = errors.New("id is in invalid range")
		return v
	}

	v.c.call += fmt.Sprintf("/?id=%d", id)
	return v
}

func (v *Validator) Result() (*APICall, error) { return v.c, v.err }

func init() {
	fmt.Println("Demonstrating Chain of Responsobility pattern:")
	defer fmt.Println()

	c, err := (&Validator{}).Init().ValidateURL("example.com").ValidateID(12).Result()
	if err != nil {
		fmt.Println(fmt.Errorf("invalid uri: %w", err))
	} else {
		c.Call()
	}
}
