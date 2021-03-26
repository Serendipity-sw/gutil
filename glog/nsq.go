package glog

/*
import (
	"errors"
	nsq "github.com/bitly/go-nsq"
	"io"
	"net"
	//"syscall"
)

type nsqLogger struct {
	Logger

	localAddr string

	rtSeconds int64
	rtItems   int64
	rtNbytes  int64
}

type nsqWriter struct {
	*nsq.Producer
	chn string
}

func (nl *nsqLogger) Close() {
	for i := DebugLevel; i < LevelCount; i++ {
		nl.out.out[i].Close()
	}
}

// options:
//   nsqdAddr: string, host:port
//   localAddr: string
//   device_id: int
func createNsqLogger(options map[string]interface{}) *nsqLogger {
	var (
		ok        bool
		err       error
		flag      int
		nsqdAddr  string
		localAddr string
	)
	if flag, ok = options["flag"].(int); !ok {
		flag = Ldate | Ltime
	}

	if localAddr, ok = options["localAddr"].(string); !ok {
		localAddr, err = getLocalAddr()
		if err != nil {
			panic("get local address failed\n")
		}
	}
	nl := &nsqLogger{
		Logger{
			flag: flag,
		},
		localAddr,
		0, 0, 0,
	}

	if nsqdAddr, ok = options["nsqdAddr"].(string); !ok {
		panic("options[\"nsqdAddr\" must be string\n")
	}

	nl.buildNsqChannel(nsqdAddr)

	return nl
}

// 获得本机网卡ip地址
func getLocalAddr() (string, error) {
	is, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, card := range is {
		// 网卡没有up
		if card.Flags&net.FlagUp == 0 {
			continue
		}
		// loopback地址
		if card.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := card.Addrs()
		if err != nil {
			continue
		}
		return addrs[0].String(), nil
	}

	return "", errors.New("Not found proper interface address.")
}

// 建立于nsq server
func (nl *nsqLogger) buildNsqChannel(nsqdAddr string) {
	var err error

	nl.out.prefix = map[int]string{
		DebugLevel: "DEBUG",
		InfoLevel:  "INFO",
		WarnLevel:  "WARN",
		ErrorLevel: "ERROR",
		FatalLevel: "FATAL",
		PanicLevel: "PANIC",
	}

	nl.out.out = make(map[int]io.WriteCloser)

	for i := DebugLevel; i < LevelCount; i++ {
		nl.out.out[i], err = nl.openPublisher(nsqdAddr, i)
		if err != nil {
			panic("open nsq publisher failed!")
		}
	}
}

func (nl *nsqLogger) openPublisher(nsqdAddr string, level int) (io.WriteCloser, error) {
	conf := nsq.NewConfig()

	chn := nl.out.prefix[level] + "-" + nl.localAddr
	p, err := nsq.NewProducer(nsqdAddr, conf)
	if err != nil {
		println("create producer " + chn + " failed")
		return nil, err
	}
	n := &nsqWriter{p, chn}
	return n, nil
}

// nsqWrite implement io.WriteCloser
func (nw *nsqWriter) Write(data []byte) (n int, err error) {
	err = nw.Publish(nw.chn, data)
	return len(data), err
}

func (nw *nsqWriter) Close() error {
	nw.Stop()
	return nil
}
*/
