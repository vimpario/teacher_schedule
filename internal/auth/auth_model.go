package auth

type Auth struct{
	AuthID uint `gorm:"column:authid"`
	UserID uint `gorm:"column:userid"`
	Login string `gorm:"column:login;size:50"`
	PasswordHash string `gorm:"column:passwordhash;size:225"`
	CreatedAt string `gorm:"column:createdat"`
	LastLogin string `gorm:"column:lastlogin"`
}