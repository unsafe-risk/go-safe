package safe

func StrZero() string {
	return ""
}

func NumZero[T number]() T {
	return 0
}

func IZero() int {
	return NumZero[int]()
}

func I8Zero() int8 {
	return NumZero[int8]()
}

func I16Zero() int16 {
	return NumZero[int16]()
}

func I32Zero() int32 {
	return NumZero[int32]()
}

func I64Zero() int64 {
	return NumZero[int64]()
}

func UiZero() uint {
	return NumZero[uint]()
}

func Ui8Zero() uint8 {
	return NumZero[uint8]()
}

func Ui16Zero() uint16 {
	return NumZero[uint16]()
}

func Ui32Zero() uint32 {
	return NumZero[uint32]()
}

func Ui64Zero() uint64 {
	return NumZero[uint64]()
}

func F32Zero() float32 {
	return NumZero[float32]()
}

func F64Zero() float64 {
	return NumZero[float64]()
}
