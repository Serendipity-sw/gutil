package glog

// 不输出任何日志，仅用于调试时提高性能

type nullLog struct {
}

func (c nullLog) GetPrefix() map[int]string {
	return nil
}

func (c nullLog) Prefix(lv int) string {
	return ""
}

func (c nullLog) SetPrefix(lv int, prefix string) {
}

func (c nullLog) Flags() int {
	return 0
}

func (c nullLog) SetFlags(flag int) {
}

func (c nullLog) Level() int {
	return DebugLevel
}
func (c nullLog) SetLevel(level int) {
}

func (c nullLog) Debug(format string, v ...interface{}) {
}

func (c nullLog) Info(format string, v ...interface{}) {
}

func (c nullLog) Warn(format string, v ...interface{}) {
}

func (c nullLog) Error(format string, v ...interface{}) {
}

func (c nullLog) Fatal(format string, v ...interface{}) {
}

func (c nullLog) Panic(format string, v ...interface{}) {
}

func (c nullLog) Close() {
}

func (c nullLog) Flush() {
}
