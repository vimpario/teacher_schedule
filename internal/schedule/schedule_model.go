package schedule

type Schedule struct {
	ID         uint  `gorm:"column:scheduleid"`
	TeacherID  uint  `gorm:"column:teacherid;not null"`
	DayID      uint  `gorm:"column:dayid;not null"`
	SlotID     uint  `gorm:"column:slotid;not null"`
	GroupID    *uint `gorm:"column:groupid"`
	SubjectID  *uint `gorm:"column:subjectid"`
	IsOccupied bool  `gorm:"column:isoccupied;default:false"`
	ClassroomID uint `gorm:"column:classroomid"`
}
