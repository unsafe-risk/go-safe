package safe

type number interface {
	int8 | int16 | int32 | int64 | int |
		uint8 | uint16 | uint32 | uint64 | uint |
		float32 | float64
}

type Safety[T any] interface {
	IsNil() bool
	Default() T
}
