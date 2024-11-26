package subjects

type Subjects struct {
	SubjectID   uint   `gorm:"column:subjectid"`
	SubjectName string `gor:"column:subjectname"`
}
