package grades

type Grade struct {
	GradeID   uint   `gorm:"column:gradeid"`
	StudentID uint   `gorm:"studentid"`
	SubjectID uint   `gorm:"subjectid"`
	Grade     uint   `gorm:"column:grade"`
	GradeDate string `gorm:"column:gradedate"`
}
