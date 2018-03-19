package main

type UserInfoSchema struct {
	Id uint32;		`json:"id" binding:"required"`
	Name string;	`json:"name" binding:"required"`
	Avatar uint32;	`json:"avatar" binding:"required"`
	Title string;	`json:"title" binding:"required"`
}

type UserSchema struct extends UserInfoSchema {
	*UserInfoSchema
	TimeRegistered string; 	`json:"time-registered" binding:"required"`	// datetime format
	Semester int;			`json:"semester" binding:"required"`
	Verified bool;			`json:"verified" binding:"required"`
	University string;		`json:"university" binding:"required"`
	Experience int;			`json:"experience" binding:"required"`
	Memorycoins int;		`json:"memorycoins" binding:"required"`
}

type RoundSchema struct {
	Pool PoolSchema			`json:"pool" binding:"required"`
	Questions []struct{
		Question Question	`json:"question" binding:"required"`
		AnswerOpponent		`json:"answer-opponent"`
	}	`json:"questions" binding:"required"`
}

type QuestionSchema struct {
	Id uint32				`json:"id" binding:"required"`
	Authro uint32			`json:"author"`
	Text string				`json:"text" binding:"required"`
	Image uint32			`json:"image"`
	Subject string			`json:"subject" binding:"required"`
	Pool Pool				`json:"pool" binding:"required"`
	Answer Answer			`json:"answer" binding:"required"`
	//should be either multiplechoice answer or textinputanswer.
}