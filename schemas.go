package main

//import "github.com/golang/protobuf/protoc-gen-go/plugin"
import(
	"time"
)

type SignUpFormular struct {

	Username string        `json:"username" binding:"required"`
	Password string    `json:"password" binding:"required"`
    Email string       `json:"email" binding:"required"`
    University string  `json:"university" binding:"required"`
    Semester int   `json:"semester" binding:"required"`
}
type UserInfoSchema struct {
	Id     uint32 `json:"id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Avatar uint32 `json:"avatar" binding:"required"`
	Title  string `json:"title" binding:"required"`
}

type UserSchema struct {
	UserInfoSchema
	TimeRegistered string `json:"time-registered" binding:"required"` // datetime format
	Semester       int    `json:"semester" binding:"required"`
	Verified       bool   `json:"verified" binding:"required"`
	University     string `json:"university" binding:"required"`
	Experience     int    `json:"experience" binding:"required"`
	Memorycoins    int    `json:"memorycoins" binding:"required"`
}

/*
 * Authentication token
 */
type AuthToken struct{
    Value string `json:"value" binding:"required"`
    Expiriation time.Time `json:"expiration" binding:"required"`
}

type TitleSchema struct {
	Name         string `json:"name" binding:"required"`
	Subject      uint32 `json:"subject" binding:"required"`
	Unlock_Score int    `json:"unlock_score" binding:"required"`
	Unlock_Win   int    `json:"unlock_win" binding:"required"`
}

type PoolSchema struct {
	Id    uint32 `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Code  string `json:"code" binding:"required"`
	Image uint32 `json:"image" binding:"required"`
}

type ExplanationSchema struct {
	Text string `json:"text" binding:"required"`
}

type RoundSchema struct {
	Pool      PoolSchema `json:"pool" binding:"required"`
	Questions []struct {
		Question       QuestionSchema `json:"question" binding:"required"`
		AnswerOpponent AnswerSchema   `json:"answer-opponent"`
	} `json:"questions" binding:"required"`
}

type QuestionSchema struct {
	Id      uint32       `json:"id" binding:"required"`
	Authro  uint32       `json:"author"`
	Text    string       `json:"text" binding:"required"`
	Image   uint32       `json:"image"`
	Subject string       `json:"subject" binding:"required"`
	Pool    PoolSchema   `json:"pool" binding:"required"`
	Answer  AnswerSchema `json:"answer" binding:"required"`
	//should be either multiplechoice answer or textinputanswer.
}

type DuelInfoSchema struct {
	Id              uint32         `json:"id" binding:"required"`
	UserChallanger  UserInfoSchema `json:"user-challanger" binding:"required"`
	UserChallanged  UserInfoSchema `json:"user-challanged" binding:"required"`
	Status          int            `json:"status" required:"true"`
	TimeStart       string         `json:"time-start" required:"true"`
	TimeChanged     string         `json:"time-changed" required:"true"`
	TimeEnd         string         `json:"time-end"`
	ScoreChallenger UserInfoSchema `json:"score-challenger"`
	ScoreChallenged UserInfoSchema `json:"score-challenged"`
	Winner          uint32         `json:"winner"`
}

type DuelSchema struct {
	DuelInfoSchema
	Round1 RoundSchema `json:"round1" required:"true"`
	Round2 RoundSchema `json:"round2"`
}

type MultipleChoiceAnswerItemSchema struct {
	Id      uint32 `json:"id" binding:"required"`
	Correct bool   `json:"name" binding:"required"`
	Image   uint32 `json:"image"`
	Text    string `json:"text" binding:"required"`
}

type TextInputAnswerSchema string

type AnswerSchema struct {
	// TODO: implement this!
}
