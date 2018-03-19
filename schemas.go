package main

type UserInfoSchema struct {
	Id uint32;		`json:"id" binding:"required"`
	Name string;	`json:"name" binding:"required"`
	Avatar uint32;	`json:"avatar" binding:"required"`
	Title string;	`json:"title" binding:"required"`
}

type UserSchema struct extends UserInfoSchema {
	TimeRegistered string; 	`json:"time-registered" binding:"required"`	// datetime format
	Semester int;			`json:"semester" binding:"required"`
	verified bool;			`json:"verified" binding:"required"`
	university string;		`json:"university" binding:"required"`
	experience int;			`json:"experience" binding:"required"`
	memorycoins int;		`json:"memorycoins" binding:"required"`
}