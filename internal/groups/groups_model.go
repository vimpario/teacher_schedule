package groups

type Group struct {
	GroupID      uint   `gorm:"column:groupid"`
	GroupNumber  string `gorm:"column:groupnumber;size:20"`
	CourseNumber uint   `gorm:"column:coursenumber"`
}
