package cache

import (
	"sync"

	"strconv"

	"github.com/garyburd/redigo/redis"
)

//CacheProvider apply cache connection config
type CacheProvider struct {
	IP   string
	Port int
	User string
	Pwd  string
}

var _conn redis.Conn
var _once sync.Once
var _cacheProvider *CacheProvider

//Instance singlon mode
func Instance() *CacheProvider {
	_once.Do(func() {
		_cacheProvider = &CacheProvider{}
		_cacheProvider.IP = "127.0.0.1"
		_cacheProvider.Port = 6379
		conn, err := redis.Dial("tcp", _cacheProvider.IP+":"+strconv.Itoa(_cacheProvider.Port))
		if err != nil {
			panic(err)
		}
		_conn = conn
	})

	return _cacheProvider
}

//SetItem apply for adding data to cache provider
func (p CacheProvider) SetItem(key string, value string) bool {
	Instance()
	_, err := _conn.Do("SET", key, value)
	if err != nil {
		return false
	}
	return true
}

//GetItem apply for getting data from cache provider by key
func (p CacheProvider) GetItem(key string) string {
	Instance()
	v, err := _conn.Do("GET", key)
	if err != nil {
		return ""
	}

	if v != nil {
		return "v.(string)"
	}

	return ""
}
