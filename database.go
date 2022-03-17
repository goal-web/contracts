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

// QueryCallback 查询回调，用于构建子查询
// query callback，for building subqueries.
type QueryCallback func(QueryBuilder) QueryBuilder

// QueryProvider 查询提供者
// query provider.
type QueryProvider func() QueryBuilder

// QueryFunc 用于构造 子 where 表达式
// Used to construct sub-where expressions.
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
	// Select 设置要选择的列
	// Set the columns to be selected.
	Select(columns ...string) QueryBuilder

	// AddSelect 追加要选择的列
	// Append the columns to be selected.
	AddSelect(columns ...string) QueryBuilder

	// SelectSub 向查询中添加子选择表达式
	// Add a subselect expression to the query.
	SelectSub(provider QueryProvider, as string) QueryBuilder

	// AddSelectSub 向查询中追加子选择表达式
	// Append a subselect expression to the query.
	AddSelectSub(provider QueryProvider, as string) QueryBuilder

	// WithCount 添加子选择查询以计算关系
	// Add subselect queries to count the relations.
	WithCount(columns ...string) QueryBuilder

	// WithAvg 添加子选择查询以包括关系列的平均值
	// Add subselect queries to include the average of the relation's column.
	WithAvg(column string, as ...string) QueryBuilder

	// WithSum 添加子选择查询以包括关系列的总和
	// Add subselect queries to include the sum of the relation's column.
	WithSum(column string, as ...string) QueryBuilder

	// WithMax 添加子选择查询以包含关系列的最大值
	// Add subselect queries to include the max of the relation's column.
	WithMax(column string, as ...string) QueryBuilder

	// WithMin 添加子选择查询以包括关系列的最小值
	// Add subselect queries to include the min of the relation's column.
	WithMin(column string, as ...string) QueryBuilder

	// Count 检索查询的“count”结果
	// Retrieve the "count" result of the query.
	Count(columns ...string) int64

	// Avg 检索给定列的平均值
	// Retrieve the average of the values of a given column.
	Avg(column string, as ...string) int64

	// Sum 检索给定列的值的总和
	// Retrieve the sum of the values of a given column.
	Sum(column string, as ...string) int64

	// Max 检索给定列的最大值
	// Retrieve the maximum value of a given column.
	Max(column string, as ...string) int64

	// Min 检索给定列的最小值
	// Retrieve the minimum value of a given column.
	Min(column string, as ...string) int64

	// Distinct 强制查询只返回不同的结果
	// Force the query to only return distinct results.
	Distinct() QueryBuilder

	// From 设置查询所针对的表
	// Set the table which the query is targeting.
	From(table string, as ...string) QueryBuilder

	// FromMany 设置许多查询所针对的表
	// Set the table that many queries are against.
	FromMany(tables ...string) QueryBuilder

	// FromSub 从子查询中“从”获取
	// Makes "from" fetch from a subquery.
	FromSub(provider QueryProvider, as string) QueryBuilder

	// Join 向查询中添加连接子句
	// Add a join clause to the query.
	Join(table string, first, condition, second string, joins ...JoinType) QueryBuilder

	// JoinSub 向查询添加子查询连接子句
	// Add a subquery join clause to the query.
	JoinSub(provider QueryProvider, as, first, condition, second string, joins ...JoinType) QueryBuilder

	// FullJoin 向查询添加全连接，两表关联查询它们的所有记录。
	// Add a full join to the query, associate the two tables, and query all their records.
	FullJoin(table string, first, condition, second string) QueryBuilder

	// FullOutJoin 向查询添加完整外部连接
	// Add a full outer join to the query
	FullOutJoin(table string, first, condition, second string) QueryBuilder

	// LeftJoin 向查询添加左连接
	// Add a left join to the query.
	LeftJoin(table string, first, condition, second string) QueryBuilder

	// RightJoin 向查询添加右连接
	// Add a right join to the query.
	RightJoin(table string, first, condition, second string) QueryBuilder

	// Where 向查询添加基本 where 子句
	// Add a basic where clause to the query.
	Where(column string, args ...interface{}) QueryBuilder

	// WhereFields 将 where 子句数组添加到查询中
	// Add an array of where clauses to the query.
	WhereFields(fields Fields) QueryBuilder

	// OrWhere 在查询中添加“或 where”子句
	// Add an "or where" clause to the query.
	OrWhere(column string, args ...interface{}) QueryBuilder

	//WhereFunc 向查询中添加嵌套的 where 语句
	// Add a nested where statement to the query.
	WhereFunc(callback QueryFunc, whereType ...WhereJoinType) QueryBuilder

	// OrWhereFunc 向查询中添加嵌套的 or where 语句
	// Add a nested "or where" statement to the query
	OrWhereFunc(callback QueryFunc) QueryBuilder

	// WhereIn 在查询中添加“where in”子句
	// Add a "where in" clause to the query.
	WhereIn(column string, args interface{}, whereType ...WhereJoinType) QueryBuilder

	// OrWhereIn 在查询中添加“or where in”子句
	// Add an "or where in" clause to the query.
	OrWhereIn(column string, args interface{}) QueryBuilder

	// WhereNotIn 在查询中添加“where not in”子句
	// Add a "where not in" clause to the query.
	WhereNotIn(column string, args interface{}, whereType ...WhereJoinType) QueryBuilder

	// OrWhereNotIn 在查询中添加“or where not in”子句
	// Add an "or where not in" clause to the query.
	OrWhereNotIn(column string, args interface{}) QueryBuilder

	// WhereBetween 在查询中添加 where between 语句
	// Add a where between statement to the query.
	WhereBetween(column string, args interface{}, whereType ...WhereJoinType) QueryBuilder

	// OrWhereBetween 在查询中添加 or where between 语句
	// Add an or where between statement to the query.
	OrWhereBetween(column string, args interface{}) QueryBuilder

	// WhereNotBetween 在查询中添加 where not between 语句
	// Add a where not between statement to the query.
	WhereNotBetween(column string, args interface{}, whereType ...WhereJoinType) QueryBuilder

	// OrWhereNotBetween 在查询中添加 or where not between 语句
	// Add an or where not between statement to the query.
	OrWhereNotBetween(column string, args interface{}) QueryBuilder

	// WhereIsNull 在查询中添加“where null”子句
	// Add a "where null" clause to the query.
	WhereIsNull(column string, whereType ...WhereJoinType) QueryBuilder

	// OrWhereIsNull 在查询中添加“or where null”子句
	// Add an "or where null" clause to the query.
	OrWhereIsNull(column string) QueryBuilder

	// OrWhereNotNull 在查询中添加“or where not null”子句
	// Add an "or where not null" clause to the query.
	OrWhereNotNull(column string) QueryBuilder

	// WhereNotNull 在查询中添加“where not null”子句
	// Add a "where not null" clause to the query.
	WhereNotNull(column string, whereType ...WhereJoinType) QueryBuilder

	// WhereExists 在查询中添加一个存在子句
	// Add an exists clause to the query.
	WhereExists(provider QueryProvider, where ...WhereJoinType) QueryBuilder

	// OrWhereExists 向查询中添加或存在子句
	// Add an or exists clause to the query.
	OrWhereExists(provider QueryProvider) QueryBuilder

	// WhereNotExists 在查询中添加 where not exists 子句
	// Add a where not exists clause to the query.
	WhereNotExists(provider QueryProvider, where ...WhereJoinType) QueryBuilder

	// OrWhereNotExists 在查询中添加 where not exists 子句
	// Add a where not exists clause to the query.
	OrWhereNotExists(provider QueryProvider) QueryBuilder

	// Union 在查询中添加联合语句
	// Add a union statement to the query.
	Union(builder QueryBuilder, unionType ...UnionJoinType) QueryBuilder

	// UnionAll 在查询中添加 union all 语句
	// Add a union all statement to the query.
	UnionAll(builder QueryBuilder) QueryBuilder

	// UnionByProvider 在查询中添加联合语句，并order by
	// Add a union statement to the query and order by.
	UnionByProvider(builder QueryProvider, unionType ...UnionJoinType) QueryBuilder

	// UnionAllByProvider 在查询中添加 union all 语句，并order by
	// Add a union all statement to the query and order by.
	UnionAllByProvider(builder QueryProvider) QueryBuilder

	// GroupBy 在查询中添加“group by”子句
	// Add a "group by" clause to the query.
	GroupBy(columns ...string) QueryBuilder

	// Having 在查询中添加“有”子句
	// Add a "having" clause to the query.
	Having(column string, args ...interface{}) QueryBuilder

	// OrHaving 在查询中添加“或有”子句
	// Add an "or having" clause to the query.
	OrHaving(column string, args ...interface{}) QueryBuilder

	// OrderBy 在查询中添加“order by”子句
	// Add an "order by" clause to the query.
	OrderBy(column string, columnOrderType ...OrderType) QueryBuilder

	// OrderByDesc 向查询中添加降序“order by”子句
	// Add a descending "order by" clause to the query.
	OrderByDesc(column string) QueryBuilder

	// InRandomOrder 将查询的结果按随机顺序排列
	// Put the query's results in random order.
	InRandomOrder() QueryBuilder

	// InRandOrder 将查询的结果按随机顺序排列
	// Put the query's results in random order.
	InRandOrder() QueryBuilder

	// When 如果给定的“值”为真，则应用回调的查询更改
	// Apply the callback's query changes if the given "value" is true.
	When(condition bool, callback QueryCallback, elseCallback ...QueryCallback) QueryBuilder

	// ToSql 获取查询的 SQL 表示
	// get the SQL representation of the query.
	ToSql() string

	// GetBindings 获取扁平数组中的当前查询值绑定
	// get the current query value bindings in a flattened array.
	GetBindings() (results []interface{})

	// Offset 设置查询的“Offset”值
	// Set the "offset" value of the query.
	Offset(offset int64) QueryBuilder

	// Skip 设置查询“Skip”值的别名
	// Alias to set the "offset" value of the query.
	Skip(offset int64) QueryBuilder

	// Limit  设置查询的“limit”值
	// Set the "limit" value of the query.
	Limit(num int64) QueryBuilder

	// Take 设置查询“limit”值的别名
	// Alias to set the "limit" value of the query.
	Take(num int64) QueryBuilder

	// WithPagination 设置给定页面的”limit”值和”offset”值
	// Set the limit and offset for a given page.
	WithPagination(perPage int64, current ...int64) QueryBuilder

	// Chunk 将查询结果分块
	// chunk the results of the query.
	Chunk(size int, handler func(collection Collection, page int) error) error

	// ChunkById 通过比较 ID 对查询结果进行分块
	// chunk the results of a query by comparing IDs.
	ChunkById(size int, handler func(collection Collection, page int) error) error

	// SelectSql 获取此 query builder 的当前规范形成的完整 SQL 字符串。
	// Gets the complete SQL string formed by the current specifications of this query builder.
	SelectSql() (string, []interface{})

	// SelectForUpdateSql 将此实例转换为 SQL 中的 UPDATE 字符串
	// Converts this instance into an UPDATE string in SQL.
	SelectForUpdateSql() (string, []interface{})
	CreateSql(value Fields, insertType2 ...InsertType) (sql string, bindings []interface{})
	InsertSql(values []Fields, insertType2 ...InsertType) (sql string, bindings []interface{})
	InsertIgnoreSql(values []Fields) (sql string, bindings []interface{})
	InsertReplaceSql(values []Fields) (sql string, bindings []interface{})
	DeleteSql() (sql string, bindings []interface{})
	UpdateSql(value Fields) (sql string, bindings []interface{})

	// SetExecutor 设置执行者
	// set executor.
	SetExecutor(executor SqlExecutor) QueryBuilder

	// Insert 向数据库中插入新记录
	// insert new records into the database.
	Insert(values ...Fields) bool

	// InsertGetId 插入一条新记录并获取主键的值
	// insert a new record and get the value of the primary key.
	InsertGetId(values ...Fields) int64

	// InsertOrIgnore 将新记录插入数据库，同时忽略错误
	// insert new records into the database while ignoring errors.
	InsertOrIgnore(values ...Fields) int64

	// InsertOrReplace 将新记录插入数据库，同时如果存在，则先删除此行数据，然后插入新的数据
	// Insert a new record into the database, and if it exists, delete this row of data first, and then insert new data.
	InsertOrReplace(values ...Fields) int64

	// Create 保存新模型并返回实例
	// Save a new model and return the instance.
	Create(fields Fields) interface{}

	// FirstOrCreate 获取与属性匹配的第一条记录或创建它
	// get the first record matching the attributes or create it.
	FirstOrCreate(values ...Fields) interface{}

	// Update 更新数据库中的记录
	// update records in the database.
	Update(fields Fields) int64

	// UpdateOrInsert 插入或更新与属性匹配的记录，并用值填充它
	// insert or update a record matching the attributes, and fill it with values.
	UpdateOrInsert(attributes Fields, values ...Fields) bool

	// UpdateOrCreate 创建或更新与属性匹配的记录，并用值填充它
	// create or update a record matching the attributes, and fill it with values.
	UpdateOrCreate(attributes Fields, values ...Fields) interface{}

	// Get 将查询作为“选择”语句执行
	// Execute the query as a "select" statement.
	Get() Collection

	// SelectForUpdate 锁定表中选定的行以进行更新
	// Lock the selected rows in the table for updating.
	SelectForUpdate() Collection

	// Find 按 ID 对单个记录执行查询
	// Execute a query for a single record by ID.
	Find(key interface{}) interface{}

	// First 执行查询并获得第一个结果
	// Execute the query and get the first result.
	First() interface{}

	// FirstOr 执行查询并获得第一个结果或调用回调
	// Execute the query and get the first result or call a callback.
	FirstOr(provider InstanceProvider) interface{}

	// FirstOrFail 执行查询并获得第一个结果或抛出异常
	// Execute the query and get the first result or throw an exception.
	FirstOrFail() interface{}

	// FirstWhere 向查询添加基本 where 子句，并返回第一个结果
	// Add a basic where clause to the query, and return the first result.
	FirstWhere(column string, args ...interface{}) interface{}

	// Delete 从数据库中删除记录
	// delete records from the database.
	Delete() int64

	// Paginate 对给定的查询进行分页。
	// paginate the given query.
	Paginate(perPage int64, current ...int64) (Collection, int64)

	// SimplePaginate 将给定的查询分页成一个简单的分页器
	// paginate the given query into a simple paginator.
	SimplePaginate(perPage int64, current ...int64) Collection

	// Bind 注册查询构造器
	// binding Query Builder.
	Bind(QueryBuilder) QueryBuilder
}

type Model interface {

	// GetClass 获取模型的类
	// Get the class of the model.
	GetClass() Class

	// GetTable 获取与模型关联的表
	// Get the table associated with the model.
	GetTable() string

	// GetConnection 获取模型的数据库连接
	// Get the database connection for the model.
	GetConnection() string

	// GetPrimaryKey 获取模型的主键
	// Get the primary key for the model.
	GetPrimaryKey() string
}

// MigrateHandler 数据库迁移处理程序
// Database Migration Handler.
type MigrateHandler func(db DBConnection) error

type Migrate struct {
	Name       string
	Connection string
	CreatedAt  time.Time
	Up         MigrateHandler
	Down       MigrateHandler
}

type Migrations []Migrate
