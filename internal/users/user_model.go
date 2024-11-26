package users

type User struct {
	ID         uint   `gorm:"primaryKey"`
	FirstName  string `grom:"size:50;not null"`
	LastName   string `grom:"size:50;not null"`
	Patronymic string `grom:"size:50"`
	BirthDate  string `gorm:"not null"`
	RoleID     uint   `gorm:"not null"`
}
