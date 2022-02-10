package valid

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/russtone/iprange"
)

//
// Directory
//

// Directory checks if validatable value is valid path to directory.
func Directory() validation.Rule {
	return &directoryRule{}
}

type directoryRule struct{}

// Validate allows directoryRule to implement validation.Rule interface.
func (r *directoryRule) Validate(value interface{}) error {
	path, _ := value.(string)

	if fi, err := os.Stat(path); os.IsNotExist(err) {
		return err
	} else if !fi.IsDir() {
		return errors.New("must be a directory, not a file")
	}

	return nil
}

//
// File
//

// File checks if validatable value is valid path to file.
func File() validation.Rule {
	return &fileRule{}
}

type fileRule struct{}

// Validate allows fileRule to implement validation.Rule interface.
func (r *fileRule) Validate(value interface{}) error {
	path, _ := value.(string)

	if fi, err := os.Stat(path); os.IsNotExist(err) {
		return err
	} else if fi.IsDir() {
		return errors.New("must be a file, not a directory")
	}

	return nil
}

//
// OneOf
//

// OneOf checks if validatable value is present among the provided values.
// If second argument is true case sensitive comparison will be used.
func OneOf(values []string, caseSensetive bool) validation.Rule {
	return &oneOfRule{values, caseSensetive}
}

type oneOfRule struct {
	values        []string
	caseSensetive bool
}

// Validate allows oneOfRule to implement validation.Rule interface.
func (r *oneOfRule) Validate(value interface{}) error {
	val, _ := value.(string)

	for _, v := range r.values {
		if (r.caseSensetive && v == val) ||
			(!r.caseSensetive && strings.EqualFold(v, val)) {
			return nil
		}
	}

	return fmt.Errorf("invalid value, expected one of %s", strings.Join(r.values, ","))
}

//
// IPRange
//

// IPRange checks if validatable value is valid IPv4 or IPv6 range.
// Can be in one of the following forms:
// - Single address. Examples: `192.168.1.1`, `2001:db8:a0b:12f0::1`
// - CIDR. Examples: `192.168.1.0/24`, `2001:db8:a0b:12f0::1`
// - Begin_End. Examples: `192.168.1.10_192.168.2.20`, `2001:db8:a0b:12f0::1_2001:db8:a0b:12f0::10`
// - Octets ranges: `192.168.1,3-5.1-10`, `2001:db8:a0b:12f0::1,1-10`
func IPRange() validation.Rule {
	return &iprangeRule{}
}

type iprangeRule struct{}

// Validate allows iprangeRule to implement validation.Rule interface.
func (r *iprangeRule) Validate(value interface{}) error {
	s := value.(string)
	if r := iprange.Parse(s); r == nil {
		return fmt.Errorf("invalid range %q", s)
	}
	return nil
}

//
// Regexp
//

// Regexp checks if validatable value is valid regular expression.
func Regexp() validation.Rule {
	return &regexpRule{}
}

type regexpRule struct{}

// Validate allows regexpRule to implement validation.Rule interface.
func (r *regexpRule) Validate(value interface{}) error {
	s := value.(string)
	if _, err := regexp.Compile(s); err != nil {
		return err
	}
	return nil
}
