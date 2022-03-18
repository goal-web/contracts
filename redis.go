package contracts

import (
	"context"
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
	RedisConnectionCtx
	// Subscribe 订阅一组给定的消息频道
	// subscribe to a set of given channels for messages.
	Subscribe(channels []string, closure RedisSubscribeFunc) error

	// PSubscribe 使用通配符订阅一组给定频道
	// subscribe to a set of given channels with wildcards.
	PSubscribe(channels []string, closure RedisSubscribeFunc) error

	// Command 对 Redis 数据库运行命令
	// Run a command against the Redis database.
	Command(method string, args ...interface{}) (interface{}, error)

	PubSubChannels(pattern string) ([]string, error)

	PubSubNumSub(channels ...string) (map[string]int64, error)

	PubSubNumPat() (int64, error)

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

	BitOpAnd(destKey string, keys ...string) (int64, error)

	BitOpNot(destKey string, key string) (int64, error)

	BitOpOr(destKey string, keys ...string) (int64, error)

	BitOpXor(destKey string, keys ...string) (int64, error)

	GetDel(key string) (string, error)

	GetEx(key string, expiration time.Duration) (string, error)

	GetRange(key string, start, end int64) (string, error)

	GetSet(key string, value interface{}) (string, error)

	ClientGetName() (string, error)

	StrLen(key string) (int64, error)

	// getter end
	// keys start

	Keys(pattern string) ([]string, error)

	Del(keys ...string) (int64, error)

	FlushAll() (string, error)

	FlushDB() (string, error)

	Dump(key string) (string, error)

	Exists(keys ...string) (int64, error)

	Expire(key string, expiration time.Duration) (bool, error)

	ExpireAt(key string, tm time.Time) (bool, error)

	PExpire(key string, expiration time.Duration) (bool, error)

	PExpireAt(key string, tm time.Time) (bool, error)

	Migrate(host, port, key string, db int, timeout time.Duration) (string, error)

	Move(key string, db int) (bool, error)

	Persist(key string) (bool, error)

	PTTL(key string) (time.Duration, error)

	TTL(key string) (time.Duration, error)

	RandomKey() (string, error)

	Rename(key, newKey string) (string, error)

	RenameNX(key, newKey string) (bool, error)

	Type(key string) (string, error)

	Wait(numSlaves int, timeout time.Duration) (int64, error)

	Scan(cursor uint64, match string, count int64) ([]string, uint64, error)

	BitCount(key string, count *BitCount) (int64, error)

	// keys end

	// setter start
	Set(key string, value interface{}, expiration time.Duration) (string, error)

	Append(key, value string) (int64, error)

	MSet(values ...interface{}) (string, error)

	MSetNX(values ...interface{}) (bool, error)

	SetNX(key string, value interface{}, expiration time.Duration) (bool, error)

	SetEX(key string, value interface{}, expiration time.Duration) (string, error)

	SetBit(key string, offset int64, value int) (int64, error)

	BitPos(key string, bit int64, pos ...int64) (int64, error)

	SetRange(key string, offset int64, value string) (int64, error)

	Incr(key string) (int64, error)

	Decr(key string) (int64, error)

	IncrBy(key string, value int64) (int64, error)

	DecrBy(key string, value int64) (int64, error)

	IncrByFloat(key string, value float64) (float64, error)

	// setter end

	// hash start
	HGet(key, field string) (string, error)

	HGetAll(key string) (map[string]string, error)

	HMGet(key string, fields ...string) ([]interface{}, error)

	HKeys(key string) ([]string, error)

	HLen(key string) (int64, error)

	HRandField(key string, count int, withValues bool) ([]string, error)

	HScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error)

	HValues(key string) ([]string, error)

	HSet(key string, values ...interface{}) (int64, error)

	HSetNX(key, field string, value interface{}) (bool, error)

	HMSet(key string, values ...interface{}) (bool, error)

	HDel(key string, fields ...string) (int64, error)

	HExists(key string, field string) (bool, error)

	HIncrBy(key string, field string, value int64) (int64, error)

	HIncrByFloat(key string, field string, value float64) (float64, error)

	// hash end

	// set start
	SAdd(key string, members ...interface{}) (int64, error)

	SCard(key string) (int64, error)

	SDiff(keys ...string) ([]string, error)

	SDiffStore(destination string, keys ...string) (int64, error)

	SInter(keys ...string) ([]string, error)

	SInterStore(destination string, keys ...string) (int64, error)

	SIsMember(key string, member interface{}) (bool, error)

	SMembers(key string) ([]string, error)

	SRem(key string, members ...interface{}) (int64, error)

	SPopN(key string, count int64) ([]string, error)

	SPop(key string) (string, error)

	SRandMemberN(key string, count int64) ([]string, error)

	SMove(source, destination string, member interface{}) (bool, error)

	SRandMember(key string) (string, error)

	SUnion(keys ...string) ([]string, error)

	SUnionStore(destination string, keys ...string) (int64, error)

	// set end

	// geo start

	GeoAdd(key string, geoLocation ...*GeoLocation) (int64, error)

	GeoHash(key string, members ...string) ([]string, error)

	GeoPos(key string, members ...string) ([]*GeoPos, error)

	GeoDist(key string, member1, member2, unit string) (float64, error)

	GeoRadius(key string, longitude, latitude float64, query *GeoRadiusQuery) ([]GeoLocation, error)

	GeoRadiusStore(key string, longitude, latitude float64, query *GeoRadiusQuery) (int64, error)

	GeoRadiusByMember(key, member string, query *GeoRadiusQuery) ([]GeoLocation, error)

	GeoRadiusByMemberStore(key, member string, query *GeoRadiusQuery) (int64, error)

	// geo end

	// lists start

	BLPop(timeout time.Duration, keys ...string) ([]string, error)

	BRPop(timeout time.Duration, keys ...string) ([]string, error)

	BRPopLPush(source, destination string, timeout time.Duration) (string, error)

	LIndex(key string, index int64) (string, error)

	LInsert(key, op string, pivot, value interface{}) (int64, error)

	LLen(key string) (int64, error)

	LPop(key string) (string, error)

	LPush(key string, values ...interface{}) (int64, error)

	LPushX(key string, values ...interface{}) (int64, error)

	LRange(key string, start, stop int64) ([]string, error)

	LRem(key string, count int64, value interface{}) (int64, error)

	LSet(key string, index int64, value interface{}) (string, error)

	LTrim(key string, start, stop int64) (string, error)

	RPop(key string) (string, error)

	RPopCount(key string, count int) ([]string, error)

	RPopLPush(source, destination string) (string, error)

	RPush(key string, values ...interface{}) (int64, error)

	RPushX(key string, values ...interface{}) (int64, error)

	// lists end

	// scripting start
	Eval(script string, keys []string, args ...interface{}) (interface{}, error)

	EvalSha(sha1 string, keys []string, args ...interface{}) (interface{}, error)

	ScriptExists(hashes ...string) ([]bool, error)

	ScriptFlush() (string, error)

	ScriptKill() (string, error)

	ScriptLoad(script string) (string, error)

	// scripting end

	// zset start

	ZAdd(key string, members ...*Z) (int64, error)

	ZCard(key string) (int64, error)

	ZCount(key, min, max string) (int64, error)

	ZIncrBy(key string, increment float64, member string) (float64, error)

	ZInterStore(destination string, store *ZStore) (int64, error)

	ZLexCount(key, min, max string) (int64, error)

	ZPopMax(key string, count ...int64) ([]Z, error)

	ZPopMin(key string, count ...int64) ([]Z, error)

	ZRange(key string, start, stop int64) ([]string, error)

	ZRangeByLex(key string, opt *ZRangeBy) ([]string, error)

	ZRevRangeByLex(key string, opt *ZRangeBy) ([]string, error)

	ZRangeByScore(key string, opt *ZRangeBy) ([]string, error)

	ZRank(key, member string) (int64, error)

	ZRem(key string, members ...interface{}) (int64, error)

	ZRemRangeByLex(key, min, max string) (int64, error)

	ZRemRangeByRank(key string, start, stop int64) (int64, error)

	ZRemRangeByScore(key, min, max string) (int64, error)

	ZRevRange(key string, start, stop int64) ([]string, error)

	ZRevRangeByScore(key string, opt *ZRangeBy) ([]string, error)

	ZRevRank(key, member string) (int64, error)

	ZScore(key, member string) (float64, error)

	ZUnionStore(key string, store *ZStore) (int64, error)

	ZScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error)
}

type RedisConnectionCtx interface {
	// SubscribeCtx 订阅一组给定的消息频道
	// SubscribeCtx to a set of given channels for messages.
	SubscribeCtx(ctx context.Context, channels []string, closure RedisSubscribeFunc) error

	// PSubscribeCtx 使用通配符订阅一组给定频道
	// SubscribeCtx to a set of given channels with wildcards.
	PSubscribeCtx(ctx context.Context, channels []string, closure RedisSubscribeFunc) error

	// CommandCtx 对 Redis 数据库运行命令
	// Run a CommandCtx against the Redis database.
	CommandCtx(ctx context.Context, method string, args ...interface{}) (interface{}, error)

	PubSubChannelsCtx(ctx context.Context, pattern string) ([]string, error)

	PubSubNumSubCtx(ctx context.Context, channels ...string) (map[string]int64, error)

	PubSubNumPatCtx(ctx context.Context) (int64, error)

	PublishCtx(ctx context.Context, channel string, message interface{}) (int64, error)

	// GetCtx 返回给定键的值
	// Returns the value of the given key.
	GetCtx(ctx context.Context, key string) (string, error)

	// MGetCtx 获取所有给定键的值
	// get the values of all the given keys.
	MGetCtx(ctx context.Context, keys ...string) ([]interface{}, error)

	// GetBitCtx 对 key 所储存的字符串值，对获取指定偏移量上的位(bit)
	// For the string value stored in key, get the bit at the specified offset.
	GetBitCtx(ctx context.Context, key string, offset int64) (int64, error)

	BitOpAndCtx(ctx context.Context, destKey string, keys ...string) (int64, error)

	BitOpNotCtx(ctx context.Context, destKey string, key string) (int64, error)

	BitOpOrCtx(ctx context.Context, destKey string, keys ...string) (int64, error)

	BitOpXorCtx(ctx context.Context, destKey string, keys ...string) (int64, error)

	GetDelCtx(ctx context.Context, key string) (string, error)

	GetExCtx(ctx context.Context, key string, expiration time.Duration) (string, error)

	GetRangeCtx(ctx context.Context, key string, start, end int64) (string, error)

	GetSetCtx(ctx context.Context, key string, value interface{}) (string, error)

	ClientGetNameCtx(ctx context.Context) (string, error)

	StrLenCtx(ctx context.Context, key string) (int64, error)

	// getter end
	// keys start

	KeysCtx(ctx context.Context, pattern string) ([]string, error)

	DelCtx(ctx context.Context, keys ...string) (int64, error)

	FlushAllCtx(ctx context.Context) (string, error)

	FlushDBCtx(ctx context.Context) (string, error)

	DumpCtx(ctx context.Context, key string) (string, error)

	ExistsCtx(ctx context.Context, keys ...string) (int64, error)

	ExpireCtx(ctx context.Context, key string, expiration time.Duration) (bool, error)

	ExpireAtCtx(ctx context.Context, key string, tm time.Time) (bool, error)

	PExpireCtx(ctx context.Context, key string, expiration time.Duration) (bool, error)

	PExpireAtCtx(ctx context.Context, key string, tm time.Time) (bool, error)

	MigrateCtx(ctx context.Context, host, port, key string, db int, timeout time.Duration) (string, error)

	MoveCtx(ctx context.Context, key string, db int) (bool, error)

	PersistCtx(ctx context.Context, key string) (bool, error)

	PTTLCtx(ctx context.Context, key string) (time.Duration, error)

	TTLCtx(ctx context.Context, key string) (time.Duration, error)

	RandomKeyCtx(ctx context.Context) (string, error)

	RenameCtx(ctx context.Context, key, newKey string) (string, error)

	RenameNXCtx(ctx context.Context, key, newKey string) (bool, error)

	TypeCtx(ctx context.Context, key string) (string, error)

	WaitCtx(ctx context.Context, numSlaves int, timeout time.Duration) (int64, error)

	ScanCtx(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64, error)

	BitCountCtx(ctx context.Context, key string, count *BitCount) (int64, error)

	// keys end

	// setter start
	SetCtx(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error)

	AppendCtx(ctx context.Context, key, value string) (int64, error)

	MSetCtx(ctx context.Context, values ...interface{}) (string, error)

	MSetNXCtx(ctx context.Context, values ...interface{}) (bool, error)

	SetNXCtx(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)

	SetEXCtx(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error)

	SetBitCtx(ctx context.Context, key string, offset int64, value int) (int64, error)

	BitPosCtx(ctx context.Context, key string, bit int64, pos ...int64) (int64, error)

	SetRangeCtx(ctx context.Context, key string, offset int64, value string) (int64, error)

	IncrCtx(ctx context.Context, key string) (int64, error)

	DecrCtx(ctx context.Context, key string) (int64, error)

	IncrByCtx(ctx context.Context, key string, value int64) (int64, error)

	DecrByCtx(ctx context.Context, key string, value int64) (int64, error)

	IncrByFloatCtx(ctx context.Context, key string, value float64) (float64, error)

	// setter end

	// hash start
	HGetCtx(ctx context.Context, key, field string) (string, error)

	HGetAllCtx(ctx context.Context, key string) (map[string]string, error)

	HMGetCtx(ctx context.Context, key string, fields ...string) ([]interface{}, error)

	HKeysCtx(ctx context.Context, key string) ([]string, error)

	HLenCtx(ctx context.Context, key string) (int64, error)

	HRandFieldCtx(ctx context.Context, key string, count int, withValues bool) ([]string, error)

	HScanCtx(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error)

	HValuesCtx(ctx context.Context, key string) ([]string, error)

	HSetCtx(ctx context.Context, key string, values ...interface{}) (int64, error)

	HSetNXCtx(ctx context.Context, key, field string, value interface{}) (bool, error)

	HMSetCtx(ctx context.Context, key string, values ...interface{}) (bool, error)

	HDelCtx(ctx context.Context, key string, fields ...string) (int64, error)

	HExistsCtx(ctx context.Context, key string, field string) (bool, error)

	HIncrByCtx(ctx context.Context, key string, field string, value int64) (int64, error)

	HIncrByFloatCtx(ctx context.Context, key string, field string, value float64) (float64, error)

	// hash end

	// set start
	SAddCtx(ctx context.Context, key string, members ...interface{}) (int64, error)

	SCardCtx(ctx context.Context, key string) (int64, error)

	SDiffCtx(ctx context.Context, keys ...string) ([]string, error)

	SDiffStoreCtx(ctx context.Context, destination string, keys ...string) (int64, error)

	SInterCtx(ctx context.Context, keys ...string) ([]string, error)

	SInterStoreCtx(ctx context.Context, destination string, keys ...string) (int64, error)

	SIsMemberCtx(ctx context.Context, key string, member interface{}) (bool, error)

	SMembersCtx(ctx context.Context, key string) ([]string, error)

	SRemCtx(ctx context.Context, key string, members ...interface{}) (int64, error)

	SPopNCtx(ctx context.Context, key string, count int64) ([]string, error)

	SPopCtx(ctx context.Context, key string) (string, error)

	SRandMemberNCtx(ctx context.Context, key string, count int64) ([]string, error)

	SMoveCtx(ctx context.Context, source, destination string, member interface{}) (bool, error)

	SRandMemberCtx(ctx context.Context, key string) (string, error)

	SUnionCtx(ctx context.Context, keys ...string) ([]string, error)

	SUnionStoreCtx(ctx context.Context, destination string, keys ...string) (int64, error)

	// set end

	// geo start

	GeoAddCtx(ctx context.Context, key string, geoLocation ...*GeoLocation) (int64, error)

	GeoHashCtx(ctx context.Context, key string, members ...string) ([]string, error)

	GeoPosCtx(ctx context.Context, key string, members ...string) ([]*GeoPos, error)

	GeoDistCtx(ctx context.Context, key string, member1, member2, unit string) (float64, error)

	GeoRadiusCtx(ctx context.Context, key string, longitude, latitude float64, query *GeoRadiusQuery) ([]GeoLocation, error)

	GeoRadiusStoreCtx(ctx context.Context, key string, longitude, latitude float64, query *GeoRadiusQuery) (int64, error)

	GeoRadiusByMemberCtx(ctx context.Context, key, member string, query *GeoRadiusQuery) ([]GeoLocation, error)

	GeoRadiusByMemberStoreCtx(ctx context.Context, key, member string, query *GeoRadiusQuery) (int64, error)

	// geo end

	// lists start

	BLPopCtx(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error)

	BRPopCtx(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error)

	BRPopLPushCtx(ctx context.Context, source, destination string, timeout time.Duration) (string, error)

	LIndexCtx(ctx context.Context, key string, index int64) (string, error)

	LInsertCtx(ctx context.Context, key, op string, pivot, value interface{}) (int64, error)

	LLenCtx(ctx context.Context, key string) (int64, error)

	LPopCtx(ctx context.Context, key string) (string, error)

	LPushCtx(ctx context.Context, key string, values ...interface{}) (int64, error)

	LPushXCtx(ctx context.Context, key string, values ...interface{}) (int64, error)

	LRangeCtx(ctx context.Context, key string, start, stop int64) ([]string, error)

	LRemCtx(ctx context.Context, key string, count int64, value interface{}) (int64, error)

	LSetCtx(ctx context.Context, key string, index int64, value interface{}) (string, error)

	LTrimCtx(ctx context.Context, key string, start, stop int64) (string, error)

	RPopCtx(ctx context.Context, key string) (string, error)

	RPopCountCtx(ctx context.Context, key string, count int) ([]string, error)

	RPopLPushCtx(ctx context.Context, source, destination string) (string, error)

	RPushCtx(ctx context.Context, key string, values ...interface{}) (int64, error)

	RPushXCtx(ctx context.Context, key string, values ...interface{}) (int64, error)

	// lists end

	// scripting start
	EvalCtx(ctx context.Context, script string, keys []string, args ...interface{}) (interface{}, error)

	EvalShaCtx(ctx context.Context, sha1 string, keys []string, args ...interface{}) (interface{}, error)

	ScriptExistsCtx(ctx context.Context, hashes ...string) ([]bool, error)

	ScriptFlushCtx(ctx context.Context) (string, error)

	ScriptKillCtx(ctx context.Context) (string, error)

	ScriptLoadCtx(ctx context.Context, script string) (string, error)

	// scripting end

	// zset start

	ZAddCtx(ctx context.Context, key string, members ...*Z) (int64, error)

	ZCardCtx(ctx context.Context, key string) (int64, error)

	ZCountCtx(ctx context.Context, key, min, max string) (int64, error)

	ZIncrByCtx(ctx context.Context, key string, increment float64, member string) (float64, error)

	ZInterStoreCtx(ctx context.Context, destination string, store *ZStore) (int64, error)

	ZLexCountCtx(ctx context.Context, key, min, max string) (int64, error)

	ZPopMaxCtx(ctx context.Context, key string, count ...int64) ([]Z, error)

	ZPopMinCtx(ctx context.Context, key string, count ...int64) ([]Z, error)

	ZRangeCtx(ctx context.Context, key string, start, stop int64) ([]string, error)

	ZRangeByLexCtx(ctx context.Context, key string, opt *ZRangeBy) ([]string, error)

	ZRevRangeByLexCtx(ctx context.Context, key string, opt *ZRangeBy) ([]string, error)

	ZRangeByScoreCtx(ctx context.Context, key string, opt *ZRangeBy) ([]string, error)

	ZRankCtx(ctx context.Context, key, member string) (int64, error)

	ZRemCtx(ctx context.Context, key string, members ...interface{}) (int64, error)

	ZRemRangeByLexCtx(ctx context.Context, key, min, max string) (int64, error)

	ZRemRangeByRankCtx(ctx context.Context, key string, start, stop int64) (int64, error)

	ZRemRangeByScoreCtx(ctx context.Context, key, min, max string) (int64, error)

	ZRevRangeCtx(ctx context.Context, key string, start, stop int64) ([]string, error)

	ZRevRangeByScoreCtx(ctx context.Context, key string, opt *ZRangeBy) ([]string, error)

	ZRevRankCtx(ctx context.Context, key, member string) (int64, error)

	ZScoreCtx(ctx context.Context, key, member string) (float64, error)

	ZUnionStoreCtx(ctx context.Context, key string, store *ZStore) (int64, error)

	ZScanCtx(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error)
}
