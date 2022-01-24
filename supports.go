package contracts

import (
	"sort"
)

type Fields map[string]interface{}

type FieldsProvider interface {
	Fields() Fields
}

type Getter interface {
	GetString(key string) string
	GetInt64(key string) int64
	GetInt(key string) int
	GetFloat64(key string) float64
	GetFloat(key string) float32
	GetBool(key string) bool
	GetFields(key string) Fields
}

type OptionalGetter interface {
	StringOption(key string, defaultValue string) string
	Int64Option(key string, defaultValue int64) int64
	IntOption(key string, defaultValue int) int
	Float64Option(key string, defaultValue float64) float64
	FloatOption(key string, defaultValue float32) float32
	BoolOption(key string, defaultValue bool) bool
	FieldsOption(key string, defaultValue Fields) Fields
}

type Collection interface {
	// sort

	sort.Interface
	Sort(sorter interface{}) Collection
	IsEmpty() bool

	// filter

	Map(filter interface{}) Collection
	Filter(filter interface{}) Collection
	Skip(filter interface{}) Collection
	Where(field string, args ...interface{}) Collection
	WhereLt(field string, arg interface{}) Collection
	WhereLte(field string, arg interface{}) Collection
	WhereGt(field string, arg interface{}) Collection
	WhereGte(field string, arg interface{}) Collection
	WhereIn(field string, arg interface{}) Collection
	WhereNotIn(field string, arg interface{}) Collection

	// keys、values

	// Pluck 数据类型为 []map、[]struct 的时候起作用
	Pluck(key string) Fields
	// Only 数据类型为 []map、[]struct 的时候起作用
	Only(keys ...string) Collection

	// First 获取首个元素, []struct或者[]map可以获取指定字段
	First(keys ...string) interface{}
	// Last 获取最后一个元素, []struct或者[]map可以获取指定字段
	Last(keys ...string) interface{}

	// union、merge...

	// Prepend 从开头插入元素
	Prepend(item ...interface{}) Collection
	// Push 从最后插入元素
	Push(items ...interface{}) Collection
	// Pull 从尾部获取并移出一个元素
	Pull(defaultValue ...interface{}) interface{}
	// Shift 从头部获取并移出一个元素
	Shift(defaultValue ...interface{}) interface{}
	// Put 替换一个元素，如果 index 不存在会执行 Push，返回新集合
	Put(index int, item interface{}) Collection
	// Offset 替换一个元素，如果 index 不存在会执行 Push
	Offset(index int, item interface{}) Collection
	// Merge 合并其他集合
	Merge(collections ...Collection) Collection
	// Reverse 返回一个顺序翻转后的集合
	Reverse() Collection
	// Chunk 分块，handler 返回 error 表示中断
	Chunk(size int, handler func(collection Collection, page int) error) error
	// Random 随机返回n个元素，默认1个
	Random(size ...uint) Collection

	// aggregate

	Sum(key ...string) float64
	Max(key ...string) float64
	Min(key ...string) float64
	Avg(key ...string) float64
	Count() int

	// convert

	ToIntArray() (results []int)
	ToInt64Array() (results []int64)
	ToInterfaceArray() []interface{}
	ToFloat64Array() (results []float64)
	ToFloatArray() (results []float32)
	ToBoolArray() (results []bool)
	ToStringArray() (results []string)
	ToFields() Fields
	ToArrayFields() []Fields
}
