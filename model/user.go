package model

import (
	. "Backend-for-Android-Project/model/base"
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
