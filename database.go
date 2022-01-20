package contracts

type DBConnector func(config Fields) DBConnection

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type DBFactory interface {
	Connection(key string) DBConnection
	Extend(name string, driver DBConnector)
}

type DBTx interface {
	SqlExecutor
	Commit() error
	Rollback() error
}

type SqlExecutor interface {
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

type Callback func(QueryBuilder) QueryBuilder
type Provider func() QueryBuilder
type WhereFunc func(QueryBuilder)
type whereJoinType string
type unionJoinType string

const (
	Union    unionJoinType = "union"
	UnionAll unionJoinType = "union all"
)

type orderType string

const (
	Desc orderType = "desc"
	Asc  orderType = "asc"
)

type joinType string

const (
	LeftJoin    joinType = "left"
	RightJoin   joinType = "right"
	InnerJoin   joinType = "inner"
	FullOutJoin joinType = "full outer"
	FullJoin    joinType = "full"
)

type insertType string

const (
	Insert        insertType = "insert"
	InsertIgnore  insertType = "insert ignore"
	InsertReplace insertType = "replace"
)

type QueryBuilder interface {
	Select(column string, columns ...string) QueryBuilder
	AddSelect(columns ...string) QueryBuilder
	SelectSub(provider Provider, as string) QueryBuilder
	AddSelectSub(provider Provider, as string) QueryBuilder
	Count(columns ...string) QueryBuilder
	Avg(column string, as ...string) QueryBuilder
	Sum(column string, as ...string) QueryBuilder
	Max(column string, as ...string) QueryBuilder
	Min(column string, as ...string) QueryBuilder

	Distinct() QueryBuilder

	From(table string, as ...string) QueryBuilder
	FromMany(tables ...string) QueryBuilder
	FromSub(provider Provider, as string) QueryBuilder

	Join(table string, first, condition, second string, joins ...joinType) QueryBuilder
	JoinSub(provider Provider, as, first, condition, second string, joins ...joinType) QueryBuilder
	FullJoin(table string, first, condition, second string) QueryBuilder
	FullOutJoin(table string, first, condition, second string) QueryBuilder
	LeftJoin(table string, first, condition, second string) QueryBuilder
	RightJoin(table string, first, condition, second string) QueryBuilder

	Where(column string, args ...interface{}) QueryBuilder
	OrWhere(column string, args ...interface{}) QueryBuilder
	WhereFunc(callback WhereFunc, whereType ...whereJoinType) QueryBuilder
	OrWhereFunc(callback WhereFunc) QueryBuilder

	WhereIn(column string, args interface{}) QueryBuilder
	OrWhereIn(column string, args interface{}) QueryBuilder
	WhereNotIn(column string, args interface{}) QueryBuilder
	OrWhereNotIn(column string, args interface{}) QueryBuilder

	WhereBetween(column string, args interface{}, whereType ...whereJoinType) QueryBuilder
	OrWhereBetween(column string, args interface{}) QueryBuilder
	WhereNotBetween(column string, args interface{}, whereType ...whereJoinType) QueryBuilder
	OrWhereNotBetween(column string, args interface{}) QueryBuilder

	WhereIsNull(column string, whereType ...string) QueryBuilder
	OrWhereIsNull(column string) QueryBuilder
	OrWhereNotNull(column string) QueryBuilder
	WhereNotNull(column string, whereType ...string) QueryBuilder

	WhereExists(provider Provider, where ...whereJoinType) QueryBuilder
	OrWhereExists(provider Provider) QueryBuilder
	WhereNotExists(provider Provider, where ...whereJoinType) QueryBuilder
	OrWhereNotExists(provider Provider) QueryBuilder

	Union(builder QueryBuilder, unionType ...unionJoinType) QueryBuilder
	UnionAll(builder QueryBuilder) QueryBuilder
	UnionByProvider(builder Provider, unionType ...unionJoinType) QueryBuilder
	UnionAllByProvider(builder Provider) QueryBuilder

	GroupBy(columns ...string) QueryBuilder
	Having(column string, args ...interface{}) QueryBuilder
	OrHaving(column string, args ...interface{}) QueryBuilder

	OrderBy(column string, columnOrderType ...orderType)
	OrderByDesc(column string)

	When(condition bool, callback Callback, elseCallback ...Callback)

	ToSql() string
	GetBindings() (results []interface{})

	SelectSql() (string, []interface{})
	CreateSql(value Fields, insertType2 ...insertType) (sql string, bindings []interface{})
	InsertSql(values []Fields, insertType2 ...insertType) (sql string, bindings []interface{})
	InsertIgnoreSql(values []Fields) (sql string, bindings []interface{})
	InsertReplaceSql(values []Fields) (sql string, bindings []interface{})
	DeleteSql() (sql string, bindings []interface{})
	UpdateSql(value Fields) (sql string, bindings []interface{})
}
