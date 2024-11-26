package users

type User struct {
	ID         uint   `gorm:"column:userid"`
	FirstName  string `grom:"column:firstname;size:50;not null"`
	LastName   string `grom:"column:lastname;size:50;not null"`
	Patronymic string `grom:"column:patronymic;size:50"`
	BirthDate  string `gorm:"column:birthdate;not null"`
	RoleID     uint   `gorm:"column:roleid;not null"`
}
