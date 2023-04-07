package contracts

type Basic interface {
	string | Number | bool
}

type Integer interface {
	int | int64 | int32 | int16 | int8 | uint | uint64 | uint32 | uint16 | uint8
}

type Float interface {
	float64 | float32
}

type Number interface {
	Integer | Float
}
