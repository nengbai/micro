package cache

import (
	"encoding/json"
	"fmt"
	"micro/global"
	"strings"
	"time"
)

type SessionInterface interface {
	Put(key string, value interface{}) error
	Get(key string) interface{}
	Remove(key string)
	setSliceMap(m map[string]interface{}, keys []string, value interface{}) map[string]interface{}
	getSliceMap(m map[string]interface{}, keys []string) interface{}
	delSliceMap(m map[string]interface{}, keys []string)
	Destroy() int64
}

type Session struct {
	Name string
	TTL  int64
	SessionInterface
}

func (s *Session) Put(key string, value interface{}) error {
	var h = global.RedisDb
	var bytes []byte
	var err error
	var content string
	m := make(map[string]interface{})
	if h.Exists(s.Name).Val() == 1 {
		content = h.Get(s.Name).Val()
	} else {
		content = "{}"
	}
	fmt.Printf("%s => %s", s.Name, content)
	json.Unmarshal([]byte(content), &m)
	if err != nil {
		return err
	}
	var keys = strings.Split(key, ".")
	var depth = len(keys)
	if depth < 2 {
		m[key] = value
	} else {
		m = setSliceMap(m, keys, value)
	}
	bytes, _ = json.Marshal(m)
	h.Set(s.Name, string(bytes), time.Duration(s.TTL)*time.Second)
	return nil
}

func (s *Session) Get(key string) interface{} {
	var h = global.RedisDb
	var m map[string]interface{}
	content := h.Get(s.Name).Val()
	json.Unmarshal([]byte(content), &m)
	var keys = strings.Split(key, ".")
	var n = len(keys)
	if n < 2 {
		return m[key]
	}
	return getSliceMap(m, keys)
}

func (s *Session) Remove(key string) {
	var h = global.RedisDb
	var m map[string]interface{}
	content := h.Get(s.Name).Val()
	json.Unmarshal([]byte(content), &m)
	var keys = strings.Split(key, ".")
	var n = len(keys)
	if n < 2 {
		delete(m, key)
		return
	}
	delSliceMap(m, keys)
}

func setSliceMap(m map[string]interface{}, keys []string, value interface{}) map[string]interface{} {
	var itMap = m
	var i int
	var limit = len(keys) - 1
	var ok bool
	for i = 0; i < limit; i++ {
		_, ok = itMap[keys[i]]
		if ok {
			fmt.Printf("%s yes\n", keys[i])
		} else {
			itMap[keys[i]] = make(map[string]interface{})
		}
		itMap = itMap[keys[i]].(map[string]interface{})
	}
	itMap[keys[limit]] = value
	return m

}
func getSliceMap(m map[string]interface{}, keys []string) interface{} {
	var itMap = m
	var i int
	var limit = len(keys) - 1
	var v interface{}
	var ok bool
	for i = 0; i < limit; i++ {
		v, ok = itMap[keys[i]]
		if !ok {
			break
		}
		itMap = v.(map[string]interface{})
	}
	v, ok = itMap[keys[i]]
	if !ok {
		return nil
	}
	return v
}
func delSliceMap(m map[string]interface{}, keys []string) {
	var itMap = m
	var i int
	var limit = len(keys) - 1
	var v interface{}
	var ok bool
	for i = 0; i < limit; i++ {
		v, ok = itMap[keys[i]]
		if !ok {
			break
		}
		itMap = v.(map[string]interface{})
	}
	_, ok = itMap[keys[i]]
	if !ok {
		return
	}
	delete(itMap, keys[i])
}

func (s *Session) Destroy() int64 {
	var h = global.RedisDb
	return h.Del(s.Name).Val()
}
