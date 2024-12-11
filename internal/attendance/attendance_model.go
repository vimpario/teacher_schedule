package attendance

type Attendance struct {
	AttendanceID uint `gorm:"column:attendanceid;primaryKey;autoIncrement"`
	StudentID    *uint `gorm:"column:studentid"`
	ScheduleID   *uint `gorm:"column:scheduleid"`
	IsPresent    bool `gorm:"column:ispresent"`
}

type NewAttendance struct {
	StudentID  *uint `gorm:"column:studentid"`
	ScheduleID *uint `gorm:"column:scheduleid"`
	IsPresent  *bool `gorm:"column:ispresent"`
}
