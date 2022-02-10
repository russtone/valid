package valid_test

import (
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"

	"github.com/russtone/valid"
)

func Test_Directory(t *testing.T) {
	var err error

	err = validation.Validate("./.github", valid.Directory())
	assert.NoError(t, err)

	err = validation.Validate("./README.md", valid.Directory())
	assert.Error(t, err)

	err = validation.Validate("./not-exists", valid.Directory())
	assert.Error(t, err)
}

func Test_File(t *testing.T) {
	var err error

	err = validation.Validate("./README.md", valid.File())
	assert.NoError(t, err)

	err = validation.Validate("./.github", valid.File())
	assert.Error(t, err)

	err = validation.Validate("./not-exists", valid.File())
	assert.Error(t, err)
}

func Test_OneOf(t *testing.T) {
	var err error

	err = validation.Validate("one", valid.OneOf([]string{"one", "two", "three"}, false))
	assert.NoError(t, err)

	err = validation.Validate("ONE", valid.OneOf([]string{"one", "two", "three"}, false))
	assert.NoError(t, err)

	err = validation.Validate("four", valid.OneOf([]string{"one", "two", "three"}, false))
	assert.Error(t, err)

	err = validation.Validate("One", valid.OneOf([]string{"one", "two", "three"}, true))
	assert.Error(t, err)
}

func Test_IPRange(t *testing.T) {
	var err error

	err = validation.Validate("192.168.1.1/24", valid.IPRange())
	assert.NoError(t, err)

	err = validation.Validate("192.168.1.1-255", valid.IPRange())
	assert.NoError(t, err)

	err = validation.Validate("192.168.1.1_192.168.1.2", valid.IPRange())
	assert.NoError(t, err)

	err = validation.Validate("256.1.1.1", valid.IPRange())
	assert.Error(t, err)
}

func Test_Regexp(t *testing.T) {
	var err error

	err = validation.Validate(`[a-z0-9]+\.test`, valid.Regexp())
	assert.NoError(t, err)

	err = validation.Validate(`[a-z0-9+\.test`, valid.Regexp())
	assert.Error(t, err)
}
