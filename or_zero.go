package safe

import "time"

func TimeOrNow(t *time.Time) time.Time {
	return TimeOrDefault(t, time.Now)
}

func StrOrZero(s *string) string {
	return StrOrDefault(s, StrZero)
}

func IOrZero(n *int) int {
	return IOrDefault(n, IZero)
}

func Int8OrZero(n *int8) int8 {
	return I8OrDefault(n, I8Zero)
}

func Int16OrZero(n *int16) int16 {
	return I16OrDefault(n, I16Zero)
}

func Int32OrZero(n *int32) int32 {
	return I32OrDefault(n, I32Zero)
}

func Int64OrZero(n *int64) int64 {
	return I64OrDefault(n, I64Zero)
}

func UintOrZero(n *uint) uint {
	return UiOrDefault(n, UiZero)
}

func Uint8OrZero(n *uint8) uint8 {
	return Ui8OrDefault(n, Ui8Zero)
}

func Uint16OrZero(n *uint16) uint16 {
	return Ui16OrDefault(n, Ui16Zero)
}

func Uint32OrZero(n *uint32) uint32 {
	return Ui32OrDefault(n, Ui32Zero)
}

func Uint64OrZero(n *uint64) uint64 {
	return Ui64OrDefault(n, Ui64Zero)
}

func Float32OrZero(n *float32) float32 {
	return F32OrDefault(n, F32Zero)
}

func Float64OrZero(n *float64) float64 {
	return F64OrDefault(n, F64Zero)
}
