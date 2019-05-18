package base

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"io/ioutil"
)

var Db *gorm.DB
var MRedis redis.Conn
var BaseUrl string
var UploadUri	string

func init() {
	var err error
	bytes, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
	}
	var dataLoaded map[string]string
	if err := json.Unmarshal(bytes, &dataLoaded); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
	}
	Db, err = gorm.Open("mysql", dataLoaded["dataSourceName"])
	if err != nil {
		fmt.Print(err.Error())
	}

	MRedis, err = redis.Dial("tcp", dataLoaded["address"])
	if err != nil {
		fmt.Println("Connect to redis error", err)
	}

	BaseUrl = dataLoaded["BaseUrl"]
	UploadUri = dataLoaded["UploadUri"]
}