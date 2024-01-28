package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string     `gorm:"column:first_name;not null"`
	LastName  string     `gorm:"column:last_name;not null"`
	Email     string     `gorm:"column:email;not null;unique"`
	BlogPosts []BlogPost `gorm:"foreignKey:UserID;references:ID"`
}

type BlogPost struct {
	gorm.Model
	Title   string `gorm:"column:title;not null;unique"`
	Tags    string `gorm:"column:tags;not null"`
	Content string `gorm:"column:content;not null"`
	UserID  uint   `gorm:"column:user_id;not null;"`
}
