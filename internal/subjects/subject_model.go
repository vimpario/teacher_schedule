package subjects

type Subjects struct {
	SubjectID   uint   `gorm:"primaryKey;column:subjectid"`
	SubjectName string `gorm:"column:subjectname;unique;not null"`
}
