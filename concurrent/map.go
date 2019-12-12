package concurrent

import "sync"

type safemap struct{
	data map[interface{}]interface{}
	lock *sync.RWMutex
}


func (s safemap) Get(k interface{})(interface{},bool){
	s.lock.RLock()
	defer s.lock.RUnlock()
	v,ok:=s.data[k]
	return v,ok
}

func (s safemap) Set(k interface{},v interface{}){
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data[k]=v
}

func (s safemap) Del(k interface{}){
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.data,k)
}

func (s safemap) GetKeys()[]interface{}{
	s.lock.Lock()
	defer s.lock.Unlock()

	var keys []interface{}
	for key := range s.data{
		keys = append(keys,key)
	}

	return keys
}