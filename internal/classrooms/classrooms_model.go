package classrooms

type Classroom struct {
	ClassroomID   uint   `gorm:"column:classroomid"`
	ClassroomName string `gorm:"column:classroomname"`
}
