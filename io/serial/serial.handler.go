package serial

import (
	"bytes"
	"errors"
	"fmt"
	log "github.com/jerryzhengj/utils/log/zap"
	s "github.com/tarm/serial"
	"go.uber.org/zap"
	"io"
	"time"
)

func Open(options *Options)(session *Port){
	log.Info(fmt.Sprintf("Open connection %s:%d", options.PortName, options.BaudRate))

	var readTimeout int64 = -1
	if session.Opts.readMode != ReadModeActive {
		readTimeout = options.Timeout
	}
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
		Conn: conn,
		Readable: true,
		ReadChan: make(chan byte,1024),
	}
    if session.Opts.readMode == ReadModeActive{
		go readToChannel(session)
	}
	return session
}

func readToChannel(session *Port){
	p := make([]byte, 1)
	for session.Readable {
		readSize, err := session.Conn.Read(p)
		if err == nil && readSize > 0{
			session.ReadChan<- p[0]
		}else if err != nil && err != io.EOF {
			log.Errorf("readToChannel failed with error:%s",err)
			panic(err)
		}
	}
}

func (session *Port)Close(){
	session.Readable = false
	time.Sleep(100)
	close(session.ReadChan)
	session.Conn.Close()
}

func (session *Port)Read(startTimeMilSec int64,size int)([]byte,error){
    if session.Opts.readMode == ReadModeActive {
        return session.readFromChannel(startTimeMilSec,size)
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
		readSize, err := session.Conn.Read(p)
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
		}else if session.Opts.Timeout > 0 && time.Now().UnixNano()/1000000 - startTimeMilSec > session.Opts.Timeout{
			return nil,errors.New("read serial timeout")
		}
	}

	data := buffer.Bytes()
	log.Debugf("readFromSerial bytes:%v",data)
	return data,nil
}

func (session Port)readFromChannel(startTimeMilSec int64,size int)([]byte,error) {
	hasRead := 0
	buffer := make([]byte, size)
	for {
		select {
		case b := <-session.ReadChan:
			buffer[hasRead] = b
			hasRead++
			if hasRead >= size {
				log.Debugf("readFromSerial bytes:%v",buffer)
				return buffer, nil
			}
		case <-time.After(time.Duration(startTimeMilSec) * time.Millisecond):
			return buffer, ReadTimeoutErr
		}
	}
}
