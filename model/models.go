package model

type Tokens struct {
	Id           uint   `jeson:"id"`
	UserId       uint   `jeson:"userId"`
	AccessToken  string `jeson:"accessToken"`
	RefreshToken string `jeson:"refreshToken"`
	Exp          int64  `jeson:"exp"`
}

type User struct {
	Id    uint   `jeson:"id"`
	Login string `jeson:"login"`
	Name  string `jeson:"name"`
	Pass  string `jeson:"pass"`
}

type Task struct {
	Id          uint   `jeson:"id"`
	UserId      uint   `jeson:"userId"`
	Title       string `jeson:"title"`
	Description string `jeson:"description"`
	Completed   bool   `jeson:"completed" sql:"DEFAULT:false;type:boolean;index" gorm:"not null"`
	StartDate   int64  `jeson:"startDate"`
	EndDate     int64  `jeson:"endDate"`
}
