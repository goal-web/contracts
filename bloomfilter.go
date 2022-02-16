package contracts

type BloomFactory interface {
	Extend(name string, driver BloomFilterDriver)

	Filter(name string) BloomFilter

	Start() error
	Close()
}

type BloomFilterDriver func(config Fields) BloomFilter

type BloomFilter interface {
	Add(bytes []byte)
	AddString(str string)

	Test(bytes []byte) bool
	TestString(str string) bool

	TestOrAdd(bytes []byte) bool
	TestOrAddString(str string) bool

	TestAndAdd(bytes []byte) bool
	TestAndAddString(str string) bool

	Clear()
	Size() uint
	Count() uint
	Load()
	Save()
}
