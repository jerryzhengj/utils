package serial

import (
	"bytes"
	"errors"
	log "github.com/jerryzhengj/utils/log/zap"
	s "github.com/tarm/serial"
	"go.uber.org/zap"
	"io"
	"sync"
	"time"
)

const defaultTimeout = time.Duration(8760) * time.Hour

func Open(options *Options)(session *Port){
	log.Infof("Open connection %s:%d", options.PortName, options.BaudRate)

	var readTimeout int64 = -1
	if options.ReadMode != ReadModeActive {
		readTimeout = options.Timeout
	}
	log.Infof("port readTimeout=%d, readMode:%d", readTimeout, options.ReadMode)
	c := &s.Config{
		Name: options.PortName,
		Baud: options.BaudRate,
		ReadTimeout: time.Duration(readTimeout)* time.Millisecond,
	}
	conn, err := s.OpenPort(c)

	if err != nil {
		log.Error("Open serial port failed",zap.Error(err))
		panic(err)
	}
	session = &Port{
		Opts: *options,
		conn: conn,
		readable: true,
		readChan: make(chan byte,1024),
		lock: new(sync.RWMutex),
	}
    if session.Opts.ReadMode == ReadModeActive{
		log.Infof("start to readToChannel")
		go readToChannel(session)
	}
	return session
}

func readToChannel(session *Port){
	p := make([]byte, 1)
	for session.readable {
		readSize, err := session.conn.Read(p)
		if err == nil && readSize > 0{
			session.readChan<- p[0]
		}else if err != nil && err != io.EOF {
			log.Errorf("readToChannel failed with error:%s",err)
			panic(err)
		}
	}
}

func (session *Port)Close(){
	session.readable = false
	time.Sleep(100)
	close(session.readChan)
	session.conn.Close()
}

func (session *Port)Write(data []byte)(int, error) {
	defer session.lock.Unlock()
	session.lock.Lock()
	return session.conn.Write(data)
}

func (session *Port)Read(startTimeMilSec int64,size int)([]byte,error){
    if session.Opts.ReadMode == ReadModeActive {
        return session.readFromChannel(size)
	}else{
		return session.readFromSerial(startTimeMilSec,size)
	}
}

func (session Port)readFromSerial(startTimeMilSec int64,size int)([]byte,error){
	log.Debugf("readNbytes size:%d", size)
	hasRead := 0
	buffer := bytes.Buffer{}
	for {
		p := make([]byte, size - hasRead)
		readSize, err := session.conn.Read(p)
		log.Debugf("readNbytes size:%d,error=%v", readSize,err)
		if err != nil {
			if err != io.EOF {
				log.Errorf("readNbytesFromSerial failed with error:%s",err)
				panic(err)
			}
		}
		if readSize > 0 {
			hasRead += readSize
			buffer.Write(p[:readSize])
		}

		if hasRead >= size {
			break
		}else if startTimeMilSec > 0 && session.Opts.Timeout > 0 &&
			time.Now().UnixNano()/1000000 - startTimeMilSec > session.Opts.Timeout{
			return nil,errors.New("read serial timeout")
		}
	}

	data := buffer.Bytes()
	log.Debugf("readFromSerial bytes:%v",data)
	return data,nil
}

func (session Port)readFromChannel(size int)([]byte,error) {
	timeout := defaultTimeout
	if session.Opts.Timeout > 0{
		timeout =time.Duration(session.Opts.Timeout) * time.Millisecond
	}

	hasRead := 0
	buffer := make([]byte, size)
	for {
		log.Debugf("start to execute select")
		select {
		case b := <-session.readChan:
			buffer[hasRead] = b
			hasRead++
			if hasRead >= size {
				log.Debugf("readFromChannel bytes:%v",buffer)
				return buffer, nil
			}
		case <-time.After(timeout):
			return buffer, ReadTimeoutErr
		}
	}
}
