package contracts

type DBConnector func(config Fields) DBConnection

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type DBFactory interface {
	Connection(key ...string) DBConnection
	Extend(name string, driver DBConnector)
}

type DBTx interface {
	SqlExecutor
	Commit() error
	Rollback() error
}

type SqlExecutor interface {
	Query(query string, args ...interface{}) ([]Fields, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (Result, error)
}

type DBConnection interface {
	SqlExecutor
	Begin() (DBTx, error)
	Transaction(func(executor SqlExecutor) error) error
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
	Select(column string, columns ...string) QueryBuilder
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

	When(condition bool, callback QueryCallback, elseCallback ...QueryCallback) QueryBuilder

	ToSql() string
	GetBindings() (results []interface{})

	Offset(offset int64) QueryBuilder
	Skip(offset int64) QueryBuilder
	Limit(num int64) QueryBuilder
	Take(num int64) QueryBuilder
	WithPagination(perPage int64, current ...int64) QueryBuilder

	SelectSql() (string, []interface{})
	CreateSql(value Fields, insertType2 ...InsertType) (sql string, bindings []interface{})
	InsertSql(values []Fields, insertType2 ...InsertType) (sql string, bindings []interface{})
	InsertIgnoreSql(values []Fields) (sql string, bindings []interface{})
	InsertReplaceSql(values []Fields) (sql string, bindings []interface{})
	DeleteSql() (sql string, bindings []interface{})
	UpdateSql(value Fields) (sql string, bindings []interface{})

	// SetTX 预留给实现端添加事物
	SetTX(tx interface{}) QueryBuilder

	Insert(values ...Fields) bool
	InsertGetId(values ...Fields) int64
	InsertOrIgnore(values ...Fields) int64
	InsertOrReplace(values ...Fields) int64

	Create(fields Fields) interface{}
	FirstOrCreate(values ...Fields) interface{}

	Update(fields Fields) int64
	UpdateOrInsert(attributes Fields, values ...Fields) bool
	UpdateOrCreate(attributes Fields, values ...Fields) interface{}

	Get() interface{}
	Find(key interface{}) interface{}
	First() interface{}
	FirstOr(provider InstanceProvider) interface{}
	FirstOrFail() interface{}
	FirstWhere(column string, args ...interface{}) interface{}

	Delete() int64

	Paginate(perPage int64, current ...int64) (interface{}, int64)
	SimplePaginate(perPage int64, current ...int64) interface{}
}
