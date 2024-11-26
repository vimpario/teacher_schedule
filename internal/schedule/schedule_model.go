package schedule

type Schedule struct {
	ID         uint  `gorm:"primaryKey"`
	TeacherID  uint  `gorm:"not null"`
	DayID      uint  `gorm:"not null"`
	SlotID     uint  `gorm:"not null"`
	GroupID    *uint `gorm:""`
	SubjectID  *uint `gorm:""`
	IsOccupied bool  `gorm:"default:false"`
}
