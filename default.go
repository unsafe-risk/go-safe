package safe

import (
	"time"
)

func TimeOrDefault(t *time.Time, def func() time.Time) time.Time {
	if t == nil {
		return def()
	}
	return *t
}

func StrOrDefault(s *string, def func() string) string {
	if s == nil {
		return def()
	}
	return *s
}

func NumOrDefault[T number](n *T, def func() T) T {
	if n == nil {
		return def()
	}

	return *n
}

func Safe[T Safety[T]](in T) T {
	if in.IsNil() {
		return in.Default()
	}

	return in
}

func IOrDefault(n *int, def func() int) int {
	return NumOrDefault(n, def)
}

func I8OrDefault(n *int8, def func() int8) int8 {
	return NumOrDefault(n, def)
}

func I16OrDefault(n *int16, def func() int16) int16 {
	return NumOrDefault(n, def)
}

func I32OrDefault(n *int32, def func() int32) int32 {
	return NumOrDefault(n, def)
}

func I64OrDefault(n *int64, def func() int64) int64 {
	return NumOrDefault(n, def)
}

func UiOrDefault(n *uint, def func() uint) uint {
	return NumOrDefault(n, def)
}

func Ui8OrDefault(n *uint8, def func() uint8) uint8 {
	return NumOrDefault(n, def)
}

func Ui16OrDefault(n *uint16, def func() uint16) uint16 {
	return NumOrDefault(n, def)
}

func Ui32OrDefault(n *uint32, def func() uint32) uint32 {
	return NumOrDefault(n, def)
}

func Ui64OrDefault(n *uint64, def func() uint64) uint64 {
	return NumOrDefault(n, def)
}

func F32OrDefault(n *float32, def func() float32) float32 {
	return NumOrDefault(n, def)
}

func F64OrDefault(n *float64, def func() float64) float64 {
	return NumOrDefault(n, def)
}
