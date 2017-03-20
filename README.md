# go-common
    go web端的公共类
    
## 文档
    
    提供common基础类库
        asyncThreadProcess.go 提供异步线程处理类,待完善
        
#### fileProcess.go
    1. CreateFileProcess
        文件夹创建方法 接收参数文件夹路径  返回错误对象
        根据文件夹路径创建文件,如文件存在则不做任何操作
        
    2. PathExists
        判断文件或文件夹是否存在 接收参数文件或文件夹路径 返回是否存在及错误对象 true表示存在
    
    3. FileCreateAndWrite
        写文件
        输入参数:文件内容 写入文件的路劲(包含文件名) 是否追加写入
        输出参数:错误对象
        
    4. ReadFileByLine
        文件读取逐行进行读取
        输入参数: 文件路劲
        输出参数: 字符串数组(数组每一项对应文件的每一行) 错误对象
        
    5. FileOpen
        文件打开
        输入参数:文件路径 是否追加
        输出参数:文件对象 错误对象
        
#### mysqlDB.go

    1. MySqlDBStruct
        数据库连接对象
        type MySqlDBStruct struct {
        	DbUser string //数据库用户名
        	DbHost string //数据库地址
        	DbPort int    //数据库端口
        	DbPass string //数据库密码
        	DbName string //数据库库名
        }
    
    2. MySqlSQlConntion
        数据库连接
        输入参数：数据库对象
        输出对象：数据库连接对象
        
    3. MySqlClose
        数据库关闭
        输入参数：数据库连接对象
        
    4. MySqlSelect
        查询方法
        输入参数: dbs数据库连接对象 model数据库对象 sqlStr 要执行的sql语句 param执行SQL的语句参数化传递
        输出参数: 查询返回条数  错误对象输出
        
    5. MySqlSqlExec
        数据库运行方法
        输入参数: dbs数据库连接对象 model数据库对象 sqlStr 要执行的sql语句  param执行SQL的语句参数化传递
        输出参数: 执行结果对象  错误对象输出
        
#### servicePIDProcess.go
    
    1. WritePid
        写PID文件
        输入参数：文件路径
        
    2. CheckPid
        检查pid文件是否存在，pid文件中的进程是否存在
        输入参数：pid文件路径
        输出参数：bool类型（true： 文件不存在或者进程不存在 false: 进程已存在）
        
    3. RmPidFile
        删除PID文件
        输入参数：pid文件路径
