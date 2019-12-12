package concurrent

import (
	"sync"
	"testing"
)


func TestSafemap_Set(t *testing.T) {
   m := initMap()
   m.Set("two",&testStruct{
	   str: "string value2",
	   i: 1,
	   arr : []string{"arr10","arr20"},
   })
}

func TestSafemap_Get(t *testing.T) {
	m := initMap()
	v, ok := m.Get("one")
	if ok{
        ts := v.(*testStruct)
        ts.i = 1
        if ts.i != 1 || ts.str != "string value" || ts.arr[0] != "arr1"{
			t.Fail()
		}
	}else{
		t.Fail()
	}
}

func TestSafemap_Del(t *testing.T) {
	m := initMap()
	m.Del("one")

	_, ok := m.Get("one")
	if ok{
		t.Fail()
	}
}

func TestSafemap_GetKeys(t *testing.T) {
	m := initMap()
	keys := m.GetKeys()
	k := make([]string,len(keys))
	for i,key := range keys{
		k[i] = key.(string)
	}
	if k[0] == "one" && k[1] == "two" || k[1] == "one" && k[0] == "two"{

	}else{
		t.Fail()
	}
}



type testStruct struct{
	str string
	i int
	arr []string
}
func initMap() *safemap{
	m:= &safemap{
		data: make(map[interface{}]interface{}),
		lock: new(sync.RWMutex),
	}

	m.Set("one",&testStruct{
		str: "string value",
		i: 1,
		arr : []string{"arr1","arr2"},
	})

	m.Set("two",&testStruct{
		str: "string value2",
		i: 1,
		arr : []string{"arr11","arr21"},
	})

	return m
}