package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func main() {
	//simpleSetAndGet()
	//setExpireTime()
	//existKey()
	//deleteKey()
	//dataToJson()
	//expireTime()
	pushList()
}

/*
简单的插入Key,根据Key获取指定的value
*/
func simpleSetAndGet() {
	//连接Redis
	c, err := redis.Dial("tcp", "192.168.1.33:6379")
	if err != nil {
		fmt.Println("连接Redis错误", err)
		return
	}
	_, err = c.Do("SET", "mykey", "gfgdg")
	if err != nil {
		fmt.Println("redis set failed", err)
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed", err)
		return
	}
	fmt.Printf("Get mykey: %v \n", username)
}

/*
设置过期时间
*/
func setExpireTime() {
	c, err := redis.Dial("tcp", "192.168.1.33:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "mykey", "superWang", "EX", "5")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	time.Sleep(8 * time.Second)

	username, err = redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
}

func existKey() {
	conn, err := redis.Dial("tcp", "192.168.1.33:6379")

	if err != nil {
		fmt.Println("connect err:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("SET", "mykey", "Yaphet")
	if err != nil {
		fmt.Println("redis set failed", err)
	}

	isKeyExist, err := redis.Bool(conn.Do("EXISTS", "mykey1"))
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("exists or not: %v \n", isKeyExist)
	}

	isKeyExist, err = redis.Bool(conn.Do("EXISTS", "mykey"))
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("exists or not: %v \n", isKeyExist)
	}
}

func deleteKey() {
	c, err := redis.Dial("tcp", "192.168.1.33:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	_, err = c.Do("DEL", "mykey")
	if err != nil {
		fmt.Println("redis delelte failed:", err)
	}

	username, err = redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
}

func dataToJson() {
	c, err := redis.Dial("tcp", "192.168.1.33:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	key := "profile"
	imap := map[string]string{"username": "dfsfg", "phonenumber": "sgsrgr"}
	value, _ := json.Marshal(imap)
	_, err = c.Do("DEL", "profile")
	n, err := c.Do("SETNX", key, value) //这个方法应该是有个Bug，不会覆盖
	if err != nil {
		fmt.Println(err)
	}

	if n == int64(1) {
		fmt.Println("success")
	}

	var imapGet map[string]string

	valueGet, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
	}

	errShal := json.Unmarshal(valueGet, &imapGet)
	if errShal != nil {
		fmt.Println(err)
	}
	fmt.Println(imapGet["username"])
	fmt.Println(imapGet["phonenumber"])
}

func expireTime() {
	conn, err := redis.Dial("tcp", "192.168.1.33:6379")
	if err != nil {
		fmt.Println("connect err:", err)
		return
	}
	defer conn.Close()

	//设置
	_, err = conn.Do("SET", "testExpire", "ddfdff")

	n, _ := conn.Do("EXPIRE", "testExpire", 20)
	if n == int64(1) {
		fmt.Println("success")
	}
}
/*集合操作*/
func pushList() {
	c, err := redis.Dial("tcp", "192.168.1.33:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("lpush", "runoobkey", "redis")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	_, err = c.Do("lpush", "runoobkey", "mongodb")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	_, err = c.Do("lpush", "runoobkey", "mysql")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	values, _ := redis.Values(c.Do("lrange", "runoobkey", "0", "100"))

	for _, v := range values {
		fmt.Println(string(v.([]byte)))
	}
}
