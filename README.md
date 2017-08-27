# gutil
    go 公共类
    golang.org/x/crypto 无法下载,可以在下面地中获取所需要的crypto包.并存放到对应的golang.org文件夹的x目录下
    https://github.com/swgloomy/crypto.git
    
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
        
    6. RWFileByWhere
        根据条件读写文件
        输入参数:文件路径 文件写入对象 条件平判断方法
        输出参数:错误对象
        
    7. ReadFileLineNumber
        读取文件行数
            文件路径
            输出行数与错误对象
            
    8. GetMyFileName
        获取文件名称及后缀名 未防止文件无后缀名,固这里返回值为数组对象
            输入参数: 文件路劲或文件名称
            输出参数: 文件名 文件后缀名 数组 第一项为文件名称 第二项为文件后缀名

    9. GetMyAllFileByDir
        根据文件夹路径获取文件夹下所有文件
        输入参数: 文件夹路径
        输出参数: 文件名列表 错误对象
        
#### mysqlDB.go

    1. MySqlDBStruct
        数据库连接对象
        type MySqlDBStruct struct {
        	DbUser string //数据库用户名
        	DbHost string //数据库地址
        	DbPort int    //数据库端口
        	DbPass string //数据库密码
        	DbName string //数据库库名
	        MaxOpenConns int    // 用于设置最大打开的连接数，默认值为0表示不限制
	        MaxIdleConns int    // 用于设置闲置的连接数
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
        
    6. MysqlSelectUnknowColumn
        查询所有字段值
        字段名称 字段值数组 错误对象
        
#### servicePIDProcess.go

    const ProgramServicePIDPath = "./programRunPID.pid" // PID文件生成路径
    当下列方法接收到的传入参数为空时,将使用默认的PID文件路径
    
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
	    MaxOpenConns int    // 用于设置最大打开的连接数，默认值为0表示不限制
	    MaxIdleConns int    // 用于设置闲置的连接数
    }
    
    其余同mysql类
    
    提供方法:
        GpSqlConntion GP数据库连接
        GpSqlClose GP数据库关闭
        GpSqlSelect 查询方法
        GpSqlExec 数据库运行方法
        GPSelectUnknowColumn 查询所有字段值
        
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
         
    2. FtpRemoveFile
         FTP文件删除
         文件名 ftp配置对象
         错误对象
         
    3. FtpRenameFile
        ftp修正远程服务器文件名称
        源文件 修正后的文件名称 ftp配置对象
        错误对象
        
#### fileDataRecording.go
     切块文件写入,防止文件过大 (编写缘由用于文件load进入数据库)
     文件数据记录对象
    type FileDataRecording struct {
    	sync.Mutex                         // 锁
    	F                         *os.File // 文件对象
    	FilePre                   string   // 文件开头字符串
    	Fn                        string   // 文件路径
    	Bytes                     int      // 文件大小
    	Seq                       int      // 第几个
    	FileProgram               string   // 文件存放路径
    	MaxFileDataRecordingBytes int      // 文件大小
    }
    
    const maxFileDataRecordingBytes = 1000000 // 默认文件大小
    
    1. OpenLoadFile
         打开文件数据记录
         文件存放目录地址 文件开头字符串 文件大小
         文件数据对象
        
    2. Exit
         文件退出
        
    3. Close
        文件关闭
        
    4. Rotate
        文件切换
        
    5. CreateNewFile
        创建新文件
        错误对象
        
    6. WriteData
        写入数据
        需要写入的数据
        错误对象
        
    7. FileList
        获取所有完成的文件列表
        文件列表
        
    8. RemoveOldFileList
        删除过期文件
        几天前
        
#### excelUtil.go

    1. ReadExcel
       excel数据获取
       sheet名称 数据内容 错误对象
       
    2. ExcelSave
       excel保存
       
#### sftpUtil.go

    // sftp配置
    type SftpConfigStruct struct {
    	Account      string // 登录用户名
    	Password     string // 登录密码
    	Port         int    // 服务器端口
    	ConntionSize int    // MaxPacket sets the maximum size of the payload
    	Addr         string // 连接地址
    }
    
    1. SftpClose 
        sftp 关闭
        
    2. SftpReadDir
        sftp读取文件夹内容
        
#### XML帮助类

    1.XmlContentReplace
        生成xml文件修正xml节点内容
        
#### watchFileUtil.go
    如文件上传完成,统一使用文件修改回调方法.watchFile方法会保证文件上传完毕进行回调
    1.WatchFile
        文件夹监控方法
            监控文件夹路径
            匹配字段(例: match*  则匹配任何match开头的. *为全匹配)
            删除文件回调方法
            文件修改回调方法
            文件改名回调方法
            文件创建回调方法

#### fileSendUtil.go
    http文件上传
    1.HttpSendFile
        文件发送处理方法
            发送http地址
            文件路径
            文件存放变量名