package attendance

type Attendance struct{
	AttendanceID uint `gorm:"column:attendanceid"`
	StudentID uint `gorm:"column:studentid"`
	ScheduleID uint `gorm:"column:scheduleid"`
	IsPresent uint `gorm:"column:ispresent"`
}