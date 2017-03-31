/**
mySql数据库组件类
创建人：邵炜
创建时间：2017年03月11日15:55:14
*/
package common

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/smtc/glog"
)

/**
数据库连接对象
创建人：邵炜
创建时间：2017年03月11日16:01:20
*/
type MySqlDBStruct struct {
	DbUser string //数据库用户名
	DbHost string //数据库地址
	DbPort int    //数据库端口
	DbPass string //数据库密码
	DbName string //数据库库名
}

/**
数据库连接
创建人：邵炜
创建时间：2017年03月11日15:56:06
输入参数：数据库对象
输出对象：数据库连接对象
*/
func MySqlSQlConntion(model MySqlDBStruct) *sql.DB {
	dbClause := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&allowAllFiles=true", model.DbUser, model.DbPass, model.DbHost, model.DbPort, model.DbName)
	dbs, err := sql.Open("mysql", dbClause)
	if err != nil {
		glog.Error("mysql can't connection dbClause: %s err: %s \n", dbClause, err.Error())
		return nil
	}
	err = dbs.Ping()
	if err != nil {
		glog.Error("mysql can't ping dbClause: %s err: %s \n", dbClause, err.Error())
		return nil
	}
	return dbs
}

/**
数据库关闭
创建人：邵炜
创建时间：2017年03月11日16:14:50
输入参数：数据库连接对象
*/
func MySqlClose(dbs *sql.DB) {
	dbs.Close()
}

/**
查询方法
创建人:邵炜
创建时间:2015年12月29日17:26:41
修正时间：2017年03月11日16:21:45
输入参数: dbs数据库连接对象 model数据库对象 sqlStr 要执行的sql语句 param执行SQL的语句参数化传递
输出参数: 查询返回条数  错误对象输出
*/
func MySqlSelect(dbs *sql.DB, model MySqlDBStruct, sqlStr string, param ...interface{}) (*sql.Rows, error) {

	var (
		row *sql.Rows
		err error
	)
	err = dbs.Ping()
	if err != nil {
		glog.Error("mysql can't ping %s \n", err.Error())
		MySqlClose(dbs)
		dbs = MySqlSQlConntion(model)
	}
	if param == nil {
		row, err = dbs.Query(sqlStr)
	} else {
		row, err = dbs.Query(sqlStr, param...)
	}
	if err != nil {
		glog.Error("mysql query can't select sql: %s err: %s \n", sqlStr, err.Error())
		return nil, err
	}
	return row, nil
}

/**
数据库运行方法
创建人:邵炜
创建时间:2015年12月29日17:33:06
修正时间：2017年03月11日16:21:36
输入参数: dbs数据库连接对象 model数据库对象 sqlStr 要执行的sql语句  param执行SQL的语句参数化传递
输出参数: 执行结果对象  错误对象输出
*/
func MySqlSqlExec(dbs *sql.DB, model MySqlDBStruct, sqlStr string, param ...interface{}) (sql.Result, error) {
	var (
		exec sql.Result
		err  error
	)

	err = dbs.Ping()

	if err != nil {
		glog.Error("mysql can't ping %s \n", err.Error())
		MySqlClose(dbs)
		dbs = MySqlSQlConntion(model)
	}

	if param == nil {
		exec, err = dbs.Exec(sqlStr)
	} else {
		exec, err = dbs.Exec(sqlStr, param...)
	}

	if err != nil {
		glog.Error("mysql exec can't carried out sql: %s err: %s \n", sqlStr, err.Error())
		return nil, err
	}

	return exec, nil
}
