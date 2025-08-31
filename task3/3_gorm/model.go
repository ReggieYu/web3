package main

type User struct {
	Id        uint   `gorm:"primarykey"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Posts     []Post `gorm:"foreignKey:UserId"`
	PostCount uint   `gorm:"default:0"`
}

type Post struct {
	Id            uint      `gorm:"primarykey"`
	Title         string    `gorm:"not null"`
	Content       string    `gorm:"type:text"`
	UserId        uint      `gorm:"not null"`
	Comments      []Comment `gorm:"foreignKey:PostId"`
	CommentStatus string    `gorm:"default: ''"`
}

type Comment struct {
	Id      uint   `gorm:"primarykey"`
	Content string `gorm:"type:text"`
	PostId  uint   `gorm:"not null"`
}
