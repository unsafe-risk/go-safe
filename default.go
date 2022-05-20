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

func IOrDefault(n *int, def func() int) int {
	if n == nil {
		return def()
	}
	return *n
}

func I8OrDefault(n *int8, def func() int8) int8 {
	if n == nil {
		return def()
	}
	return *n
}

func I16OrDefault(n *int16, def func() int16) int16 {
	if n == nil {
		return def()
	}
	return *n
}

func I32OrDefault(n *int32, def func() int32) int32 {
	if n == nil {
		return def()
	}
	return *n
}

func I64OrDefault(n *int64, def func() int64) int64 {
	if n == nil {
		return def()
	}
	return *n
}

func UiOrDefault(n *uint, def func() uint) uint {
	if n == nil {
		return def()
	}
	return *n
}

func Ui8OrDefault(n *uint8, def func() uint8) uint8 {
	if n == nil {
		return def()
	}
	return *n
}

func Ui16OrDefault(n *uint16, def func() uint16) uint16 {
	if n == nil {
		return def()
	}
	return *n
}

func Ui32OrDefault(n *uint32, def func() uint32) uint32 {
	if n == nil {
		return def()
	}
	return *n
}

func Ui64OrDefault(n *uint64, def func() uint64) uint64 {
	if n == nil {
		return def()
	}
	return *n
}

func F32OrDefault(n *float32, def func() float32) float32 {
	if n == nil {
		return def()
	}
	return *n
}

func F64OrDefault(n *float64, def func() float64) float64 {
	if n == nil {
		return def()
	}
	return *n
}
