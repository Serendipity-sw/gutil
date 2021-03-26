package glog

// 2015-09-30
// 简化一下：
// 当前正在使用的日志命名为.log
// 间隔只能以Hour, 或者以天为单位
//

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"path"
	"path/filepath"
	//"strconv"
	"strings"
	"time"

	"github.com/guotie/days"
)

type fileLogger struct {
	Logger
	dir    string
	format string // file suffix, such as "{{program}}-{{host}}-{{username}}-{{yyyy}}{{mm}}{{dd}}-{{pid}}"

	rotDuration time.Duration
	rot         chan struct{}
	exit        chan struct{}
	exited      chan bool
	//contact     bool
	duration string // must be hour or day
}

var (
	pid      = os.Getpid()
	program  = filepath.Base(os.Args[0])
	host     = "unknownhost"
	userName = "unknownuser"

	prefixFn = map[int]string{
		DebugLevel: "DEBUG",
		InfoLevel:  "INFO",
		WarnLevel:  "WARN",
		ErrorLevel: "ERROR",
		FatalLevel: "FATAL",
		PanicLevel: "PANIC",
	}
)

func init() {
	h, err := os.Hostname()
	if err == nil {
		host = shortHostname(h)
	}

	current, err := user.Current()
	if err == nil {
		userName = current.Username
	}

	// Sanitize userName since it may contain filepath separators on Windows.
	userName = strings.Replace(userName, `\`, "_", -1)
}

// shortHostname returns its argument, truncating at the first period.
// For instance, given "www.google.com" it returns "www".
func shortHostname(hostname string) string {
	if i := strings.Index(hostname, "."); i >= 0 {
		return hostname[:i]
	}
	return hostname
}

/*
func cleanTmpLogs(dir string, contact bool) {
	var (
		err error
		fns []string = make([]string, 0)
	)

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, ".tmp") {
			fns = append(fns, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("clean tmp logs failed: walk dir %s failed: %v\n", dir, err)
	}

	for _, fn := range fns {
		renameTmpLogs(dir, fn, contact)
	}

	return
}

func renameTmpLogs(dir, fn string, contact bool) {
	if len(fn) <= 4 {
		return
	}

	// 不断的尝试rename fn - ".tmp" + xxx + ".log", 直到成功为止
	var (
		seq    int = 1
		err    error
		baseFn = filepath.Base(fn)
		nfn    string
	)

	if len(baseFn) < 4 {
		fmt.Printf("tmp log filename %s %s invalid: filename should contain .tmp\n", fn, baseFn)
		return
	}
	baseFn = baseFn[0 : len(baseFn)-4]

	if contact {
		contactLog(path.Join(dir, baseFn+".log"), fn)
	} else {
		seq = checkSequence(dir, baseFn) + 1
		nfn = path.Join(dir, baseFn+fmt.Sprintf(".seq%03d", seq)+".log")

		if err = os.Rename(fn, nfn); err != nil {
			fmt.Printf("Cannot rename tmp log file %s, remove it\n", fn, err)
			os.Remove(fn)
		}
	}

	return
}

// 检查是否由以tmp结尾的文件, 可能有雨程序崩溃, 存在tmp结尾的文件，把这些文件转换为对应的log后缀
func checkSequence(dir, fnDate string) int {
	var (
		err      error
		seq      int
		sequence int
		fns      []string = make([]string, 0)
	)

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		baseFn := filepath.Base(path)
		if strings.Contains(baseFn, fnDate) {
			fns = append(fns, baseFn)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("check sequence: walk dir %s failed: %v\n", dir, err)
	}

	for _, fn := range fns {
		segs := strings.Split(fn, ".")
		if len(segs) < 2 {
			continue
		}
		sseq := segs[len(segs)-2]
		if strings.HasPrefix(sseq, "seq") == false {
			continue
		}
		seq, err = strconv.Atoi(sseq[3:])
		if err != nil {
			fmt.Printf("convert sequence %s failed: %v\n", sseq[3:], err)
			continue
		}
		if sequence < seq {
			sequence = seq
		}
	}

	return sequence
}
*/

func contactLog(logf, tmp string) {
	file, err := os.OpenFile(logf, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm|os.ModeTemporary)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write([]byte("\n\n--------------------------------------------------------\n\n\n"))

	tmpFile, err := os.Open(tmp)
	if err != nil {
		panic(err)
	}
	defer tmpFile.Close()

	buff := make([]byte, 256*1024)
	n := 0
	for err != io.EOF {
		n, err = tmpFile.Read(buff)
		if err != nil && err != io.EOF {
			log.Printf("Read file %s failed: %s\n", tmp, err.Error())
			return
		}
		file.Write(buff[0:n])
	}
	//buff, _ := ioutil.ReadAll(tmpFile)
	//file.Write(buff)
	os.Remove(tmp)
}

// options:
//    flag: int
//    prefix: map[int]string
//    dir: string
//    contcat:
//    format: string
//
func createFileLogger(options map[string]interface{}) *fileLogger {
	var (
		ok   bool
		err  error
		flag int
		//sequence int
		//contact  bool
		dir      string
		fnSuffix string
		prefix   map[int]string
		duration = "day" // 按天来rotate日志或按小时rotate日志
	)

	if flag, ok = options["flag"].(int); !ok {
		flag = Ldate | Ltime
	}

	// 使用不同文件来记录不同等级的log，不需要加前缀
	if prefix, ok = options["prefix"].(map[int]string); !ok {
		prefix = nil
	}

	if dir, ok = options["dir"].(string); !ok {
		dir = "./logs"
	}

	if duration, ok := options["duration"].(string); !ok {
		duration = "day"
	} else {
		duration = strings.ToLower(strings.TrimSpace(duration))
		if duration != "hour" && duration != "day" {
			log.Printf("duration [%s] invalid, must be day or hour, set to day\n", duration)
			duration = "day"
		}
	}
	if fnSuffix, ok = options["suffix"].(string); !ok {
		if duration == "day" {
			fnSuffix = "-{{yyyy}}{{mm}}{{dd}}"
		} else {
			fnSuffix = "-{{yyyy}}{{mm}}{{dd}}-{{HH}}"
		}
	}

	// 判断目录是否存在
	_, err = os.Stat(dir)
	dirExisted := err == nil || os.IsExist(err)
	if !dirExisted {
		os.Mkdir(dir, os.ModePerm)
	}

	// 2015-09-30 不存在tmp logs, sequence也不需要了
	// 清理tmp log文件
	//cleanTmpLogs(dir, contact)
	//sequence = checkSequence(dir, formatSuffix(fnSuffix))

	fl := &fileLogger{
		Logger{
			flag: flag,
		},
		dir,
		fnSuffix,

		0,
		make(chan struct{}),
		make(chan struct{}),
		make(chan bool),
		//contact,
		duration,
	}

	if duration == "day" {
		fl.rotDuration = time.Duration(86400) * time.Second
	} else {
		fl.rotDuration = time.Duration(3600) * time.Second
	}

	err = fl.buildFileOut(prefix)
	if err != nil {
		panic(err)
	}

	go fl.rotate()
	return fl
}

func CreateDirIfNotExist(dir string) error {
	_, err := os.Stat(dir)
	// 目录存在
	if err == nil {
		return nil
	}

	// 其他错误
	if !os.IsNotExist(err) {
		return err
	}
	return os.MkdirAll(dir, os.ModeDir|os.ModePerm)
}

func (fl *fileLogger) buildFileOut(prefix map[int]string) (err error) {
	if err = CreateDirIfNotExist(fl.dir); err != nil {
		return
	}

	// avoid panic on nil map
	if fl.out.prefix = prefix; prefix == nil {
		fl.out.prefix = make(map[int]string)
	}

	fl.out.out, err = fl.openLogFiles()
	return
}

func (fl *fileLogger) openLogFiles() (wr map[int]io.WriteCloser, err error) {
	var (
		f  *os.File
		fn string
	)

	//suffix := formatSuffix(fl.format)
	wr = make(map[int]io.WriteCloser)

	for i := DebugLevel; i < LevelCount; i++ {
		fn = path.Join(fl.dir, prefixFn[i]+".log")
		//log.Printf("open log level %d, fn=%s\n", i, fn)
		if f, err = os.OpenFile(fn, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm); err != nil {
			log.Printf("open log file %s failed: %v\n", fn, err)
			continue
		}

		wr[i] = f
	}

	return
}

func formatSuffix(format string, tm time.Time) (res string) {
	if format == "" {
		return
	}

	res = strings.Replace(format, "{{program}}", program, -1)
	res = strings.Replace(res, "{{host}}", host, -1)
	res = strings.Replace(format, "{{username}}", userName, -1)
	res = strings.Replace(res, "{{yyyy}}", fmt.Sprintf("%d", tm.Year()), -1)
	res = strings.Replace(res, "{{mm}}", fmt.Sprintf("%02d", tm.Month()), -1)
	res = strings.Replace(res, "{{dd}}", fmt.Sprintf("%02d", tm.Day()), -1)
	res = strings.Replace(res, "{{HH}}", fmt.Sprintf("%02d", tm.Hour()), -1)
	res = strings.Replace(res, "{{MM}}", fmt.Sprintf("%02d", tm.Minute()), -1)
	res = strings.Replace(res, "{{SS}}", fmt.Sprintf("%02d", tm.Second()), -1)
	res = strings.Replace(res, "{{pid}}", fmt.Sprint(pid), -1)

	return
}

func (fl *fileLogger) toNextRotateSeconds(now time.Time) int {
	if fl.duration == "hour" {
		return 3600 - int(now.Unix()%3600)
	}

	tom := days.Tomorrow(now).Unix()

	return int(tom - now.Unix())
}

// 2014-10-17 guotie
// TODO: rotate file logs
func (fl *fileLogger) rotate() {
	//tk := time.NewTicker(time.Second * time.Duration(maxCacheSeconds))
	left := fl.toNextRotateSeconds(time.Now())
	//log.Println("left second to timer:", left)
	tmr := time.NewTimer(time.Duration(left) * time.Second)
	go func() {
		for {
			select {
			case <-tmr.C:
				fl.rot <- struct{}{}

				left = fl.toNextRotateSeconds(time.Now())
				tmr.Reset(time.Duration(left) * time.Second)
				//case <-tk.C:
				//	fl.flush()
			}
		}
	}()

	for {
		select {
		case <-fl.rot:
			fl.mu.Lock()
			owr := fl.out.out
			fl.closeLogFiles(owr, "rotate")
			wr, err := fl.openLogFiles()
			if err != nil {
				log.Printf("rotate log files failed: %v\n", err)
			} else {
				fl.out.out = wr
			}
			fl.mu.Unlock()

		case <-fl.exit:
			//tk.Stop()
			tmr.Stop()

			fl.mu.Lock()
			owr := fl.out.out
			fl.closeLogFiles(owr, "exit")
			fl.mu.Unlock()
			close(fl.exited)

			return
		}
	}
}

// mod: 关闭的方式
//	       "rotate": 使用昨天或者上一个小时的时间作为文件后缀
//         "exit":   使用当前时间作为后缀
func (fl *fileLogger) closeLogFiles(fs map[int]io.WriteCloser, mod string) {
	var (
		err     error
		fn, nfn string
		suffix  string
		tm      time.Time
	)

	if mod != "rotate" && mod != "exit" {
		panic(`invalid param mod: must be "rotate" or "exit"!`)
	}

	tm = time.Now()

	if mod != "exit" {
		if fl.duration == "day" {
			tm = days.Yesterday(tm)
		} else {
			tm = tm.Add(-1 * time.Hour)
		}
	}
	suffix = formatSuffix(fl.format, tm)

	//fl.sequence++
	for _, r := range fs {
		f := r.(*os.File)
		f.Close()
		fn = f.Name()
		nfn = fn[0:len(fn)-4] + suffix + ".log"
		err = renameLog(fn, nfn)
		if err != nil {
			log.Println("closeLogFiles failed:", err)
		}
	}
}

func renameLog(ofn, nfn string) error {
	f, err := os.Open(nfn)
	if err == nil {
		//
		f.Close()
		contactLog(nfn, ofn)
		return nil
	}

	if os.IsNotExist(err) {
		return os.Rename(ofn, nfn)
	}

	return err
}

func (fl *fileLogger) Close() {
	fl.exit <- struct{}{}
	<-fl.exited
}
