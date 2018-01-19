/**
mySql数据库组件类
创建人：邵炜
创建时间：2017年03月11日15:55:14
*/
package gutil

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/**
数据库连接对象
创建人：邵炜
创建时间：2017年03月11日16:01:20
*/
type MySqlDBStruct struct {
	DbUser       string //数据库用户名
	DbHost       string //数据库地址
	DbPort       int    //数据库端口
	DbPass       string //数据库密码
	DbName       string //数据库库名
	MaxOpenConns int    // 用于设置最大打开的连接数，默认值为0表示不限制
	MaxIdleConns int    // 用于设置闲置的连接数
}

/**
数据库连接
创建人：邵炜
创建时间：2017年03月11日15:56:06
输入参数：数据库对象
输出对象：数据库连接对象
*/
func MySqlSQlConntion(model MySqlDBStruct) (*sql.DB, error) {
	dbClause := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&allowAllFiles=true", model.DbUser, model.DbPass, model.DbHost, model.DbPort, model.DbName)
	dbs, err := sql.Open("mysql", dbClause)
	if err != nil {
		return nil, err
	}
	dbs.SetMaxOpenConns(model.MaxOpenConns)
	dbs.SetMaxIdleConns(model.MaxIdleConns)
	err = dbs.Ping()
	if err != nil {
		return nil, err
	}
	return dbs, err
}

/**
数据库关闭
创建人：邵炜
创建时间：2017年03月11日16:14:50
输入参数：数据库连接对象
*/
func MySqlClose(dbs *sql.DB) {
	if dbs == nil {
		return
	}
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
	if dbs == nil || dbs.Ping() != nil {
		MySqlClose(dbs)
		dbs, err = MySqlSQlConntion(model)
		if err != nil {
			return nil, err
		}
	}
	if param == nil {
		row, err = dbs.Query(sqlStr)
	} else {
		row, err = dbs.Query(sqlStr, param...)
	}
	if err != nil {
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
	if dbs == nil || dbs.Ping() != nil {
		MySqlClose(dbs)
		dbs, err = MySqlSQlConntion(model)
		if err != nil {
			return nil, err
		}
	}

	if param == nil {
		exec, err = dbs.Exec(sqlStr)
	} else {
		exec, err = dbs.Exec(sqlStr, param...)
	}

	if err != nil {
		return nil, err
	}

	return exec, nil
}

//查询返回map
//create by gloomy 2018-01-16 17:52:15
func MysqlSelectMap(dbs *sql.DB, model MySqlDBStruct, sqlStr string, param ...interface{}) (*[]map[string]string, error) {
	var (
		columnArrayIn *[]string
		dataArrayIn   *[][]string
		err           error
	)
	if param == nil {
		columnArrayIn, dataArrayIn, err = MysqlSelectUnknowColumn(dbs, model, sqlStr)
	} else {
		columnArrayIn, dataArrayIn, err = MysqlSelectUnknowColumn(dbs, model, sqlStr, param...)
	}
	if err != nil {
		return nil, err
	}
	var (
		list     []map[string]string
		dicArray map[string]string
	)
	for _, rowArray := range *dataArrayIn {
		dicArray = make(map[string]string)
		for index, columnName := range *columnArrayIn {
			dicArray[columnName] = rowArray[index]
		}
		list = append(list, dicArray)
	}
	return &list, err
}

// 查询所有字段值
// create by gloomy 2017-5-12 16:38:58
func MysqlSelectUnknowColumn(dbs *sql.DB, model MySqlDBStruct, sqlStr string, param ...interface{}) (*[]string, *[][]string, error) {

	var (
		row             *sql.Rows
		err             error
		columnNameArray []string
		dataList        [][]string
	)
	if dbs == nil || dbs.Ping() != nil {
		MySqlClose(dbs)
		dbs, err = MySqlSQlConntion(model)
		if err != nil {
			return nil, nil, err
		}
	}
	if param == nil {
		row, err = dbs.Query(sqlStr)
	} else {
		row, err = dbs.Query(sqlStr, param...)
	}
	if err != nil {
		return nil, nil, err
	}
	defer row.Close()
	columnNameArray, err = row.Columns()
	if err != nil {
		return nil, nil, err
	}
	values := make([]sql.RawBytes, len(columnNameArray))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for row.Next() {
		var rowData []string
		err = row.Scan(scanArgs...)
		if err != nil {
			continue
		}
		for _, col := range values {
			if col == nil {
				rowData = append(rowData, "")
			} else {
				rowData = append(rowData, string(col))
			}
		}
		dataList = append(dataList, rowData)
	}
	return &columnNameArray, &dataList, nil
}
