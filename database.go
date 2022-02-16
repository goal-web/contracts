package contracts


import "time"

// DBConnector 获取数据库连接实例
// Get a database connection instance.
type DBConnector func(config Fields, dispatcher EventDispatcher) DBConnection

type Result interface {
	// LastInsertId 获取最后一个插入 ID
	// Get the last insert ID.
	LastInsertId() (int64, error)

	// RowsAffected 获取受影响的行数
	// Get the number of affected rows.
	RowsAffected() (int64, error)
}

type DBFactory interface {
	// Connection 获取指定的数据库连接实例
	// Get the specified database connection instance.
	Connection(key ...string) DBConnection

	// Extend 注册扩展数据库连接实例
	// Register an extension database connection resolver.
	Extend(name string, driver DBConnector)
}

type DBTx interface {
	SqlExecutor
	// Commit 提交活动的数据库事务
	// commit the active database transaction.
	Commit() error

	// Rollback 回滚活动的数据库事务
	// rollback the active database transaction.
	Rollback() error
}

type SqlExecutor interface {
	// Query 对连接执行新查询
	// Execute a new query against the connection.
	Query(query string, args ...interface{}) (Collection, error)

	// Get 将查询作为“select”语句执行。
	// Execute the query as a "select" statement.
	Get(dest interface{}, query string, args ...interface{}) error

	// Select 对数据库运行 select 语句
	// Run a select statement against the database.
	Select(dest interface{}, query string, args ...interface{}) error

	// Exec 执行一条 SQL 语句
	// Execute an SQL statement.
	Exec(query string, args ...interface{}) (Result, error)
}

type DBConnection interface {
	SqlExecutor

	// Begin 开始一个新的数据库事务。
	// Start a new database transaction.
	Begin() (DBTx, error)

	// Transaction 在事务中执行闭包
	// Execute a Closure within a transaction.
	Transaction(func(executor SqlExecutor) error) error

	// DriverName 获取驱动程序名称
	// Get the driver name.
	DriverName() string
}

type QueryCallback func(QueryBuilder) QueryBuilder
type QueryProvider func() QueryBuilder
type QueryFunc func(QueryBuilder)
type WhereJoinType string
type UnionJoinType string

const (
	Union    UnionJoinType = "union"
	UnionAll UnionJoinType = "union all"
)

type OrderType string

const (
	Desc OrderType = "desc"
	Asc  OrderType = "asc"
)

type JoinType string

const (
	LeftJoin    JoinType = "left"
	RightJoin   JoinType = "right"
	InnerJoin   JoinType = "inner"
	FullOutJoin JoinType = "full outer"
	FullJoin    JoinType = "full"
)

type InsertType string

const (
	Insert        InsertType = "insert"
	InsertIgnore  InsertType = "insert ignore"
	InsertReplace InsertType = "replace"
)

const (
	And WhereJoinType = "and"
	Or  WhereJoinType = "or"
)

type QueryBuilder interface {
	Select(columns ...string) QueryBuilder
	AddSelect(columns ...string) QueryBuilder
	SelectSub(provider QueryProvider, as string) QueryBuilder
	AddSelectSub(provider QueryProvider, as string) QueryBuilder

	WithCount(columns ...string) QueryBuilder
	WithAvg(column string, as ...string) QueryBuilder
	WithSum(column string, as ...string) QueryBuilder
	WithMax(column string, as ...string) QueryBuilder
	WithMin(column string, as ...string) QueryBuilder

	Count(columns ...string) int64
	Avg(column string, as ...string) int64
	Sum(column string, as ...string) int64
	Max(column string, as ...string) int64
	Min(column string, as ...string) int64

	Distinct() QueryBuilder

	From(table string, as ...string) QueryBuilder
	FromMany(tables ...string) QueryBuilder
	FromSub(provider QueryProvider, as string) QueryBuilder

	Join(table string, first, condition, second string, joins ...JoinType) QueryBuilder
	JoinSub(provider QueryProvider, as, first, condition, second string, joins ...JoinType) QueryBuilder
	FullJoin(table string, first, condition, second string) QueryBuilder
	FullOutJoin(table string, first, condition, second string) QueryBuilder
	LeftJoin(table string, first, condition, second string) QueryBuilder
	RightJoin(table string, first, condition, second string) QueryBuilder

	Where(column string, args ...interface{}) QueryBuilder
	WhereFields(fields Fields) QueryBuilder
	OrWhere(column string, args ...interface{}) QueryBuilder
	WhereFunc(callback QueryFunc, whereType ...WhereJoinType) QueryBuilder
	OrWhereFunc(callback QueryFunc) QueryBuilder

	WhereIn(column string, args interface{}, whereType ...WhereJoinType) QueryBuilder
	OrWhereIn(column string, args interface{}) QueryBuilder
	WhereNotIn(column string, args interface{}, whereType ...WhereJoinType) QueryBuilder
	OrWhereNotIn(column string, args interface{}) QueryBuilder

	WhereBetween(column string, args interface{}, whereType ...WhereJoinType) QueryBuilder
	OrWhereBetween(column string, args interface{}) QueryBuilder
	WhereNotBetween(column string, args interface{}, whereType ...WhereJoinType) QueryBuilder
	OrWhereNotBetween(column string, args interface{}) QueryBuilder

	WhereIsNull(column string, whereType ...WhereJoinType) QueryBuilder
	OrWhereIsNull(column string) QueryBuilder
	OrWhereNotNull(column string) QueryBuilder
	WhereNotNull(column string, whereType ...WhereJoinType) QueryBuilder

	WhereExists(provider QueryProvider, where ...WhereJoinType) QueryBuilder
	OrWhereExists(provider QueryProvider) QueryBuilder
	WhereNotExists(provider QueryProvider, where ...WhereJoinType) QueryBuilder
	OrWhereNotExists(provider QueryProvider) QueryBuilder

	Union(builder QueryBuilder, unionType ...UnionJoinType) QueryBuilder
	UnionAll(builder QueryBuilder) QueryBuilder
	UnionByProvider(builder QueryProvider, unionType ...UnionJoinType) QueryBuilder
	UnionAllByProvider(builder QueryProvider) QueryBuilder

	GroupBy(columns ...string) QueryBuilder
	Having(column string, args ...interface{}) QueryBuilder
	OrHaving(column string, args ...interface{}) QueryBuilder

	OrderBy(column string, columnOrderType ...OrderType) QueryBuilder
	OrderByDesc(column string) QueryBuilder
	InRandomOrder() QueryBuilder

	When(condition bool, callback QueryCallback, elseCallback ...QueryCallback) QueryBuilder

	ToSql() string
	GetBindings() (results []interface{})

	Offset(offset int64) QueryBuilder
	Skip(offset int64) QueryBuilder
	Limit(num int64) QueryBuilder
	Take(num int64) QueryBuilder
	WithPagination(perPage int64, current ...int64) QueryBuilder

	Chunk(size int, handler func(collection Collection, page int) error) error
	ChunkById(size int, handler func(collection Collection, page int) error) error

	SelectSql() (string, []interface{})
	SelectForUpdateSql() (string, []interface{})
	CreateSql(value Fields, insertType2 ...InsertType) (sql string, bindings []interface{})
	InsertSql(values []Fields, insertType2 ...InsertType) (sql string, bindings []interface{})
	InsertIgnoreSql(values []Fields) (sql string, bindings []interface{})
	InsertReplaceSql(values []Fields) (sql string, bindings []interface{})
	DeleteSql() (sql string, bindings []interface{})
	UpdateSql(value Fields) (sql string, bindings []interface{})

	SetExecutor(executor SqlExecutor) QueryBuilder

	Insert(values ...Fields) bool
	InsertGetId(values ...Fields) int64
	InsertOrIgnore(values ...Fields) int64
	InsertOrReplace(values ...Fields) int64

	Create(fields Fields) interface{}
	FirstOrCreate(values ...Fields) interface{}

	Update(fields Fields) int64
	UpdateOrInsert(attributes Fields, values ...Fields) bool
	UpdateOrCreate(attributes Fields, values ...Fields) interface{}

	Get() Collection
	SelectForUpdate() Collection
	Find(key interface{}) interface{}
	First() interface{}
	FirstOr(provider InstanceProvider) interface{}
	FirstOrFail() interface{}
	FirstWhere(column string, args ...interface{}) interface{}

	Delete() int64

	Paginate(perPage int64, current ...int64) (Collection, int64)
	SimplePaginate(perPage int64, current ...int64) Collection

	Bind(QueryBuilder) QueryBuilder
}

type Model interface {
	GetClass() Class
	GetTable() string
	GetConnection() string
	GetPrimaryKey() string
}

type MigrateHandler func(db DBConnection) error

type Migrate struct {
	Name       string
	Connection string
	CreatedAt  time.Time
	Up         MigrateHandler
	Down       MigrateHandler
}

type Migrations []Migrate
