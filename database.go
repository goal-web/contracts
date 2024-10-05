package contracts

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
	Commit() Exception

	// Rollback 回滚活动的数据库事务
	// rollback the active database transaction.
	Rollback() Exception
}

type SqlExecutor interface {
	// Query 对连接执行新查询
	// Execute a new query against the connection.
	Query(query string, args ...any) (Collection[Fields], Exception)

	// Get 将查询作为“select”语句执行。
	// Execute the query as a "select" statement.
	Get(dest any, query string, args ...any) Exception

	// Select 对数据库运行 select 语句
	// Run a select statement against the database.
	Select(dest any, query string, args ...any) Exception

	// Exec 执行一条 SQL 语句
	// Execute an SQL statement.
	Exec(query string, args ...any) (Result, Exception)
}

type DBConnection interface {
	SqlExecutor

	// Begin 开始一个新的数据库事务。
	// Start a new database transaction.
	Begin() (DBTx, Exception)

	// Transaction 在事务中执行闭包
	// Execute a Closure within a transaction.
	Transaction(func(executor SqlExecutor) Exception) Exception

	// DriverName 获取驱动程序名称
	// Get the driver name.
	DriverName() string
}

// QueryCallback 查询回调，用于构建子查询
// query callback，for building subqueries.
type QueryCallback[T any] func(QueryBuilder[T]) QueryBuilder[T]

// QueryProvider 查询提供者
// query provider.
type QueryProvider[T any] func() QueryBuilder[T]

// QueryFunc 用于构造 子 where 表达式
// Used to construct sub-where expressions.
type QueryFunc[T any] func(QueryBuilder[T])

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

type QueryExecutor[T any] interface {

	// Count 检索查询的“count”结果
	// Retrieve the "count" result of the query.
	Count(columns ...string) int64

	// Avg 检索给定列的平均值
	// Retrieve the average of the values of a given column.
	Avg(column string) float64

	// Sum 检索给定列的值的总和
	// Retrieve the sum of the values of a given column.
	Sum(column string) float64

	// Max 检索给定列的最大值
	// Retrieve the maximum value of a given column.
	Max(column string) float64

	// Min 检索给定列的最小值
	// Retrieve the minimum value of a given column.
	Min(column string) float64

	// CountE 检索查询的“count”结果
	// Retrieve the "count" result of the query.
	CountE(columns ...string) (int64, Exception)

	// AvgE 检索给定列的平均值
	// Retrieve the average of the values of a given column.
	AvgE(column string) (float64, Exception)

	// SumE 检索给定列的值的总和
	// Retrieve the sum of the values of a given column.
	SumE(column string) (float64, Exception)

	// MaxE 检索给定列的最大值
	// Retrieve the maximum value of a given column.
	MaxE(column string) (float64, Exception)

	// MinE 检索给定列的最小值
	// Retrieve the minimum value of a given column.
	MinE(column string) (float64, Exception)

	// Chunk 将查询结果分块
	// chunk the results of the query.
	Chunk(size int, handler func(collection Collection[*T], page int) Exception) Exception

	// ChunkById 通过比较 ID 对查询结果进行分块
	// chunk the results of a query by comparing IDs.
	ChunkById(size int, handler func(collection Collection[*T], page int) (any, Exception)) Exception

	// ChunkByIdDesc 通过比较 ID 对查询结果进行分块
	// chunk the results of a query by comparing IDs.
	ChunkByIdDesc(size int, handler func(collection Collection[*T], page int) (any, Exception)) Exception

	// Insert 向数据库中插入新记录
	// insert new records into the database.
	Insert(values ...Fields) bool

	// InsertE 向数据库中插入新记录
	// insert new records into the database.
	InsertE(values ...Fields) Exception

	// InsertGetId 插入一条新记录并获取主键的值
	// insert a new record and get the value of the primary key.
	InsertGetId(values ...Fields) int64

	// InsertGetIdE 插入一条新记录并获取主键的值
	// insert a new record and get the value of the primary key.
	InsertGetIdE(values ...Fields) (int64, Exception)

	// InsertOrIgnore 将新记录插入数据库，同时忽略错误
	// insert new records into the database while ignoring Exceptions.
	InsertOrIgnore(values ...Fields) int64

	// InsertOrIgnoreE 将新记录插入数据库，同时忽略错误
	// insert new records into the database while ignoring Exceptions.
	InsertOrIgnoreE(values ...Fields) (int64, Exception)

	// InsertOrReplace 将新记录插入数据库，同时如果存在，则先删除此行数据，然后插入新的数据
	// Insert a new record into the database, and if it exists, delete this row of data first, and then insert new data.
	InsertOrReplace(values ...Fields) int64

	// InsertOrReplaceE 将新记录插入数据库，同时如果存在，则先删除此行数据，然后插入新的数据
	// Insert a new record into the database, and if it exists, delete this row of data first, and then insert new data.
	InsertOrReplaceE(values ...Fields) (int64, Exception)

	// Create 保存新模型并返回实例
	// Save a new model and return the instance.
	Create(fields Fields) *T

	// CreateE 保存新模型并返回实例
	// Save a new model and return the instance.
	CreateE(fields Fields) (*T, Exception)

	// FirstOrCreateE 获取与属性匹配的第一条记录或创建它
	// get the first record matching the attributes or create it.
	FirstOrCreateE(where Fields, values ...Fields) (*T, Exception)

	// FirstOrCreate 获取与属性匹配的第一条记录或创建它
	// get the first record matching the attributes or create it.
	FirstOrCreate(where Fields, values ...Fields) *T

	// Update 更新数据库中的记录
	// update records in the database.
	Update(fields Fields) int64

	// UpdateE 更新数据库中的记录
	// update records in the database.
	UpdateE(fields Fields) (int64, Exception)

	// UpdateOrInsert 插入或更新与属性匹配的记录，并用值填充它
	// insert or update a record matching the attributes, and fill it with values.
	UpdateOrInsert(attributes Fields, values Fields) bool

	// UpdateOrInsertE 插入或更新与属性匹配的记录，并用值填充它
	// insert or update a record matching the attributes, and fill it with values.
	UpdateOrInsertE(attributes Fields, values Fields) Exception

	// UpdateOrCreate 创建或更新与属性匹配的记录，并用值填充它
	// create or update a record matching the attributes, and fill it with values.
	UpdateOrCreate(attributes, values Fields) *T

	// UpdateOrCreateE 创建或更新与属性匹配的记录，并用值填充它
	// create or update a record matching the attributes, and fill it with values.
	UpdateOrCreateE(attributes, values Fields) (*T, Exception)

	// Get 将查询作为“选择”语句执行
	// Execute the query as a "select" statement.
	Get() Collection[*T]

	// GetE 将查询作为“选择”语句执行
	// Execute the query as a "select" statement.
	GetE() (Collection[*T], Exception)

	// SelectForUpdate 锁定表中选定的行以进行更新
	// Lock the selected rows in the table for updating.
	SelectForUpdate() Collection[*T]

	// SelectForUpdateE 锁定表中选定的行以进行更新
	// Lock the selected rows in the table for updating.
	SelectForUpdateE() (Collection[*T], Exception)

	// Find 按 ID 对单个记录执行查询
	// Execute a query for a single record by ID.
	Find(key any) *T

	// FindOrFail 按 ID 对单个记录执行查询
	// Execute a query for a single record by ID.
	FindOrFail(key any) *T

	// First 执行查询并获得第一个结果
	// Execute the query and get the first result.
	First() *T

	// FirstE 执行查询并获得第一个结果
	// Execute the query and get the first result.
	FirstE() (*T, Exception)

	// FirstOr 执行查询并获得第一个结果或调用回调
	// Execute the query and get the first result or call a callback.
	FirstOr(provider InstanceProvider[*T]) *T

	// FirstOrFail 执行查询并获得第一个结果或抛出异常
	// Execute the query and get the first result or throw an exception.
	FirstOrFail() *T

	// FirstWhere 向查询添加基本 where 子句，并返回第一个结果
	// Add a basic where clause to the query, and return the first result.
	FirstWhere(column string, args ...any) *T

	// FirstWhereE 向查询添加基本 where 子句，并返回第一个结果
	// Add a basic where clause to the query, and return the first result.
	FirstWhereE(column string, args ...any) (*T, Exception)

	// Delete 从数据库中删除记录
	// delete records from the database.
	Delete() int64

	// DeleteE 从数据库中删除记录
	// delete records from the database.
	DeleteE() (int64, Exception)

	// Paginate 对给定的查询进行分页。
	// paginate the given query.
	Paginate(perPage int64, current ...int64) (Collection[*T], int64)

	// SimplePaginate 将给定的查询分页成一个简单的分页器
	// paginate the given query into a simple paginator.
	SimplePaginate(perPage int64, current ...int64) Collection[*T]
}

type QueryBuilder[T any] interface {
	QueryExecutor[T]

	With(...RelationType) QueryBuilder[T]

	GetWith() []RelationType

	// Select 设置要选择的列
	// Set the columns to be selected.
	Select(columns ...string) QueryBuilder[T]

	// AddSelect 追加要选择的列
	// Append the columns to be selected.
	AddSelect(columns ...string) QueryBuilder[T]

	// SelectSub 向查询中添加子选择表达式
	// Add a subselect expression to the query.
	SelectSub(provider QueryProvider[T], as string) QueryBuilder[T]

	// AddSelectSub 向查询中追加子选择表达式
	// Append a subselect expression to the query.
	AddSelectSub(provider QueryProvider[T], as string) QueryBuilder[T]

	// WithCount 添加子选择查询以计算关系
	// Add subselect queries to count the relations.
	WithCount(columns ...string) QueryBuilder[T]

	// WithAvg 添加子选择查询以包括关系列的平均值
	// Add subselect queries to include the average of the relation's column.
	WithAvg(column string, as ...string) QueryBuilder[T]

	// WithSum 添加子选择查询以包括关系列的总和
	// Add subselect queries to include the sum of the relation's column.
	WithSum(column string, as ...string) QueryBuilder[T]

	// WithMax 添加子选择查询以包含关系列的最大值
	// Add subselect queries to include the max of the relation's column.
	WithMax(column string, as ...string) QueryBuilder[T]

	// WithMin 添加子选择查询以包括关系列的最小值
	// Add subselect queries to include the min of the relation's column.
	WithMin(column string, as ...string) QueryBuilder[T]

	// Distinct 强制查询只返回不同的结果
	// Force the query to only return distinct results.
	Distinct() QueryBuilder[T]

	// From 设置查询所针对的表
	// Set the table which the query is targeting.
	From(table string, as ...string) QueryBuilder[T]

	// FromMany 设置许多查询所针对的表
	// Set the table that many queries are against.
	FromMany(tables ...string) QueryBuilder[T]

	// FromSub 从子查询中 “从”获取
	// Makes "from" fetch from a subquery.
	FromSub(provider QueryProvider[T], as string) QueryBuilder[T]

	// Join 向查询中添加连接子句
	// Add a join clause to the query.
	Join(table string, first, condition, second string, joins ...JoinType) QueryBuilder[T]

	// JoinSub 向查询添加子查询连接子句
	// Add a subquery join clause to the query.
	JoinSub(provider QueryProvider[T], as, first, condition, second string, joins ...JoinType) QueryBuilder[T]

	// FullJoin 向查询添加全连接，两表关联查询它们的所有记录。
	// Add a full join to the query, associate the two tables, and query all their records.
	FullJoin(table string, first, condition, second string) QueryBuilder[T]

	// FullOutJoin 向查询添加完整外部连接
	// Add a full outer join to the query
	FullOutJoin(table string, first, condition, second string) QueryBuilder[T]

	// LeftJoin 向查询添加左连接
	// Add a left join to the query.
	LeftJoin(table string, first, condition, second string) QueryBuilder[T]

	// RightJoin 向查询添加右连接
	// Add a right join to the query.
	RightJoin(table string, first, condition, second string) QueryBuilder[T]

	// Where 向查询添加基本 where 子句
	// Add a basic where clause to the query.
	Where(column string, args ...any) QueryBuilder[T]

	// WhereFields 将 where 子句数组添加到查询中
	// Add an array of where clauses to the query.
	WhereFields(fields Fields) QueryBuilder[T]

	// OrWhere 在查询中添加“或 where”子句
	// Add an "or where" clause to the query.
	OrWhere(column string, args ...any) QueryBuilder[T]

	//WhereFunc 向查询中添加嵌套的 where 语句
	// Add a nested where statement to the query.
	WhereFunc(callback QueryFunc[T], whereType ...WhereJoinType) QueryBuilder[T]

	// OrWhereFunc 向查询中添加嵌套的 or where 语句
	// Add a nested "or where" statement to the query
	OrWhereFunc(callback QueryFunc[T]) QueryBuilder[T]

	// WhereIn 在查询中添加“where in”子句
	// Add a "where in" clause to the query.
	WhereIn(column string, args any, whereType ...WhereJoinType) QueryBuilder[T]

	// OrWhereIn 在查询中添加“or where in”子句
	// Add an "or where in" clause to the query.
	OrWhereIn(column string, args any) QueryBuilder[T]

	// WhereNotIn 在查询中添加“where not in”子句
	// Add a "where not in" clause to the query.
	WhereNotIn(column string, args any, whereType ...WhereJoinType) QueryBuilder[T]

	// OrWhereNotIn 在查询中添加“or where not in”子句
	// Add an "or where not in" clause to the query.
	OrWhereNotIn(column string, args any) QueryBuilder[T]

	// WhereBetween 在查询中添加 where between 语句
	// Add a where between statement to the query.
	WhereBetween(column string, args any, whereType ...WhereJoinType) QueryBuilder[T]

	// OrWhereBetween 在查询中添加 or where between 语句
	// Add an or where between statement to the query.
	OrWhereBetween(column string, args any) QueryBuilder[T]

	// WhereNotBetween 在查询中添加 where not between 语句
	// Add a where not between statement to the query.
	WhereNotBetween(column string, args any, whereType ...WhereJoinType) QueryBuilder[T]

	// OrWhereNotBetween 在查询中添加 or where not between 语句
	// Add an or where not between statement to the query.
	OrWhereNotBetween(column string, args any) QueryBuilder[T]

	// WhereIsNull 在查询中添加“where null”子句
	// Add a "where null" clause to the query.
	WhereIsNull(column string, whereType ...WhereJoinType) QueryBuilder[T]

	// OrWhereIsNull 在查询中添加“or where null”子句
	// Add an "or where null" clause to the query.
	OrWhereIsNull(column string) QueryBuilder[T]

	// OrWhereNotNull 在查询中添加“or where not null”子句
	// Add an "or where not null" clause to the query.
	OrWhereNotNull(column string) QueryBuilder[T]

	// WhereNotNull 在查询中添加“where not null”子句
	// Add a "where not null" clause to the query.
	WhereNotNull(column string, whereType ...WhereJoinType) QueryBuilder[T]

	// WhereExists 在查询中添加一个存在子句
	// Add an exists clause to the query.
	WhereExists(provider QueryProvider[T], where ...WhereJoinType) QueryBuilder[T]

	// OrWhereExists 向查询中添加或存在子句
	// Add an or exists clause to the query.
	OrWhereExists(provider QueryProvider[T]) QueryBuilder[T]

	// WhereNotExists 在查询中添加 where not exists 子句
	// Add a where not exists clause to the query.
	WhereNotExists(provider QueryProvider[T], where ...WhereJoinType) QueryBuilder[T]

	// OrWhereNotExists 在查询中添加 where not exists 子句
	// Add a where not exists clause to the query.
	OrWhereNotExists(provider QueryProvider[T]) QueryBuilder[T]

	// Union 在查询中添加联合语句
	// Add a union statement to the query.
	Union(builder QueryBuilder[T], unionType ...UnionJoinType) QueryBuilder[T]

	// UnionAll 在查询中添加 union all 语句
	// Add a union all statement to the query.
	UnionAll(builder QueryBuilder[T]) QueryBuilder[T]

	// UnionByProvider 在查询中添加联合语句，并order by
	// Add a union statement to the query and order by.
	UnionByProvider(builder QueryProvider[T], unionType ...UnionJoinType) QueryBuilder[T]

	// UnionAllByProvider 在查询中添加 union all 语句，并order by
	// Add a union all statement to the query and order by.
	UnionAllByProvider(builder QueryProvider[T]) QueryBuilder[T]

	// GroupBy 在查询中添加“group by”子句
	// Add a "group by" clause to the query.
	GroupBy(columns ...string) QueryBuilder[T]

	// Having 在查询中添加“有”子句
	// Add a "having" clause to the query.
	Having(column string, args ...any) QueryBuilder[T]

	// OrHaving 在查询中添加“或有”子句
	// Add an "or having" clause to the query.
	OrHaving(column string, args ...any) QueryBuilder[T]

	// OrderBy 在查询中添加 “order by” 子句
	// Add an "order by" clause to the query.
	OrderBy(column string, columnOrderType ...OrderType) QueryBuilder[T]

	// OrderByDesc 向查询中添加降序 “order by” 子句
	// Add a descending "order by" clause to the query.
	OrderByDesc(column string) QueryBuilder[T]

	// InRandomOrder 将查询的结果按随机顺序排列
	// Put the query's results in random order.
	InRandomOrder(orderFunc ...OrderType) QueryBuilder[T]

	// When 如果给定的 “值” 为真，则应用回调的查询更改
	// Apply the callback's query changes if the given "value" is true.
	When(condition bool, callback QueryCallback[T], elseCallback ...QueryCallback[T]) QueryBuilder[T]

	// ToSql 获取查询的 SQL 表示
	// get the SQL representation of the query.
	ToSql() string

	// GetBindings 获取扁平数组中的当前查询值绑定
	// get the current query value bindings in a flattened array.
	GetBindings() (results []any)

	// Offset 设置查询的 “Offset” 值
	// Set the "offset" value of the query.
	Offset(offset int64) QueryBuilder[T]

	// Skip 设置查询 “Skip” 值的别名
	// Alias to set the "offset" value of the query.
	Skip(offset int64) QueryBuilder[T]

	// Limit  设置查询的“limit”值
	// Set the "limit" value of the query.
	Limit(num int64) QueryBuilder[T]

	// Take 设置查询 “limit” 值的别名
	// Alias to set the "limit" value of the query.
	Take(num int64) QueryBuilder[T]

	// WithPagination 设置给定页面的 “limit” 值和 “offset” 值
	// Set the limit and offset for a given page.
	WithPagination(perPage int64, current ...int64) QueryBuilder[T]

	// SelectSql 获取此 query builder 的当前规范形成的完整 SQL 字符串。
	// Gets the complete SQL string formed by the current specifications of this query builder.
	SelectSql() (string, []any)

	// SelectForUpdateSql 将此实例转换为 SQL 中的 UPDATE 字符串
	// Converts this instance into an UPDATE string in SQL.
	SelectForUpdateSql() (string, []any)
	CreateSql(value Fields, insertType2 ...InsertType) (sql string, bindings []any)
	InsertSql(values []Fields, insertType2 ...InsertType) (sql string, bindings []any)
	InsertIgnoreSql(values []Fields) (sql string, bindings []any)
	InsertReplaceSql(values []Fields) (sql string, bindings []any)
	DeleteSql() (sql string, bindings []any)
	UpdateSql(value Fields) (sql string, bindings []any)

	// Bind 注册查询构造器
	// binding Query executor.
	Bind(executor QueryExecutor[T]) QueryBuilder[T]
}

type ModelContext interface {
	Set(fields Fields)
	Get(string string) any
}

// RelationType 关联关系
type RelationType string

type RelationCollector func(keys []any) map[string][]any
type ForeignKeysCollector[T any] func(item *T) any
type RelationSetter[T any] func(item *T, value []any)

type Relation[T any, P any] interface {
	GetRelationCollector() RelationCollector
	GetForeignKeysCollector() ForeignKeysCollector[P]
	GetRelationSetter() RelationSetter[P]
	GetRelation() RelationType
}
