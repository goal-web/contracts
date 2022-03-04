package contracts

type BloomFactory interface {

	// Extend 扩展布隆过滤器驱动程序
	// Extended bloom filter driver.
	Extend(name string, driver BloomFilterDriver)

	// Filter 按名称获取布隆过滤器
	// Get bloom filter by name.
	Filter(name string) BloomFilter

	// Start 开启布隆过滤器驱动程序
	// Enable bloom filter driver.
	Start() error

	// Close 关闭布隆过滤器驱动程序
	// Turn off bloom filter driver.
	Close()
}

// BloomFilterDriver 通过名称与配置信息获取布隆过滤器
// Get bloom filter by name and configuration info.
type BloomFilterDriver func(name string, config Fields) BloomFilter

type BloomFilter interface {
	// Add 添加字节数组
	// add byte array.
	Add(bytes []byte)

	// AddString 添加字符串
	// add string .
	AddString(str string)


	// Test 测试字节数组
	// test byte array.
	Test(bytes []byte) bool
	// TestString 测试字符串
	// test string.
	TestString(str string) bool


	// TestOrAdd 相当于调用 test(bytes) 如果不存在 add(bytes)
	// is the equivalent to calling test(bytes) then if not present add(bytes).
	TestOrAdd(bytes []byte) bool
	// TestOrAddString 相当于调用 testString(str) 如果不存在 addString(str)
	// is the equivalent to calling testString(str) then if not present addString(str).
	TestOrAddString(str string) bool


	// TestAndAdd 相当于调用 test(bytes) 然后 add(bytes)
	// is the equivalent to calling test(bytes) then add(bytes).
	TestAndAdd(bytes []byte) bool
	// TestAndAddString 相当于调用 testString(string) 然后 addString(string)
	// is the equivalent to calling testString(string) then addString(string).
	TestAndAddString(str string) bool


	// Clear 清除布隆过滤器中的数据，删除键
	// clears the data in a Bloom filter, removing keys
	Clear()
	// Size  存储的项目数
	// number of items stored.
	Size() uint
	// Count 统计有多少位bit
	// count how many bits there are.
	Count() uint
	// Load 加载布隆过滤器
	// load Bloom filter.
	Load()
	// Save 保存布隆过滤器
	// save bloom filter.
	Save()
}
