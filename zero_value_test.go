package safe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIZero(t *testing.T) {
	assert.Zero(t, IZero())
}

func TestI8Zero(t *testing.T) {
	assert.Zero(t, I8Zero())
}

func TestI16Zero(t *testing.T) {
	assert.Zero(t, I16Zero())
}

func TestI32Zero(t *testing.T) {
	assert.Zero(t, I32Zero())
}

func TestI64Zero(t *testing.T) {
	assert.Zero(t, I64Zero())
}

func TestUiZero(t *testing.T) {
	assert.Zero(t, UiZero())
}

func TestUi8Zero(t *testing.T) {
	assert.Zero(t, Ui8Zero())
}

func TestUi16Zero(t *testing.T) {
	assert.Zero(t, Ui16Zero())
}

func TestUi32Zero(t *testing.T) {
	assert.Zero(t, Ui32Zero())
}

func TestUi64Zero(t *testing.T) {
	assert.Zero(t, Ui64Zero())
}

func TestF32Zero(t *testing.T) {
	assert.Zero(t, F32Zero())
}

func TestF64Zero(t *testing.T) {
	assert.Zero(t, F64Zero())
}
