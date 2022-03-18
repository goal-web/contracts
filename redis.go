package contracts

import (
	"time"
)

type RedisFactory interface {
	// Connection 解析一个redis连接实例
	// Resolve a redis connection instance.
	Connection(name ...string) RedisConnection
}

type GeoPos struct {
	Longitude, Latitude float64
}

type BitCount struct {
	Start, End int64
}

type GeoLocation struct {
	Name                      string
	Longitude, Latitude, Dist float64
	GeoHash                   int64
}

type GeoRadiusQuery struct {
	Radius float64
	// Can be m, km, ft, or mi. Default is km.
	Unit        string
	WithCoord   bool
	WithDist    bool
	WithGeoHash bool
	Count       int
	// Can be ASC or DESC. Default is no sort order.
	Sort      string
	Store     string
	StoreDist string
}

type ZStore struct {
	Keys    []string
	Weights []float64
	// Can be SUM, MIN or MAX.
	Aggregate string
}

type Z struct {
	Score  float64
	Member interface{}
}

type ZRangeBy struct {
	Min, Max      string
	Offset, Count int64
}

// RedisSubscribeFunc 订阅给定的消息频道
// Subscribe to a given message channel.
type RedisSubscribeFunc func(message, channel string)

type RedisConnection interface {
	// Subscribe 订阅一组给定的消息频道
	// subscribe to a set of given channels for messages.
	Subscribe(channels []string, closure RedisSubscribeFunc) error

	// PSubscribe 使用通配符订阅一组给定频道
	// subscribe to a set of given channels with wildcards.
	PSubscribe(channels []string, closure RedisSubscribeFunc) error

	// Command 对 Redis 数据库运行命令
	// Run a command against the Redis database.
	Command(method string, args ...interface{}) (interface{}, error)

	// PubSubChannels 列出当前active channels.活跃是指信道含有一个或多个订阅者(不包括从模式接收订阅的客户端) 如果pattern未提供，所有的信道都被列出，否则只列出匹配上指定全局-类型模式的信道被列出.
	// List the currently active channels. Active means that the channel contains one or more subscribers (excluding clients receiving subscriptions from the pattern). If pattern is not provided, all channels are listed, otherwise only lists matching the specified global- Channels of type mode are listed.
	PubSubChannels(pattern string) ([]string, error)


	// PubSubNumSub 列出指定信道的订阅者个数(不包括订阅模式的客户端订阅者)
	// List the number of subscribers to the specified channel (excluding client subscribers in subscription mode).
	PubSubNumSub(channels ...string) (map[string]int64, error)

	// PubSubNumPat 返回订阅模式的数量(使用命令PSUBSCRIBE实现).注意， 这个命令返回的不是订阅模式的客户端的数量， 而是客户端订阅的所有模式的数量总和。
	// Returns the number of subscribed schemas (implemented using the command PSUBSCRIBE). Note that this command returns not the number of clients subscribed to the schema, but the sum of all schemas subscribed by the client.
	PubSubNumPat() (int64, error)

	// Publish 将信息 message 发送到指定的频道 channel
	// Send the information message to the specified channel channel.
	Publish(channel string, message interface{}) (int64, error)

	// Get 返回给定键的值
	// Returns the value of the given key.
	Get(key string) (string, error)

	// MGet 获取所有给定键的值
	// get the values of all the given keys.
	MGet(keys ...string) ([]interface{}, error)

	// GetBit 对 key 所储存的字符串值，对获取指定偏移量上的位(bit)
	// For the string value stored in key, get the bit at the specified offset.
	GetBit(key string, offset int64) (int64, error)

	// BitOpAnd 对一个或多个 key 求逻辑并，并将结果保存到 destkey
	// Take a logical union of one or more keys and save the result to destkey.
	BitOpAnd(destKey string, keys ...string) (int64, error)

	// BitOpNot 对给定 key 求逻辑非，并将结果保存到 destkey
	// Take the logical negation of the given key and save the result to destkey.
	BitOpNot(destKey string, key string) (int64, error)

	// BitOpOr 对一个或多个 key 求逻辑或，并将结果保存到 destkey
	// Logical OR of one or more keys and save the result to destkey.
	BitOpOr(destKey string, keys ...string) (int64, error)

	// BitOpXor 对一个或多个 key 求逻辑异或，并将结果保存到 destkey
	// XOR one or more keys and save the result to destkey.
	BitOpXor(destKey string, keys ...string) (int64, error)

	// GetDel 返回key对应的字符串value的字符串，并删除key
	// Returns the string value of the string corresponding to the key, and deletes the key.
	GetDel(key string) (string, error)

	// GetEx 零到期会删除与密钥关联的 TTL（即 GETEX 密钥持久）
	// An expiration of zero removes the TTL associated with the key (i.e. GETEX key persist).
	GetEx(key string, expiration time.Duration) (string, error)

	// GetRange 返回key对应的字符串value的子串，这个子串是由start和end位移决定的
	// Returns the substring of the string value corresponding to the key, which is determined by the start and end displacements.
	GetRange(key string, start, end int64) (string, error)

	// GetSet 自动将key对应到value并且返回原来key对应的value。如果key存在但是对应的value不是字符串，就返回错误
	// Automatically map the key to the value and return the value corresponding to the original key. If the key exists but the corresponding value is not a string, return an error.
	GetSet(key string, value interface{}) (string, error)

	// ClientGetName 返回连接的名称
	// returns the name of the connection.
	ClientGetName() (string, error)

	// StrLen 返回key的string类型value的长度。如果key对应的非string类型，就返回错误
	// Returns the length of the string value of key. If the key corresponds to a non-string type, an error is returned.
	StrLen(key string) (int64, error)

	// getter end
	// keys start

	// Keys 查找所有符合给定模式pattern（正则表达式）的 key
	// Find all keys matching the given pattern pattern (regular expression).
	Keys(pattern string) ([]string, error)

	// Del 删除指定的一批keys，如果删除中的某些key不存在，则直接忽略
	// Delete the specified batch of keys, if some keys in the deletion do not exist, they will be ignored directly.
	Del(keys ...string) (int64, error)

	// FlushAll 删除所有数据库里面的所有数据，注意不是当前数据库，而是所有数据库
	// Delete all data in all databases, note that not the current database, but all databases.
	FlushAll() (string, error)

	// FlushDB 删除当前数据库里面的所有数据
	// Delete all data in the current database.
	FlushDB() (string, error)

	// Dump 序列化给定 key ，并返回被序列化的值
	// Serialize the given key and return the serialized value.
	Dump(key string) (string, error)

	// Exists 返回key是否存在
	// Returns whether the key exists.
	Exists(keys ...string) (int64, error)

	// Expire 设置key的过期时间，超过时间后，将会自动删除该key
	// Set the expiration time of the key. After the time expires, the key will be automatically deleted.
	Expire(key string, expiration time.Duration) (bool, error)

	// ExpireAt EXPIREAT 的作用和 EXPIRE类似，都用于为 key 设置生存时间。不同在于 EXPIREAT 命令接受的时间参数是 UNIX 时间戳
	// The role of EXPIREAT is similar to that of EXPIRE, both are used to set the time-to-live for the key. The difference is that the time parameter accepted by the EXPIREAT command is the UNIX timestamp.
	ExpireAt(key string, tm time.Time) (bool, error)

	// PExpire 这个命令和EXPIRE命令的作用类似，但是它以毫秒为单位设置 key 的生存时间，而不像EXPIRE命令那样，以秒为单位
	// This command works like the EXPIRE command, but it sets the key's lifetime in milliseconds instead of seconds like the EXPIRE command.
	PExpire(key string, expiration time.Duration) (bool, error)

	// PExpireAt PEXPIREAT 这个命令和EXPIREAT命令类似，但它以毫秒为单位设置 key 的过期 unix 时间戳，而不是像EXPIREAT那样，以秒为单位
	// PEXPIREAT This command is similar to the EXPIREAT command, but it sets the expiry unix timestamp of the key in milliseconds instead of seconds like EXPIREAT.
	PExpireAt(key string, tm time.Time) (bool, error)

	// Migrate 将 key 原子性地从当前实例传送到目标实例的指定数据库上，一旦传送成功， key 保证会出现在目标实例上，而当前实例上的 key 会被删除
	// Atomically transfer the key from the current instance to the specified database of the target instance. Once the transfer is successful, the key is guaranteed to appear on the target instance, and the key on the current instance will be deleted.
	Migrate(host, port, key string, db int, timeout time.Duration) (string, error)

	// Move 将当前数据库的 key 移动到给定的数据库 db 当中
	// move the key of the current database to the given database db.
	Move(key string, db int) (bool, error)

	// Persist 移除给定key的生存时间，将这个 key 从『易失的』(带生存时间 key )转换成『持久的』(一个不带生存时间、永不过期的 key )
	// Removes the time-to-live for a given key, converting the key from "volatile" (a key with a time-to-live) to "durable" (a key with no time-to-live that never expires).
	Persist(key string) (bool, error)

	// PTTL 这个命令类似于TTL命令，但它以毫秒为单位返回 key 的剩余生存时间，而不是像TTL命令那样，以秒为单位
	// This command is similar to the TTL command, but it returns the remaining time to live for the key in milliseconds instead of seconds like the TTL command.
	PTTL(key string) (time.Duration, error)

	// TTL 返回key剩余的过期时间。 这种反射能力允许Redis客户端检查指定key在数据集里面剩余的有效期
	// Returns the remaining expiration time for the key. This reflection capability allows Redis clients to check the remaining validity period of a given key in the dataset.
	TTL(key string) (time.Duration, error)

	// RandomKey 从当前数据库返回一个随机的key
	// Returns a random key from the current database.
	RandomKey() (string, error)

	// Rename 将key重命名为newkey，如果key与newkey相同，将返回一个错误。如果newkey已经存在，则值将被覆盖
	// Rename key to newkey, returns an error if key is the same as newkey. If newkey already exists, the value will be overwritten.
	Rename(key, newKey string) (string, error)

	// RenameNX 当且仅当 newkey 不存在时，将 key 改名为 newkey, 当 key 不存在时，返回一个错误
	// If and only if newkey does not exist, rename key to newkey, if key does not exist, return an error.
	RenameNX(key, newKey string) (bool, error)

	// Type 返回key所存储的value的数据结构类型，它可以返回string, list, set, zset 和 hash等不同的类型
	// Returns the data structure type of the value stored by the key, which can return different types such as string, list, set, zset and hash.
	Type(key string) (string, error)

	// Wait 此命令阻塞当前客户端，直到所有以前的写命令都成功的传输和指定的slaves确认。如果超时，指定以毫秒为单位，即使指定的slaves还没有到达，命令任然返回
	// This command blocks the current client until all previous write commands have been successfully transmitted and acknowledged by the specified slaves. If it times out, specify in milliseconds, even if the specified slaves have not arrived, the command still returns.
	Wait(numSlaves int, timeout time.Duration) (int64, error)

	// Scan 迭代当前数据库中的key集合
	// Iterate over the key collection in the current database.
	Scan(cursor uint64, match string, count int64) ([]string, uint64, error)

	// BitCount 统计字符串被设置为1的bit数
	// The number of bits whose statistics string is set to 1.
	BitCount(key string, count *BitCount) (int64, error)

	// keys end

	// setter start
	// Set 将键key设定为指定的“字符串”值
	// Set the key key to the specified "string" value.
	Set(key string, value interface{}, expiration time.Duration) (string, error)

	// Append 如果 key 已经存在，并且值为字符串，那么这个命令会把 value 追加到原来值（value）的结尾。 如果 key 不存在，那么它将首先创建一个空字符串的key，再执行追加操作，这种情况 APPEND 将类似于 SET 操作
	// If key already exists and the value is a string, then this command appends value to the end of the original value. If the key does not exist, then it will first create a key with an empty string, and then perform the append operation, in which case APPEND will be similar to the SET operation.
	Append(key, value string) (int64, error)

	// MSet 对应给定的keys到他们相应的values上。MSET会用新的value替换已经存在的value，就像普通的SET命令一样。如果你不想覆盖已经存在的values，请参看命令MSETNX。
	// Map the given keys to their corresponding values. MSET will replace the existing value with the new value, just like a normal SET command. If you do not want to overwrite existing values, see the command MSETNX.
	MSet(values ...interface{}) (string, error)

	// MSetNX 对应给定的keys到他们相应的values上。只要有一个key已经存在，MSETNX一个操作都不会执行。 由于这种特性，MSETNX可以实现要么所有的操作都成功，要么一个都不执行，这样可以用来设置不同的key，来表示一个唯一的对象的不同字段。
	// Map the given keys to their corresponding values. As long as a key already exists, MSETNX will not perform an operation. Because of this feature, MSETNX can achieve either all operations succeed or none of them are executed, which can be used to set different keys to represent different fields of a unique object.
	MSetNX(values ...interface{}) (bool, error)

	// SetNX 将key设置值为value，如果key不存在，这种情况下等同SET命令。 当key存在时，什么也不做, SETNX是”SET if Not eXists”的简写。
	// Set the key to value, if the key does not exist, this case is equivalent to the SET command. When the key exists, do nothing, SETNX is short for "SET if Not eXists".
	SetNX(key string, value interface{}, expiration time.Duration) (bool, error)

	// SetEX 设置key对应字符串value，并且设置key在给定的seconds时间之后超时过期
	// Set the key corresponding to the string value, and set the key to expire after the given seconds time.
	SetEX(key string, value interface{}, expiration time.Duration) (string, error)

	// SetBit 设置或者清空key的value(字符串)在offset处的bit值
	// set or clear the bit value of the key's value (string) at offset.
	SetBit(key string, offset int64, value int) (int64, error)

	// BitPos 返回字符串里面第一个被设置为1或者0的bit位
	// Returns the first bit in the string that is set to 1 or 0.
	BitPos(key string, bit int64, pos ...int64) (int64, error)

	// SetRange 这个命令的作用是覆盖key对应的string的一部分，从指定的offset处开始，覆盖value的长度。如果offset比当前key对应string还要长，那这个string后面就补0以达到offset。不存在的keys被认为是空字符串，所以这个命令可以确保key有一个足够大的字符串，能在offset处设置value
	// The function of this command is to overwrite the part of the string corresponding to the key, starting from the specified offset, overwriting the length of the value. If the offset is longer than the string corresponding to the current key, then the string is followed by 0 to achieve the offset. Non-existing keys are considered empty strings, so this command ensures that the key has a string large enough to set the value at offset.
	SetRange(key string, offset int64, value string) (int64, error)

	// Incr 对存储在指定key的数值执行原子的加1操作。如果指定的key不存在，那么在执行incr操作之前，会先将它的值设定为0。
	// Performs an atomic increment operation on the value stored at the specified key. If the specified key does not exist, its value will be set to 0 before performing the incr operation.
	Incr(key string) (int64, error)

	// Decr 对key对应的数字做减1操作。如果key不存在，那么在操作之前，这个key对应的值会被置为0。如果key有一个错误类型的value或者是一个不能表示成数字的字符串，就返回错误。这个操作最大支持在64位有符号的整型数字。
	// Subtract 1 from the number corresponding to the key. If the key does not exist, the value corresponding to the key will be set to 0 before the operation. Returns an error if key has a value of the wrong type or is a string that cannot be represented as a number. This operation supports up to 64-bit signed integer numbers.
	Decr(key string) (int64, error)

	// IncrBy 将key对应的数字加decrement。如果key不存在，操作之前，key就会被置为0。如果key的value类型错误或者是个不能表示成数字的字符串，就返回错误。这个操作最多支持64位有符号的正型数字。
	// Add decrement to the number corresponding to the key. If the key does not exist, the key will be set to 0 before the operation. Returns an error if the value of key has the wrong type or is a string that cannot be represented as a number. This operation supports up to 64-bit signed positive numbers.
	IncrBy(key string, value int64) (int64, error)

	// DecrBy 将key对应的数字减decrement。如果key不存在，操作之前，key就会被置为0。如果key的value类型错误或者是个不能表示成数字的字符串，就返回错误。这个操作最多支持64位有符号的正型数字。
	// Subtract decrement from the number corresponding to the key. If the key does not exist, the key will be set to 0 before the operation. Returns an error if the value of key has the wrong type or is a string that cannot be represented as a number. This operation supports up to 64-bit signed positive numbers.
	DecrBy(key string, value int64) (int64, error)

	// IncrByFloat 通过指定浮点数key来增长浮点数(存放于string中)的值. 当键不存在时,先将其值设为0再操作.
	// Increase the value of the floating-point number (stored in the string) by specifying the floating-point number key. When the key does not exist, set its value to 0 before operating.
	IncrByFloat(key string, value float64) (float64, error)

	// setter end

	// hash start
	// HGet 返回 key 指定的哈希集中该字段所关联的值
	// Returns the value associated with the field in the hash set specified by key.
	HGet(key, field string) (string, error)

	// HGetAll 返回 key 指定的哈希集中所有的字段和值。返回值中，每个字段名的下一个是它的值，所以返回值的长度是哈希集大小的两倍
	// Returns all fields and values in the hash set specified by key. In the return value, next to each field name is its value, so the length of the return value is twice the size of the hash set.
	HGetAll(key string) (map[string]string, error)

	// HMGet 返回 key 指定的哈希集中指定字段的值。
	// Returns the value of the specified field in the hash set specified by key.
	HMGet(key string, fields ...string) ([]interface{}, error)

	// HKeys 返回 key 指定的哈希集中所有字段的名字。
	// Returns the names of all fields in the hash set specified by key.
	HKeys(key string) ([]string, error)

	// HLen 返回 key 指定的哈希集包含的字段的数量。
	// Returns the number of fields contained in the hashset specified by key.
	HLen(key string) (int64, error)

	// HRandField Redis 6.2中新增的命令，用于随机获取指定哈希表中的域。
	// A new command added in Redis 6.2 to randomly obtain the fields in the specified hash table.
	HRandField(key string, count int, withValues bool) ([]string, error)

	// HScan 用于增量式的迭代获取哈希表中的所有域，并返回其域名称及其值。
	// Used for incremental iteration to get all fields in the hash table and return their field names and their values.
	HScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error)

	// HValues 返回 key 指定的哈希集中所有字段的值。
	// Returns the values of all fields in the hashset specified by key.
	HValues(key string) ([]string, error)

	// HSet 设置 key 指定的哈希集中指定字段的值。
	// Sets the value of the specified field in the hash set specified by key.
	HSet(key string, values ...interface{}) (int64, error)

	// HSetNX 只在 key 指定的哈希集中不存在指定的字段时，设置字段的值。如果 key 指定的哈希集不存在，会创建一个新的哈希集并与 key 关联。如果字段已存在，该操作无效果。
	// Sets the field's value only if the specified field does not exist in the hash set specified by key. If the hash set specified by key does not exist, a new hash set is created and associated with key. If the field already exists, this operation has no effect.
	HSetNX(key, field string, value interface{}) (bool, error)

	// HMSet 设置 key 指定的哈希集中指定字段的值。该命令将重写所有在哈希集中存在的字段。如果 key 指定的哈希集不存在，会创建一个新的哈希集并与 key 关联
	// Sets the value of the specified field in the hash set specified by key. This command will rewrite all fields present in the hashset. If the hash set specified by key does not exist, a new hash set will be created and associated with key.
	HMSet(key string, values ...interface{}) (bool, error)

	// HDel 从 key 指定的哈希集中移除指定的域。在哈希集中不存在的域将被忽略。
	// Removes the specified domain from the hash set specified by key. Domains that do not exist in the hashset will be ignored.
	HDel(key string, fields ...string) (int64, error)

	// HExists 返回hash里面field是否存在
	// Returns whether the field exists in the hash.
	HExists(key string, field string) (bool, error)

	// HIncrBy 增加 key 指定的哈希集中指定字段的数值。如果 key 不存在，会创建一个新的哈希集并与 key 关联。如果字段不存在，则字段的值在该操作执行前被设置为 0 HINCRBY 支持的值的范围限定在 64位 有符号整数
	// Increments the value of the specified field in the hash set specified by key. If key does not exist, a new hash set is created and associated with key. If the field does not exist, the value of the field is set to 0 before the operation is performed HINCRBY The range of supported values is limited to 64-bit signed integers.
	HIncrBy(key string, field string, value int64) (int64, error)

	// HIncrByFloat 为指定key的hash的field字段值执行float类型的increment加。如果field不存在，则在执行该操作前设置为0.
	// Performs an increment of float type for the field field value of the hash of the specified key. If field does not exist, set to 0 before performing the operation.
	HIncrByFloat(key string, field string, value float64) (float64, error)

	// hash end

	// set start
	// SAdd 添加一个或多个指定的member元素到集合的 key中.指定的一个或者多个元素member 如果已经在集合key中存在则忽略.如果集合key 不存在，则新建集合key,并添加member元素到集合key中.如果key 的类型不是集合则返回错误.
	// Add one or more specified member elements to the set key. If the specified one or more member elements already exist in the set key, it will be ignored. If the set key does not exist, create a new set key and add the member element to the set key. If the type of key is not a collection, an error is returned.
	SAdd(key string, members ...interface{}) (int64, error)

	// SCard 返回集合存储的key的基数 (集合元素的数量).
	// Returns the cardinality of the key stored in the collection (the number of elements in the collection).
	SCard(key string) (int64, error)

	// SDiff 返回一个集合与给定集合的差集的元素.
	// Returns the elements of the difference between a set and the given set.
	SDiff(keys ...string) ([]string, error)

	// SDiffStore 该命令类似于 SDIFF, 不同之处在于该命令不返回结果集，而是将结果存放在destination集合中.如果destination已经存在, 则将其覆盖重写.
	// This command is similar to SDIFF, except that the command does not return a result set, but stores the result in the destination set. If the destination already exists, it will.
	SDiffStore(destination string, keys ...string) (int64, error)

	// SInter 返回指定所有的集合的成员的交集.
	// Returns the intersection of all members of the specified set.
	SInter(keys ...string) ([]string, error)

	// SInterStore 这个命令与SINTER命令类似, 但是它并不是直接返回结果集,而是将结果保存在 destination集合中.如果destination 集合存在, 则会被重写.
	// This command is similar to the SINTER command, but instead of returning the result set directly, it stores the result in the destination collection. If the destination collection exists, it will be overwritten.
	SInterStore(destination string, keys ...string) (int64, error)

	// SIsMember 返回成员 member 是否是存储的集合 key的成员.
	// Returns whether member member is a member of the stored collection key.
	SIsMember(key string, member interface{}) (bool, error)

	// SMembers 返回key集合所有的元素.
	// Returns all elements of the key collection.
	SMembers(key string) ([]string, error)

	// SRem 在key集合中移除指定的元素. 如果指定的元素不是key集合中的元素则忽略 如果key集合不存在则被视为一个空的集合，该命令返回0. 如果key的类型不是一个集合,则返回错误.
	// Removes the specified element from the key set. If the specified element is not an element in the key set, it is ignored. If the key set does not exist, it is treated as an empty set, and the command returns 0. If the type of key is not a set, then returns an error.
	SRem(key string, members ...interface{}) (int64, error)

	// SPopN 从存储在key的集合中移除并返回多个随机元素
	// Redis `SPOP key count` command.Remove and return multiple random elements from the collection stored at key.
	SPopN(key string, count int64) ([]string, error)

	// SPop 从存储在key的集合中移除并返回一个。
	// Remove from the collection stored at key and return a.
	SPop(key string) (string, error)

	// SRandMemberN 随机返回key集合中的多个元素.
	// Redis `SRANDMEMBER key count` command. Randomly returns multiple elements in the key collection.
	SRandMemberN(key string, count int64) ([]string, error)

	// SMove 将member从source集合移动到destination集合中. 对于其他的客户端,在特定的时间元素将会作为source或者destination集合的成员出现.
	// Moves the member from the source collection to the destination collection. For other clients, the element will appear as a member of the source or destination collection at a specific time.
	SMove(source, destination string, member interface{}) (bool, error)

	// SRandMember 随机返回key集合中的一个元素.
	// Randomly returns an element in the key collection.
	SRandMember(key string) (string, error)

	// SUnion 返回给定的多个集合的并集中的所有成员.
	// Returns all members of the union of the given collections.
	SUnion(keys ...string) ([]string, error)

	// SUnionStore 该命令作用类似于SUNION命令,不同的是它并不返回结果集,而是将结果存储在destination集合中.如果destination 已经存在,则将其覆盖.
	// his command works like the SUNION command, except that it does not return a result set, but stores the result in the destination set. If the destination already exists, it will be overwritten.
	SUnionStore(destination string, keys ...string) (int64, error)

	// set end

	// geo start
	// GeoAdd 将指定的地理空间位置（纬度、经度、名称）添加到指定的key中。
	// Adds the specified geospatial location (latitude, longitude, name) to the specified key.
	GeoAdd(key string, geoLocation ...*GeoLocation) (int64, error)

	// GeoHash 返回一个或多个位置元素的 Geohash 表示。
	// Returns a Geohash representation of one or more location elements.
	GeoHash(key string, members ...string) ([]string, error)

	// GeoPos 从key里返回所有给定位置元素的位置（经度和纬度）。
	// Returns the location (longitude and latitude) of all given location elements from key.
	GeoPos(key string, members ...string) ([]*GeoPos, error)

	// GeoDist 返回两个给定位置之间的距离。如果两个位置之间的其中一个不存在， 那么命令返回空值。
	// Returns the distance between two given locations. If one of the two positions does not exist, the command returns null.
	GeoDist(key string, member1, member2, unit string) (float64, error)

	// GeoRadius 以给定的经纬度为中心， 返回键包含的位置元素当中， 与中心的距离不超过给定最大距离的所有位置元素。
	// Taking the given latitude and longitude as the center, among the position elements contained in the return key, all position elements whose distance from the center does not exceed the given maximum distance.
	GeoRadius(key string, longitude, latitude float64, query *GeoRadiusQuery) ([]GeoLocation, error)

	// GeoRadiusStore 是一个写入 GEORADIUS 命令。
	// is a writing GEORADIUS command.
	GeoRadiusStore(key string, longitude, latitude float64, query *GeoRadiusQuery) (int64, error)

	// GeoRadiusByMember 这个命令和 GEORADIUS 命令一样， 都可以找出位于指定范围内的元素， 但是 GEORADIUSBYMEMBER 的中心点是由给定的位置元素决定的， 而不是像 GEORADIUS 那样， 使用输入的经度和纬度来决定中心点 指定成员的位置被用作查询的中心。
	// This command is the same as the GEORADIUS command, which can find the elements within the specified range, but the center point of GEORADIUSBYMEMBER is determined by the given position element, instead of using the input latitude and longitude to determine the center point specification like GEORADIUS The member's location is used as the center of the query.
	GeoRadiusByMember(key, member string, query *GeoRadiusQuery) ([]GeoLocation, error)

	// GeoRadiusByMemberStore 是一个写入 GEORADIUSBYMEMBER 命令。
	// is a writing GEORADIUSBYMEMBER command.
	GeoRadiusByMemberStore(key, member string, query *GeoRadiusQuery) (int64, error)

	// geo end

	// lists start
	// BLPop BLPOP 是阻塞式列表的弹出原语。 它是命令 LPOP 的阻塞版本，这是因为当给定列表内没有任何元素可供弹出的时候， 连接将被 BLPOP 命令阻塞。 当给定多个 key 参数时，按参数 key 的先后顺序依次检查各个列表，弹出第一个非空列表的头元素。
	// BLPOP is a popping primitive for blocking lists. It is the blocking version of the command LPOP because the connection will be blocked by the BLPOP command when there are no elements to pop in the given list. When multiple key parameters are given, each list is checked in order of the parameter keys, and the head element of the first non-empty list is popped.
	BLPop(timeout time.Duration, keys ...string) ([]string, error)

	// BRPop BRPOP 是一个阻塞的列表弹出原语。 它是 RPOP 的阻塞版本，因为这个命令会在给定list无法弹出任何元素的时候阻塞连接。 该命令会按照给出的 key 顺序查看 list，并在找到的第一个非空 list 的尾部弹出一个元素。
	// BRPOP is a blocking list pop primitive. It is a blocking version of RPOP, as this command blocks the connection if no elements can be popped from the given list. This command looks through the list in the order given by the keys and pops an element at the end of the first non-empty list found.
	BRPop(timeout time.Duration, keys ...string) ([]string, error)

	// BRPopLPush BRPOPLPUSH 是 RPOPLPUSH 的阻塞版本。 当 source 包含元素的时候，这个命令表现得跟 RPOPLPUSH 一模一样。 当 source 是空的时候，Redis将会阻塞这个连接，直到另一个客户端 push 元素进入或者达到 timeout 时限。 timeout 为 0 能用于无限期阻塞客户端。
	// BRPOPLPUSH is a blocking version of RPOPLPUSH. When source contains elements, this command behaves exactly like RPOPLPUSH. When source is empty, Redis will block the connection until another client pushes an element in or the timeout expires. A timeout of 0 can be used to block the client indefinitely.
	BRPopLPush(source, destination string, timeout time.Duration) (string, error)

	// LIndex 返回列表里的元素的索引 index 存储在 key 里面。 下标是从0开始索引的，所以 0 是表示第一个元素， 1 表示第二个元素，并以此类推。 负数索引用于指定从列表尾部开始索引的元素。在这种方法下，-1 表示最后一个元素，-2 表示倒数第二个元素，并以此往前推。当 key 位置的值不是一个列表的时候，会返回一个error。
	// Returns the index of the element in the list index stored in key. Subscripts are 0-based, so 0 is the first element, 1 is the second, and so on. Negative indices are used to specify elements indexed from the end of the list. In this method, -1 means the last element, -2 means the second-to-last element, and so on. An error is returned when the value at the key position is not a list.
	LIndex(key string, index int64) (string, error)

	// LInsert 把 value 插入存于 key 的列表中在基准值 pivot 的前面或后面。
	// Inserts the value in the list stored at key before or after the pivot value.
	LInsert(key, op string, pivot, value interface{}) (int64, error)

	// LLen 返回存储在 key 里的list的长度。 如果 key 不存在，那么就被看作是空list，并且返回长度为 0。 当存储在 key 里的值不是一个list的话，会返回error。
	// Returns the length of the list stored in key. If the key does not exist, it is treated as an empty list, and the return length is 0. An error is returned when the value stored in key is not a list.
	LLen(key string) (int64, error)

	// LPop 移除并且返回 key 对应的 list 的第一个元素。
	// Removes and returns the first element of the list corresponding to key.
	LPop(key string) (string, error)

	// LPush 将所有指定的值插入到存于 key 的列表的头部。如果 key 不存在，那么在进行 push 操作前会创建一个空列表。 如果 key 对应的值不是一个 list 的话，那么会返回一个错误。
	// Inserts all specified values at the head of the list stored at key. If the key does not exist, an empty list is created before the push operation. If the value corresponding to key is not a list, an error will be returned.
	LPush(key string, values ...interface{}) (int64, error)

	// LPushX 只有当 key 已经存在并且存着一个 list 的时候，在这个 key 下面的 list 的头部插入 value。 与 LPUSH 相反，当 key 不存在的时候不会进行任何操作。
	// Insert value at the head of the list below the key only if the key already exists and a list exists. Contrary to LPUSH, no operation is performed when the key does not exist.
	LPushX(key string, values ...interface{}) (int64, error)

	// LRange 返回存储在 key 的列表里指定范围内的元素。
	// Returns the elements in the specified range stored in the list at key.
	LRange(key string, start, stop int64) ([]string, error)

	// LRem 从存于 key 的列表里移除前 count 次出现的值为 value 的元素
	// Removes the first count occurrences of the element whose value is value from the list stored at key.
	LRem(key string, count int64, value interface{}) (int64, error)

	// LSet 设置 index 位置的list元素的值为 value。
	// Sets the value of the list element at index position to value.
	LSet(key string, index int64, value interface{}) (string, error)

	// LTrim 修剪(trim)一个已存在的 list，这样 list 就会只包含指定范围的指定元素。
	// Trim an existing list so that the list contains only the specified elements in the specified range.
	LTrim(key string, start, stop int64) (string, error)

	// RPop 推出并返回存于 key 的 list 的最后一个元素。
	// Push out and return the last element of the list stored at key.
	RPop(key string) (string, error)

	// RPopCount 推出右边指定数量的元素并返回存于 key 的 list。
	// Extracts the specified number of elements on the right and returns the list stored at key.
	RPopCount(key string, count int) ([]string, error)

	// RPopLPush 原子性地返回并移除存储在 source 的列表的最后一个元素（列表尾部元素）， 并把该元素放入存储在 destination 的列表的第一个元素位置（列表头部）。
	// Atomically returns and removes the last element of the list stored in source (the tail of the list) and places that element in the position of the first element of the list stored in destination (the head of the list).
	RPopLPush(source, destination string) (string, error)

	// RPush 向存于 key 的列表的尾部插入所有指定的值。如果 key 不存在，那么会创建一个空的列表然后再进行 push 操作。 当 key 保存的不是一个列表，那么会返回一个错误。
	// Inserts all specified values to the end of the list stored at key. If the key does not exist, an empty list is created and then the push operation is performed. When key does not hold a list, an error is returned.
	RPush(key string, values ...interface{}) (int64, error)

	// RPushX 将值 value 插入到列表 key 的表尾, 当且仅当 key 存在并且是一个列表。 和 RPUSH 命令相反, 当 key 不存在时，RPUSHX 命令什么也不做。
	// Inserts the value value at the end of the list key if and only if key exists and is a list. Contrary to the RPUSH command, the RPUSHX command does nothing when the key does not exist.
	RPushX(key string, values ...interface{}) (int64, error)

	// lists end

	// scripting start

	// Eval 在服务器端执行 LUA 脚本。
	// Execute LUA scripts on the server side.
	Eval(script string, keys []string, args ...interface{}) (interface{}, error)

	// EvalSha 根据给定的 SHA1 校验码，对缓存在服务器中的脚本进行求值。 将脚本缓存到服务器的操作可以通过 SCRIPT LOAD 命令进行。 这个命令的其他地方，比如参数的传入方式，都和 EVAL命令一样。
	// Evaluate the script cached in the server against the given SHA1 checksum. Caching scripts to the server can be done with the SCRIPT LOAD command. The rest of this command, such as the way parameters are passed in, are the same as the EVAL command.
	EvalSha(sha1 string, keys []string, args ...interface{}) (interface{}, error)

	// ScriptExists 检查脚本是否存在脚本缓存里面。
	// Check if the script exists in the script cache.
	ScriptExists(hashes ...string) ([]bool, error)

	// ScriptFlush 清空Lua脚本缓存
	// Flush the Lua scripts cache.
	ScriptFlush() (string, error)

	// ScriptKill 杀死当前正在运行的 Lua 脚本，当且仅当这个脚本没有执行过任何写操作时，这个命令才生效。
	// Kills the currently running Lua script, if and only if the script has not performed any write operations, this command will take effect.
	ScriptKill() (string, error)

	// ScriptLoad 将脚本 script 添加到脚本缓存中，但并不立即执行该脚本。在脚本被加入到缓存之后，通过 EVALSHA 命令，可以使用脚本的 SHA1 校验和来调用这个脚本。
	// Adds the script script to the script cache, but does not execute the script immediately. After the script has been added to the cache, the script can be called with the SHA1 checksum of the script via the EVALSHA command.
	ScriptLoad(script string) (string, error)

	// scripting end

	// zset start

	// ZAdd 将所有指定成员添加到键为key有序集合（sorted set）里面。
	// Adds all specified members to the keyed sorted set (sorted set).
	ZAdd(key string, members ...*Z) (int64, error)

	// ZCard 返回key的有序集元素个数。
	// Returns the number of elements in the sorted set for key.
	ZCard(key string) (int64, error)

	// ZCount 返回有序集key中，score值在min和max之间(默认包括score值等于min或max)的成员。 关于参数min和max的详细使用方法，请参考ZRANGEBYSCORE命令。
	// Returns the members of the sorted set key whose score value is between min and max (by default, the score value is equal to min or max). For details on how to use the parameters min and max, please refer to the ZRANGEBYSCORE command.
	ZCount(key, min, max string) (int64, error)

	// ZIncrBy 为有序集key的成员member的score值加上增量increment。如果key中不存在member，就在key中添加一个member，score是increment（就好像它之前的score是0.0）。如果key不存在，就创建一个只含有指定member成员的有序集合。为有序集key的成员member的score值加上增量increment。如果key中不存在member，就在key中添加一个member，score是increment（就好像它之前的score是0.0）。如果key不存在，就创建一个只含有指定member成员的有序集合。
	// Increment is added to the score value of the member member of the sorted set key. If a member does not exist in the key, add a member to the key with a score of increment (as if the previous score was 0.0). If the key does not exist, create an ordered set containing only the specified member members. Increment is added to the score value of the member member of the sorted set key. If a member does not exist in the key, add a member to the key with a score of increment (as if the previous score was 0.0). If the key does not exist, create an ordered set containing only the specified member members.
	ZIncrBy(key string, increment float64, member string) (float64, error)

	// ZInterStore 计算给定的numkeys个有序集合的交集，并且把结果放到destination中。 在给定要计算的key和其它参数之前，必须先给定key个数(numberkeys)。
	// Computes the intersection of the given numkeys sorted sets and places the result in destination. Before the key and other parameters to be calculated are given, the number of keys (numberkeys) must be given.
	ZInterStore(destination string, store *ZStore) (int64, error)

	// ZLexCount 计算有序集合中指定成员之间的成员数量。
	// Counts the number of members between the specified members in a sorted set.
	ZLexCount(key, min, max string) (int64, error)

	// ZPopMax 删除并返回有序集合key中的最多count个具有最高得分的成员。
	// Removes and returns at most count members with the highest score in the sorted set key.
	ZPopMax(key string, count ...int64) ([]Z, error)

	// ZPopMin 删除并返回有序集合key中的最多count个具有最低得分的成员。
	// Removes and returns at most count members with the lowest score in the sorted set key.
	ZPopMin(key string, count ...int64) ([]Z, error)

	// ZRange 返回存储在有序集合key中的指定范围的元素。 返回的元素可以认为是按得分从最低到最高排列。 如果得分相同，将按字典排序。
	// Returns the specified range of elements stored in the sorted set key. The returned elements can be thought of as ordered from lowest to highest score. If the scores are the same, they will be sorted lexicographically.
	ZRange(key string, start, stop int64) ([]string, error)

	// ZRangeByLex 返回指定成员区间内的成员，按成员字典正序排序, 分数必须相同。 在某些业务场景中,需要对一个字符串数组按名称的字典顺序进行排序时,可以使用Redis中SortSet这种数据结构来处理。
	// Returns the members in the specified member range, sorted in the positive lexicographic order of the members, and the scores must be the same. In some business scenarios, when a string array needs to be sorted according to the lexicographical order of names, a data structure such as SortSet in Redis can be used for processing.
	ZRangeByLex(key string, opt *ZRangeBy) ([]string, error)

	// ZRevRangeByLex 返回指定成员区间内的成员，按成员字典倒序排序, 分数必须相同。
	// Returns the members in the specified member range, sorted in reverse lexicographic order of the members, and the scores must be the same.
	ZRevRangeByLex(key string, opt *ZRangeBy) ([]string, error)

	// ZRangeByScore 返回有序集合中指定分数区间的成员列表。有序集成员按分数值递增(从小到大)次序排列。
	// Returns a list of members of the ordered collection with the specified score interval. The members of the ordered set are arranged in order of increasing score value (from small to large).
	ZRangeByScore(key string, opt *ZRangeBy) ([]string, error)

	// ZRank 返回有序集中指定成员的排名。其中有序集成员按分数值递增(从小到大)顺序排列。
	// Returns the rank of the specified member in the sorted set. The ordered set members are arranged in order of increasing score value (from small to large).
	ZRank(key, member string) (int64, error)

	// ZRem 移除有序集中的一个或多个成员，不存在的成员将被忽略。
	// Removes one or more members from the sorted set, non-existing members are ignored.
	ZRem(key string, members ...interface{}) (int64, error)

	// ZRemRangeByLex 移除有序集合中给定的字典区间的所有成员。
	// Removes all members of the given dictionary range in the sorted set.
	ZRemRangeByLex(key, min, max string) (int64, error)

	// ZRemRangeByRank 移除有序集中，指定排名(rank)区间内的所有成员。
	// Removes all members from an ordered set, specifying a rank interval.
	ZRemRangeByRank(key string, start, stop int64) (int64, error)

	// ZRemRangeByScore 移除有序集中，指定分数（score）区间内的所有成员。
	// Removes all members from an ordered set with a specified score interval.
	ZRemRangeByScore(key, min, max string) (int64, error)

	// ZRevRange 返回有序集中，指定区间内的成员。其中成员的位置按分数值递减(从大到小)来排列。具有相同分数值的成员按字典序的逆序(reverse lexicographical order)排列。
	// Returns a sorted set with members in the specified range. The positions of the members are arranged in decreasing order of score value (from largest to smallest). Members with the same score value are arranged in reverse lexicographical order.
	ZRevRange(key string, start, stop int64) ([]string, error)

	// ZRevRangeByScore 返回有序集中指定分数区间内的所有的成员。有序集成员按分数值递减(从大到小)的次序排列。具有相同分数值的成员按字典序的逆序(reverse lexicographical order )排列。
	// Returns all members of the sorted set within the specified score interval. The ordered set members are arranged in order of decreasing score value (largest to smallest). Members with the same score value are arranged in reverse lexicographical order.
	ZRevRangeByScore(key string, opt *ZRangeBy) ([]string, error)

	// ZRevRank 返回有序集中成员的排名。其中有序集成员按分数值递减(从大到小)排序。
	// Returns the rank of the members in the sorted set. The members of the ordered set are sorted by decreasing score value (large to small).
	ZRevRank(key, member string) (int64, error)

	// ZScore 返回有序集中，成员的分数值。 如果成员元素不是有序集 key 的成员，或 key 不存在，返回 nil 。
	// Returns the fractional value of the members in the sorted set. Returns nil if the member element is not a member of the sorted set key, or if the key does not exist.
	ZScore(key, member string) (float64, error)

	// ZUnionStore 计算给定的一个或多个有序集的并集，其中给定 key 的数量必须以 numkeys 参数指定，并将该并集(结果集)储存到 destination 。
	// Computes the union of one or more ordered sets given, where the number of given keys must be specified with the numkeys parameter, and stores the union (result set) in destination .
	ZUnionStore(key string, store *ZStore) (int64, error)

	// ZScan 用于迭代有序集合中的元素（包括元素成员和元素分值）
	// Used to iterate over elements in a sorted set (including element members and element scores).
	ZScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error)
}
