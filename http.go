package main

import (
	"fmt"
	mysqlClinet "golottery/mysql"
	"log"
	"net/http"
)

func Start()  {
	http.HandleFunc("/draw", draw);
	http.ListenAndServe("127.0.0.1:8080", nil);
}

func draw(rsp http.ResponseWriter, req *http.Request)  {
	log.Println("有一次返问");
	//redisClient.Set("testgo", "123", 100);
	data := make(map[string]interface{})
	data["name"] = "zwl";
	data["age"] = 12;

	id, err := mysqlClinet.Save("draw", data)

	if (err != nil) {
		log.Println("Save error", err);
	}

	//defer redisClient.Close()
	//rsp.Write([]byte("欢迎参加抽奖"))
	msg := fmt.Sprintf("写入数据库的id是:%d", id);
	log.Println(msg);

	rsp.Write([]byte(msg))
}