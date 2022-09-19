package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Data struct {
	Value   string
	TimeVal time.Time
}

type Cache struct {
	Key  map[string]*Data
	lock sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		Key: make(map[string]*Data),
	}
}

//Add 添加缓存
// key 缓存key值  Value 缓存值  second 过期时间 0 代表永不过期
func (c *Cache) Add(key, value string, second int) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	t := time.Time{}
	if second > 0 {
		t = time.Now()
		t = t.Add(time.Duration(second) * time.Second)
	}
	data := &Data{
		Value:   value,
		TimeVal: t,
	}
	c.Key[key] = data
}

//Delete 删除数组元素
//key 为空值 则删除整个缓存
func (c *Cache) Delete(key string) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if len(key) > 0 {
		delete(c.Key, key)
	} else {
		c.Key = make(map[string]*Data)
	}

}

//Update 更新缓存
// second 为0时，不更新缓存时间 second > 0 时，更新缓存过期时间为当前时间 + second秒
func (c *Cache) Update(key, val string, second int) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if ok, _ := c.Select(key); ok {
		if second > 0 {
			t := time.Now().Add(time.Duration(second) * time.Second)
			c.Key[key].Value = val
			c.Key[key].TimeVal = t
		} else {
			c.Key[key].Value = val
		}
	}
	return true
}

//Select 查询数组
//result bool是否存在 string 对应value值
func (c *Cache) Select(key string) (bool, string) {
	if _, ok := c.Key[key]; !ok {
		return false, ""
	}
	if !c.Key[key].TimeVal.IsZero() && time.Now().After(c.Key[key].TimeVal) {
		return false, ""
	}
	return true, c.Key[key].Value
}

func cacheTest() {
	cache := NewCache()
	for i := 0; i < 15; i++ {
		key := "book" + strconv.Itoa(i)
		value := "书籍" + strconv.Itoa(i)
		cache.Add(key, value, i)
	}

	key := "book0"
	ok, val := cache.Select(key)
	fmt.Printf("key:%s,value:%s,：%v \n", key, val, ok)

	time.Sleep(1 * time.Second)

	key = "book1"
	ok, val = cache.Select(key)
	fmt.Printf("key:%s,value:%s,：%v \n", key, val, ok)

	key = "book3"
	ok, val = cache.Select(key)
	fmt.Printf("key:%s,value:%s,：%v \n", key, val, ok)

	cache.Update(key, "GO", 0)

	time.Sleep(2 * time.Second)

	ok, val = cache.Select(key)
	fmt.Printf("key:%s,value:%s,：%v \n", key, val, ok)

	key = "book14"
	ok, val = cache.Select(key)
	fmt.Printf("key:%s,value:%s,：%v \n", key, val, ok)

	cache.Update(key, "JAVA", 0)
	fmt.Println(time.Now())

	ok, val = cache.Select(key)
	fmt.Printf("key:%s,value:%s,更新为java：%v \n", key, val, ok)

	cache.Update(key, "PHP", 1)

	time.Sleep(2 * time.Second)

	ok, val = cache.Select(key)
	fmt.Printf("key:%s,value:%s,更新之后过期：%v \n", key, val, ok)

	key = "book0"
	ok, val = cache.Select(key)
	fmt.Printf("key:%s,value:%s,永不过期：%v \n", key, val, ok)

}
