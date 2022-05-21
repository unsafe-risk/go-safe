package safe

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestTimeOrNow(t *testing.T) {
	var val *time.Time
	assert.NotNil(t, TimeOrNow(val))

	val = new(time.Time)
	assert.Equal(t, TimeOrNow(val), *val)
}

func TestStrOrZero(t *testing.T) {
	var str *string
	assert.Zero(t, StrOrZero(str))

	str = new(string)
	*str = "hello world"
	assert.NotZero(t, StrOrZero(str))
}

func TestIOrZero(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var i *int
	assert.Zero(t, IOrZero(i))

	i = new(int)
	*i = rand.Int()
	assert.NotZero(t, IOrZero(i))
}

func TestInt8OrZero(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var i8 *int8
	assert.Zero(t, I8OrZero(i8))

	i8 = new(int8)
	*i8 = int8(rand.Int())
	assert.NotZero(t, I8OrZero(i8))
}

func TestInt16OrZero(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var i16 *int16
	assert.Zero(t, I16OrZero(i16))

	i16 = new(int16)
	*i16 = int16(rand.Int())
	assert.NotZero(t, I16OrZero(i16))
}

func TestInt32OrZero(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var i32 *int32
	assert.Zero(t, I32OrZero(i32))

	i32 = new(int32)
	*i32 = rand.Int31()
	assert.NotZero(t, I32OrZero(i32))
}

func TestInt64OrZero(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var i64 *int64
	assert.Zero(t, I64OrZero(i64))

	i64 = new(int64)
	*i64 = rand.Int63()
	assert.NotZero(t, I64OrZero(i64))
}

func TestUintOrZero(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var ui *uint
	assert.Zero(t, UiOrZero(ui))

	ui = new(uint)
	*ui = uint(rand.Uint64())
	assert.NotZero(t, UiOrZero(ui))
}

func TestUint8OrZero(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var ui8 *uint8
	assert.Zero(t, Ui8OrZero(ui8))

	ui8 = new(uint8)
	*ui8 = uint8(rand.Uint64())
	assert.NotZero(t, Ui8OrZero(ui8))
}

func TestUint16OrZero(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var ui16 *uint16
	assert.Zero(t, Ui16OrZero(ui16))

	ui16 = new(uint16)
	*ui16 = uint16(rand.Uint64())
	assert.NotZero(t, Ui16OrZero(ui16))
}

func TestUint32OrZero(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var ui32 *uint32
	assert.Zero(t, Ui32OrZero(ui32))

	ui32 = new(uint32)
	*ui32 = rand.Uint32()
	assert.NotZero(t, Ui32OrZero(ui32))
}

func TestUint64OrZero(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var ui64 *uint64
	assert.Zero(t, Ui64OrZero(ui64))

	ui64 = new(uint64)
	*ui64 = rand.Uint64()
	assert.NotZero(t, Ui64OrZero(ui64))
}

func TestFloat32OrZero(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var f32 *float32
	assert.Zero(t, F32OrZero(f32))

	f32 = new(float32)
	*f32 = rand.Float32()
	assert.NotZero(t, F32OrZero(f32))
}

func TestFloat64OrZero(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var f64 *float64
	assert.Zero(t, F64OrZero(f64))

	f64 = new(float64)
	*f64 = rand.Float64()
	assert.NotZero(t, F64OrZero(f64))
}
