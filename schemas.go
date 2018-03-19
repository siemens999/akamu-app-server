package main

//import "github.com/golang/protobuf/protoc-gen-go/plugin"

type SignUpFormular struct {
	Name string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
    Email string `json:"email" binding:"required"`
    University string `json:"university" binding:"required"`
    Semester string `json:"semester" binding:"required"`
}
type UserInfoSchema struct {
	Id uint32		`json:"id" binding:"required"`
	Name string	`json:"name" binding:"required"`
	Avatar uint32	`json:"avatar" binding:"required"`
	Title string	`json:"title" binding:"required"`
}

type UserSchema struct {
	*UserInfoSchema
	TimeRegistered string 	`json:"time-registered" binding:"required"`	// datetime format
	Semester int			`json:"semester" binding:"required"`
	Verified bool			`json:"verified" binding:"required"`
	University string		`json:"university" binding:"required"`
	Experience int			`json:"experience" binding:"required"`
	Memorycoins int		`json:"memorycoins" binding:"required"`
}

type TitleSchema struct {
	Name string 		`json:"name" binding:"required"`
	Subject uint32		`json:"subject" binding:"required"`
	Unlock_Score int	`json:"unlock_score" binding:"required"`
	Unlock_Win int		`json:"unlock_win" binding:"required"`
}

type PoolSchema struct{
	Id uint32			`json:"id" binding:"required"`
	Name string			`json:"name" binding:"required"`
	Code string			`json:"code" binding:"required"`
	Image uint32		`json:"image" binding:"required"`
}

type ExplanationSchema struct {
	Text string `json:"text" binding:"required"`
}

type MultipleChoiceAnswerItemSchema struct {
	Id uint32 			`json:"id" binding:"required"`
	Correct bool 		`json:"name" binding:"required"`
	Image uint32 		`json:"image"`
	Text string 		`json:"text" binding:"required"`
}

type TextInputAnswer string

// type Answer struct {} // TODO: implement this!