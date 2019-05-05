package model

import . "Backend-for-Android-Project/model/base"

type Lesson struct {
	Lid        	uint    	`gorm:"primary_key"`
	LessonName 	string  	`gorm:"size:255"`
	LessonTime 	string  	`gorm:"size:255"`
	Name		string
}

func (Lesson) TableName() string {
	return "Lesson"
}

type StuLesson struct {
	Sid		uint
	Lid		uint
	SLID	uint	`gorm:"primary_key"`
}

func (StuLesson) TableName() string {
	return "StuLesson"
}

type Attachment struct {
	AID	uint
	FileName	string
	FilePath	string
	UniName		string
	LID			uint
}

func (Attachment) TableName() string {
	return "Attachment"
}

func GetStudentLessons(SID uint) ([]Lesson, bool) {
	var lessons []Lesson
	Db.Table("Lesson").Select("Lesson.*, user.name").Joins("inner join teacher on teacher.tid = Lesson.tid").Joins("inner join user on user.id = teacher.tid").Joins("inner join StuLesson on StuLesson.LID = Lesson.LID").Joins("inner join student on StuLesson.SID = student.SID").Where(Student{Sid:SID}).Find(&lessons)
	if len(lessons) == 0 {
		return nil, false
	}else {
		return lessons, true
	}
}
