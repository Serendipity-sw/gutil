# go-common
    go web端的公共类
    
## 文档
    
    提供common基础类库
        asyncThreadProcess.go 提供异步线程处理类,待完善
        
#### const.go
    1.WithNanos 时间格式化字符串 2006-01-02 15:04:05
        
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
        
    5. RWFileByWhere
        根据条件读写文件
        输入参数:文件路径 文件写入对象 条件平判断方法
        输出参数:错误对象
        
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
#### dateProcess.go

    1.DateFormat
        时间格式化处理
        输入参数:需要格式化的时间 格式化方式 示例yyyy-MM-dd hh:mm:ss.tttttt   2017-03-22 10:21:55.379415
        
#### redisHelp.go
    
    1. OpenRedis
        redis通道开启
        addr IP地址+端口     idx 仓库数  
        
    2. CloseRedis
        redis通道关闭
        
    3. SetRedisCache
        设置redis缓存
        key 存储键名    value 存储值    cacheBssSeconds 存储时间(单位秒)
        
    4. GetRedisCache
        获取redis缓存
        key 存储键名
        
#### decryptionProcess.go

    1. AesEncrypt
        字符串加密
        输入参数: 需要加密的字符串
        输出参数: 加密后字符串 错误对象
        
    2. AesDecrypt
        字符串解密
        输入参数: 需要解密的字符串  解密后字符串长度
        输出参数: 解密后字符串  错误对象
        
#### gpDB.go
    
    // GP数据库连接对象
    // create by gloomy 2017-3-30 15:27:26
    type GpDBStruct struct {
    	DbUser string //数据库用户名
    	DbHost string //数据库地址
    	DbPort int    //数据库端口
    	DbPass string //数据库密码
    	DbName string //数据库库名
    }
    
    其余同mysql类
    
    提供方法:
        GpSqlConntion GP数据库连接
        GpSqlClose GP数据库关闭
        GpSqlSelect 查询方法
        GpSqlExec 数据库运行方法
        
#### mathProcess.go
    
    1. Rounding
        四舍五入取舍
        
    2. RoundingByInt
        四舍五入取舍
        除数 被除数 取舍几位
        
    3. RoundingPercentageByInt
        四舍五入取舍 百分比
        除数 被除数 取舍几位
        
#### ftpHelpProcess.go

    FTP帮助类实体
    type FtpHelpStruct struct {
    	IpAddr string // ip 地址
    	Port int // 端口
    	TimeOut time.Duration // 超时时间
    	UserName string // 用户名
    	PassWord string // 密码
    	FilePaths string // 目标服务器路径
    }
    
    1. FtpFileStor
         FTP文件传输        
         FTP配置实体 文件内容 创建目标服务器的文件名
         错误对象