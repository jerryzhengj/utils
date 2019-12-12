package serial


import (
	"errors"
	s "github.com/tarm/serial"
)

const (
	// 打开串口连接后，等到调用读操作时，才开始读从串口读数据
	ReadModePassive = 1+ iota
	// 打开串口连接后，立即开始从串口读数据
	ReadModeActive
)

var ReadTimeoutErr = errors.New("read timeout")

type Port struct {
	// 配置
	opts Options

	// 连接
	conn *s.Port

	readable bool

	readChan chan byte
}

type Options struct {

	// 码率
	BaudRate int

	// 串口
	PortName string

	// 读串口超时时间，单位：毫秒
	Timeout int64

	// 读模式. 1:默认模式，调用readNbytes才开始读  2: 打开端口连接后就持续读
	readMode int
}

type Callback func(data []byte)
