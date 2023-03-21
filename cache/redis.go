package cache

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hjk-cloud/douyin/cache/redisPool"
	"strconv"
)

type userInfo struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Password int    `db:"password"`
}

func GetAll() {
	//从连接池当中获取链接
	conn := redisPool.Pool.Get()
	//先查看redis中是否有数据
	//conn,_ :=redis.Dial("tcp","localhost:6379")
	defer conn.Close()
	values, _ := redis.Values(conn.Do("lrange", "mlist", 0, -1))

	if len(values) > 0 {
		//如果有数据
		fmt.Println("从redis获取数据")
		//从redis中直接获取
		for _, key := range values {
			pid := string(key.([]byte))
			id, _ := strconv.Atoi(pid)
			results, _ := redis.Bytes(conn.Do("GET", id))
			var u userInfo
			err := json.Unmarshal(results, &u)
			if err != nil {
				fmt.Println("json 反序列化出错")
			} else {
				fmt.Printf("id = %d\n", u.Id)
				fmt.Printf("name = %s\n", u.Name)
				fmt.Printf("password = %d\n", u.Password)
			}
		}
	} else {
		fmt.Println("从mysql中获取")

		//查询数据库
		db, _ := sql.Open("mysql", "root:541688@tcp(localhost:3306)/douyin")
		defer db.Close()

		var userInfos []userInfo

		rows, _ := db.Query("select id, name, password from user")
		for rows.Next() {
			var id int
			var name string
			var password int
			rows.Scan(&id, &name, &password)
			per := userInfo{id, name, password}
			userInfos = append(userInfos, per)

		}
		//写入到redis中:将userinfo以hash的方式写入到redis中
		for _, v := range userInfos {

			v_byte, _ := json.Marshal(v)
			_, err1 := conn.Do("SETNX", v.Id, v_byte)
			_, err2 := conn.Do("rpush", "mlist", v.Id) //rpush从右侧输入，lpush从左侧输入
			// 设置过期时间
			conn.Do("EXPIRE", v.Id, 60*5)
			if err1 != nil || err2 != nil {
				fmt.Println("写入失败")
			} else {
				fmt.Println("写入成功")
			}
		}
		conn.Do("EXPIRE", "mlist", 60*5)
	}
}
