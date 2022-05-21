package safe

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimeOrDefault(t *testing.T) {
	var val *time.Time

	tVal1 := time.Now().Add(time.Hour)
	def := func() time.Time {
		return tVal1
	}

	assert.Equal(t, TimeOrDefault(val, def), tVal1)

	val = new(time.Time)
	assert.Equal(t, TimeOrDefault(val, def), *val)
}

func TestStrOrDefault(t *testing.T) {
	var str *string
	def := func() string {
		return "hello world"
	}

	assert.Equal(t, StrOrDefault(str, def), "hello world")

	str = new(string)
	assert.Equal(t, StrOrDefault(str, def), "")
}

func TestIOrDefault(t *testing.T) {
	var i *int
	def := func() int {
		return 101010101
	}

	assert.Equal(t, IOrDefault(i, def), 101010101)

	i = new(int)
	assert.Equal(t, IOrDefault(i, def), 0)
}

func TestI8OrDefault(t *testing.T) {
	var i8 *int8
	def := func() int8 {
		return 101
	}

	assert.Equal(t, I8OrDefault(i8, def), int8(101))

	i8 = new(int8)
	assert.Equal(t, I8OrDefault(i8, def), int8(0))
}

func TestI16OrDefault(t *testing.T) {
	var i16 *int16
	def := func() int16 {
		return 10101
	}

	assert.Equal(t, I16OrDefault(i16, def), int16(10101))

	i16 = new(int16)
	assert.Equal(t, I16OrDefault(i16, def), int16(0))
}

func TestI32OrDefault(t *testing.T) {
	var i32 *int32
	def := func() int32 {
		return 1010101
	}

	assert.Equal(t, I32OrDefault(i32, def), int32(1010101))

	i32 = new(int32)
	assert.Equal(t, I32OrDefault(i32, def), int32(0))
}

func TestI64OrDefault(t *testing.T) {
	var i64 *int64
	def := func() int64 {
		return 1010101
	}

	assert.Equal(t, I64OrDefault(i64, def), int64(1010101))

	i64 = new(int64)
	assert.Equal(t, I64OrDefault(i64, def), int64(0))
}

func TestUiOrDefault(t *testing.T) {
	var ui *uint
	def := func() uint {
		return 1101010101
	}

	assert.Equal(t, UiOrDefault(ui, def), uint(1101010101))

	ui = new(uint)
	assert.Equal(t, UiOrDefault(ui, def), uint(0))
}

func TestUi8OrDefault(t *testing.T) {
	var ui8 *uint8
	def := func() uint8 {
		return 110
	}

	assert.Equal(t, Ui8OrDefault(ui8, def), uint8(110))

	ui8 = new(uint8)
	assert.Equal(t, Ui8OrDefault(ui8, def), uint8(0))
}

func TestUi16OrDefault(t *testing.T) {
	var ui16 *uint16
	def := func() uint16 {
		return 11010
	}

	assert.Equal(t, Ui16OrDefault(ui16, def), uint16(11010))

	ui16 = new(uint16)
	assert.Equal(t, Ui16OrDefault(ui16, def), uint16(0))
}

func TestUi32OrDefault(t *testing.T) {
	var ui32 *uint32
	def := func() uint32 {
		return 11010101
	}

	assert.Equal(t, Ui32OrDefault(ui32, def), uint32(11010101))

	ui32 = new(uint32)
	assert.Equal(t, Ui32OrDefault(ui32, def), uint32(0))
}

func TestUi64OrDefault(t *testing.T) {
	var ui64 *uint64
	def := func() uint64 {
		return 11010101010
	}

	assert.Equal(t, Ui64OrDefault(ui64, def), uint64(11010101010))

	ui64 = new(uint64)
	assert.Equal(t, Ui64OrDefault(ui64, def), uint64(0))
}

func TestF32OrDefault(t *testing.T) {
	var f32 *float32
	def := func() float32 {
		return 1.123
	}

	assert.Equal(t, F32OrDefault(f32, def), float32(1.123))

	f32 = new(float32)
	assert.Equal(t, F32OrDefault(f32, def), float32(0))
}

func TestF64OrDefault(t *testing.T) {
	var f64 *float64
	def := func() float64 {
		return 3.141592
	}

	assert.Equal(t, F64OrDefault(f64, def), 3.141592)

	f64 = new(float64)
	assert.Equal(t, F64OrDefault(f64, def), float64(0))
}
