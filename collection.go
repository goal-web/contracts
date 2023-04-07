package contracts

import "sort"

type Collection[T any] interface {
	Json
	// sort

	sort.Interface
	Sort(sorter func(int, int, T, T) bool) Collection[T]
	IsEmpty() bool

	// filter
	Foreach(handle func(int, T)) Collection[T]
	Each(handle func(int, T) T) Collection[T]
	Filter(filter func(int, T) bool) Collection[T]
	Skip(skipper func(int, T) bool) Collection[T]
	Map(filter any) Collection[T]

	Where(field string, args ...any) Collection[T]
	WhereLt(field string, arg any) Collection[T]
	WhereLte(field string, arg any) Collection[T]
	WhereGt(field string, arg any) Collection[T]
	WhereGte(field string, arg any) Collection[T]
	WhereIn(field string, arg any) Collection[T]
	WhereNotIn(field string, arg any) Collection[T]

	// keys、values

	// Pluck 数据类型为 []map、[]struct 的时候起作用
	Pluck(key string) map[string]T
	// GroupBy 数据类型为 []map、[]struct 的时候起作用
	GroupBy(key string) map[string][]T
	// Only 数据类型为 []map、[]struct 的时候起作用
	Only(keys ...string) Collection[T]

	// First 获取首个元素
	First() *T
	// Last 获取最后一个元素
	Last() *T

	// union、merge...

	// Prepend 从开头插入元素
	Prepend(item ...T) Collection[T]
	// Push 从最后插入元素
	Push(items ...T) Collection[T]
	// Pull 从尾部获取并移出一个元素
	Pull(defaultValue ...T) *T
	// Shift 从头部获取并移出一个元素
	Shift(defaultValue ...T) *T
	// Put 替换一个元素，如果 index 不存在会执行 Push，返回新集合
	Put(index int, item T) Collection[T]
	// Offset 替换一个元素，如果 index 不存在会执行 Push
	Offset(index int, item T) Collection[T]
	// Merge 合并其他集合
	Merge(collections ...Collection[T]) Collection[T]
	// Reverse 返回一个顺序翻转后的集合
	Reverse() Collection[T]
	// Chunk 分块，handler 返回 error 表示中断
	Chunk(size int, handler func(collection Collection[T], page int) error) error
	// Random 随机返回n个元素，默认1个
	Random(size ...uint) Collection[T]

	// aggregate

	Sum(key ...string) float64
	Max(key ...string) float64
	Min(key ...string) float64
	Avg(key ...string) float64
	Count() int

	// convert

	ToArray() []T
	ToAnyArray() []any
	ToArrayFields() []Fields
}
