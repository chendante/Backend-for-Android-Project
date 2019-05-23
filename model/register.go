package model

import (
	. "Backend-for-Android-Project/model/base"
	"time"
)

type Register struct {
	Rid				uint	`gorm:"primary_key"`
	Lid				uint
	BeginTime		time.Time
	//LessonName		string
	DeleteStatus	int
}

type RegisterInfo struct {
	Rid				uint
	Lid				uint
	BeginTime		time.Time
	LessonName		string
	DeleteStatus	int
}

func (Register) TableName() string {
	return "Register"
}

type AttendanceBook struct {
	Adid			uint		`gorm:"primary_key"`
	Rid				uint
	Sid				uint
	AnswerTime		time.Time
}

func (AttendanceBook) TableName() string {
	return "AttendanceBook"
}

type AttendanceInfo struct {
	Adid	uint
	Rid		uint
	Sid		uint
	AnswerTime	time.Time
	Name	string
}

func CreateRegister(lid uint) int {
	register := Register{Lid:lid, BeginTime:time.Now(), DeleteStatus: 1}
	Db.Create(&register)
	var res Register
	Db.Where(&Register{Lid:lid, DeleteStatus:1}).Last(&res)
	return int(res.Rid)
}

func DeleteRegister(rid uint) {
	var register Register
	Db.First(&register, rid)
	register.DeleteStatus = 0
	Db.Save(&register)
}

func StuSelectRegister(sid uint) (RegisterInfo, bool) {
	var register RegisterInfo
	var ok bool
	Db.Table("StuLesson").Select("Lesson.lesson_name, Register.*").Joins("inner join Lesson").Joins("inner join Register").Where("Register.delete_status = ? AND StuLesson.sid = ?", 1, sid).Last(&register)
	if register.Rid == 0{
		ok = false
	} else {
		var attendanceBook AttendanceBook
		Db.Where(&AttendanceBook{Rid:register.Rid}).First(&attendanceBook)
		if attendanceBook.Adid == 0{
			ok = true
		} else {
			ok = false
		}
	}
	return register, ok
}

func PostAttendance(sid, rid uint) bool {
	attendanceBook := AttendanceBook{Rid:rid, Sid:sid, AnswerTime:time.Now()}
	Db.Create(&attendanceBook)
	return !Db.NewRecord(attendanceBook)
}

func SelectAttendance(rid uint) []AttendanceInfo {
	var attendances []AttendanceInfo
	//Db.Where(&AttendanceBook{Rid:rid}).Find(&attendances)
	Db.Table("student").Select("AttendanceBook.*, user.name").Joins("inner Join user on student.sid = user.id").Joins("inner Join StuLesson on StuLesson.sid = student.sid").Joins("inner Join Lesson on Lesson.lid = StuLesson.lid").Joins("inner join Register on Register.lid = Lesson.lid").Joins("left join AttendanceBook on AttendanceBook.rid = Register.rid and AttendanceBook.sid = student.sid").Where("Register.rid = ?", rid).Find(&attendances)
	return attendances
}