package contracts

import (
	"reflect"
)

type Context interface {
	// Get 从上下文中检索数据
	// retrieves data from the context.
	Get(key string) any

	// Set 在上下文中保存数据
	// saves data in the context.
	Set(key string, val any)
}

type Fields map[string]any

func (f Fields) Set(fields Fields) {
	for k, v := range fields {
		f[k] = v
	}
}
func (f Fields) Get(key string) any {
	return f[key]
}

type Interface interface {
	reflect.Type
	GetType() reflect.Type

	IsSubClass(class any) bool

	// ClassName 获取类名
	ClassName() string
}

type Class[T any] interface {
	Interface

	// New 通过 Fields
	New(fields Fields) T

	NewByTag(fields Fields, tag string) T
}

type FieldsProvider interface {
	ToFields() Fields
}

type Json interface {
	ToJson() string
}

type Getter[T any] interface {
	Get(key string) T
	GetString(key string) string
	GetInt64(key string) int64
	GetInt32(key string) int32
	GetInt16(key string) int16
	GetInt8(key string) int8
	GetInt(key string) int
	GetUInt64(key string) uint64
	GetUInt32(key string) uint32
	GetUInt16(key string) uint16
	GetUInt8(key string) uint8
	GetUInt(key string) uint
	GetFloat64(key string) float64
	GetFloat(key string) float32
	GetBool(key string) bool
}

type GetterE[T any] interface {
	GetE(key string) (T, error)
	GetStringE(key string) (string, error)
	GetInt64E(key string) (int64, error)
	GetInt32E(key string) (int32, error)
	GetInt16E(key string) (int16, error)
	GetInt8E(key string) (int8, error)
	GetIntE(key string) (int, error)
	GetUInt64E(key string) (uint64, error)
	GetUInt32E(key string) (uint32, error)
	GetUInt16E(key string) (uint16, error)
	GetUInt8E(key string) (uint8, error)
	GetUIntE(key string) (uint, error)
	GetFloat64E(key string) (float64, error)
	GetFloatE(key string) (float32, error)
	GetBoolE(key string) (bool, error)
}

type OptionalGetter[T any] interface {
	Optional(key string, value T) T
	StringOptional(key string, defaultValue string) string
	Int64Optional(key string, defaultValue int64) int64
	Int32Optional(key string, defaultValue int32) int32
	Int16Optional(key string, defaultValue int16) int16
	Int8Optional(key string, defaultValue int8) int8
	IntOptional(key string, defaultValue int) int
	UInt64Optional(key string, defaultValue uint64) uint64
	UInt32Optional(key string, defaultValue uint32) uint32
	UInt16Optional(key string, defaultValue uint16) uint16
	UInt8Optional(key string, defaultValue uint8) uint8
	UIntOptional(key string, defaultValue uint) uint
	Float64Optional(key string, defaultValue float64) float64
	FloatOptional(key string, defaultValue float32) float32
	BoolOptional(key string, defaultValue bool) bool
}

type OptionalGetterE[T any] interface {
	OptionalE(key string, value T) (T, error)
	StringOptionalE(key string, defaultValue string) (string, error)
	Int64OptionalE(key string, defaultValue int64) (int64, error)
	Int32OptionalE(key string, defaultValue int32) (int32, error)
	Int16OptionalE(key string, defaultValue int16) (int16, error)
	Int8OptionalE(key string, defaultValue int8) (int8, error)
	IntOptionalE(key string, defaultValue int) (int, error)
	UInt64OptionalE(key string, defaultValue uint64) (uint64, error)
	UInt32OptionalE(key string, defaultValue uint32) (uint32, error)
	UInt16OptionalE(key string, defaultValue uint16) (uint16, error)
	UInt8OptionalE(key string, defaultValue uint8) (uint8, error)
	UIntOptionalE(key string, defaultValue uint) (uint, error)
	Float64OptionalE(key string, defaultValue float64) (float64, error)
	FloatOptionalE(key string, defaultValue float32) (float32, error)
	BoolOptionalE(key string, defaultValue bool) (bool, error)
}
