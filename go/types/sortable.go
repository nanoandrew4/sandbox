package types

type Sortable interface {
	byte | int | int8 | int16 | int32 | int64 | string | float32 | float64
}
