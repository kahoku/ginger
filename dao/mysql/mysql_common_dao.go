package mysql

import (
	"context"
	"errors"
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
)

// sql builder 映射的表struct应该实现的接口，实现该结构可使用一些通用的curd方法
type DaoMysqlSchema interface {
	TableName() string
}

// 以下提供通用的几个curd方法，具体构建方式可查看本目录的README
/*
根据条件获取单条数据
@param tableName 	string					查询的表名
@param where 		map[string]interface{}	查询条件
@param selectField 	[]string				查询选择返回的字段
@param result 		DaoMysqlSchema			带表结构存储结果的指针，接收返回的数据，实现DaoMysqlSchema接口
*/
func GetOne(tableName string, where map[string]interface{}, selectField []string, result DaoMysqlSchema) error {
	if nil == Db {
		return errors.New("mysql.DB object couldn't be nil")
	}
	condition, values, err := builder.BuildSelect(tableName, where, selectField)
	if nil != err {
		return err
	}

	row, err := Db.Query(condition, values...)
	if nil != err || nil == row {
		return err
	}
	defer row.Close()
	return scanner.Scan(row, result)
}

/*
根据条件获取多条数据
@param tableName 	string					查询的表名
@param where 		map[string]interface{}	查询条件
@param selectField 	[]string				查询选择返回的字段
@param results 		interface{}				带表结构存储结果的指针数组，接收返回的数据，接收results应与table schema struct相对应
*/
func GetMulti(tableName string, where map[string]interface{}, selectField []string, results interface{}) error {
	if nil == Db {
		return errors.New("mysql.DB object couldn't be nil")
	}
	condition, values, err := builder.BuildSelect(tableName, where, selectField)
	if nil != err {
		return err
	}

	rows, err := Db.Query(condition, values...)
	if nil != err || nil == rows {
		return err
	}
	defer rows.Close()
	return scanner.Scan(rows, results)
}

/*
插入数据
@param tableName 	string						表名
@param data 		[]map[string]interface{}	插入数据

@return LastInsertId int64						返回最新的插入id
*/
func Insert(tableName string, data []map[string]interface{}) (int64, error) {
	if nil == Db {
		return -1, errors.New("sql.DB object couldn't be nil")
	}
	condition, values, err := builder.BuildInsert(tableName, data)
	if nil != err {
		return -1, err
	}
	result, err := Db.Exec(condition, values...)
	if nil != err || nil == result {
		return -1, err
	}
	return result.LastInsertId()
}

/*
插入数据,已存在则忽略
@param tableName 	string						表名
@param data 		[]map[string]interface{}	插入数据

@return LastInsertId int64						返回最新的插入id
*/
func InsertIgnore(tableName string, data []map[string]interface{}) (int64, error) {
	if nil == Db {
		return -1, errors.New("sql.DB object couldn't be nil")
	}
	condition, values, err := builder.BuildInsertIgnore(tableName, data)
	if nil != err {
		return -1, err
	}
	result, err := Db.Exec(condition, values...)
	if nil != err || nil == result {
		return -1, err
	}
	return result.LastInsertId()
}

/*
插入数据,已存在则替换
@param tableName 	string						表名
@param data 		[]map[string]interface{}	插入数据

@return LastInsertId int64						返回最新的插入id
*/
func InsertReplace(tableName string, data []map[string]interface{}) (int64, error) {
	if nil == Db {
		return -1, errors.New("sql.DB object couldn't be nil")
	}
	condition, values, err := builder.BuildReplaceInsert(tableName, data)
	if nil != err {
		return -1, err
	}
	result, err := Db.Exec(condition, values...)
	if nil != err || nil == result {
		return -1, err
	}
	return result.LastInsertId()
}

/*
更新数据
@param tableName 	string						表名
@param where 		map[string]interface{}		查询条件
@param data 		[]map[string]interface{}	更新数据

@return RowsAffected int64						更新影响的行数
*/
func Update(tableName string, where, data map[string]interface{}) (int64, error) {
	if nil == Db {
		return -1, errors.New("sql.DB object couldn't be nil")
	}
	condition, values, err := builder.BuildUpdate(tableName, where, data)
	if nil != err {
		return -1, err
	}
	result, err := Db.Exec(condition, values...)
	if nil != err {
		return -1, err
	}
	return result.RowsAffected()
}

/*
删除数据
@param tableName 	string						表名
@param where 		map[string]interface{}		查询条件

@return RowsAffected int64						删除影响的行数
*/func Delete(tableName string, where map[string]interface{}) (int64, error) {
	if nil == Db {
		return -1, errors.New("sql.DB object couldn't be nil")
	}
	condition, values, err := builder.BuildDelete(tableName, where)
	if nil != err {
		return -1, err
	}
	result, err := Db.Exec(condition, values...)
	if nil != err {
		return -1, err
	}
	return result.RowsAffected()
}

/*
聚合查询获取符合条件的行数
@param tableName 	string						表名
@param where 		map[string]interface{}		查询条件

@return count 		int64						符合条件的行数
*/func GetCount(tableName string, where map[string]interface{}) (int64, error) {
	if nil == Db {
		return -1, errors.New("sql.DB object couldn't be nil")
	}
	res, err := builder.AggregateQuery(context.TODO(), Db, tableName, where, builder.AggregateCount("*"))
	if nil != err {
		return -1, err
	}

	return res.Int64(), err
}
