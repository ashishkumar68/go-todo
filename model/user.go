package model

type User struct {
	Identity
	CreatedMetaInfo
	UpdatedMetaInfo
	DeletedMetaInfo
	VersionMetaInfo

	FirstName string `gorm:"name:first_name;type:varchar(255);not null;" json:"firstName"`
	LastName string `gorm:"name:last_name;type:varchar(255);not null;" json:"lastName"`
	Email string `gorm:"name:email;type:varchar(255);unique;not null" json:"email"`
	Phone string `gorm:"name:phone;type:varchar(32);" json:"phone"`
}