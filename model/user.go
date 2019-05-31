package model

import (
	. "Backend-for-Android-Project/model/base"
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

type User struct {
	ID       	uint	`gorm:"primary_key"`
	Name		string	`gorm:"size:255"`
	Password	string	`gorm:"size:255"`
}

func (User) TableName() string {
	return "user"
}

type Teacher struct {
	User			User	`gorm:"ForeignKey:TID"`
	Tid 			uint	`gorm:"primary_key"`
	Introduction	string	`gorm:"size:500"`
}

func (Teacher) TableName() string {
	return "teacher"
}

type Student struct {
	User			User	`gorm:"ForeignKey:SID"`
	Sid 			uint	`gorm:"primary_key"`
	GraduatedSchool	string	`gorm:"size:255"`
}

func (Student) TableName() string {
	return "student"
}

func SignIn(id uint, password string) (string,string, bool) {
	var user User
	Db.Where(User{ID:id, Password:password}).First(&user)
	if user.Name == "" || password == "" {
		return "","", false
	}
	return getBaseInfo(id)
}

func getBaseInfo(id uint) (string, string, bool) {
	var result User
	Db.Table("teacher").Select("user.name").Joins("inner join user on user.id = teacher.TID").Where("teacher.TID = ?", id).Scan(&result)
	if result.Name != "" {
		return result.Name,"teacher", true
	}
	Db.Table("student").Select("user.name").Joins("inner join user on user.id = student.SID").Where("student.SID = ?", id).Scan(&result)
	if result.Name != "" {
		return result.Name,"student", true
	}
	return "","", false
}

func CreatToken(id uint) string {
	cruTime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(cruTime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	//var err error
	//_, err = MRedis.Do("set", token, id) //设置3天过期
	TokenID[token] = id
	//if err != nil{
	//	fmt.Println("token create error:",err.Error())
	//}
	return token
}

func Token2ID(token string) uint {
	id, ok := TokenID[token]
	//id, err := redis.Uint64(MRedis.Do("get", token))
	if ok{
		return id
	} else {
		//fmt.Println("Token2ID error:",err.Error())
		return 0
	}
}

func SignToken(token string) (string, string, bool) {
	id := Token2ID(token)
	fmt.Println(id)
	if id == 0{
		return "", "", false
	} else {
		return getBaseInfo(id)
	}
}
