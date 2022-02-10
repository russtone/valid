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

func Directory() validation.Rule {
	return &directoryRule{}
}

type directoryRule struct{}

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

func File() validation.Rule {
	return &fileRule{}
}

type fileRule struct{}

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

func OneOf(values []string, caseSensetive bool) validation.Rule {
	return &oneOfRule{values, caseSensetive}
}

type oneOfRule struct {
	values        []string
	caseSensetive bool
}

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

func IPRange() validation.Rule {
	return &iprangeRule{}
}

type iprangeRule struct{}

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

func Regexp() validation.Rule {
	return &regexpRule{}
}

type regexpRule struct{}

func (r *regexpRule) Validate(value interface{}) error {
	s := value.(string)
	if _, err := regexp.Compile(s); err != nil {
		return err
	}
	return nil
}
