package safe

import "time"

func TimeOrNow(t *time.Time) time.Time {
	return TimeOrDefault(t, time.Now)
}

func StrOrZero(s *string) string {
	return StrOrDefault(s, StrZero)
}

func NumOrZero[T number](n *T) T {
	return NumOrDefault[T](n, NumZero[T])
}

func IOrZero(n *int) int {
	return NumOrZero(n)
}

func I8OrZero(n *int8) int8 {
	return NumOrZero(n)
}

func I16OrZero(n *int16) int16 {
	return NumOrZero(n)
}

func I32OrZero(n *int32) int32 {
	return NumOrZero(n)
}

func I64OrZero(n *int64) int64 {
	return NumOrZero(n)
}

func UiOrZero(n *uint) uint {
	return NumOrZero(n)
}

func Ui8OrZero(n *uint8) uint8 {
	return NumOrZero(n)
}

func Ui16OrZero(n *uint16) uint16 {
	return NumOrZero(n)
}

func Ui32OrZero(n *uint32) uint32 {
	return NumOrZero(n)
}

func Ui64OrZero(n *uint64) uint64 {
	return NumOrZero(n)
}

func F32OrZero(n *float32) float32 {
	return NumOrZero(n)
}

func F64OrZero(n *float64) float64 {
	return NumOrZero(n)
}
