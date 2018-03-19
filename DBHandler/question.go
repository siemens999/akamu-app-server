package DBHandler

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

type Question struct {
	Id       uint32
	Author   uint32
	Text     string
	Image    uint32
	Answer   Answer
	subject  uint32
	reviewed bool
	verified bool
}

type Answer struct {
	Id uint32
}

type MultipleChoiceAnswer struct {
	Answer
	Options []Option
}

type Option struct {
	Id      uint32
	Text    string
	Correct bool
}

type TextInputAnswer struct {
	Answer
	Correct string
}

/*
 * Selects questions from the database. Parameters id, subjectid, poolid, reviewed,
 * verified, authorid are filter parameters, that will be applied with sql 'where'.
 * Only the Answer reference will be populated. All other reference ids remain.
 */
func SelectQuestion(id uint32, subjectid uint32, poolid uint32, reviewed bool, verified bool, authorid uint32) ([]Question, error) {

	sqStmt := sq.Select("*").From("question")

	if id != 0 {
		sqStmt = sqStmt.Where("id=?", id)
	} else { // Id is unique. Any other filters irrelevant.
		if subjectid != 0 {
			sqStmt = sqStmt.Where("subject=?", subjectid)
		}
		if poolid != 0 {
			sqStmt = sqStmt.Where("pool=?", poolid)
		}
	}
	if reviewed {
		sqStmt = sqStmt.Where("reviewed=?", reviewed)
	}
	if verified {
		sqStmt = sqStmt.Where("verified=?", verified)
	}
	if authorid != 0 {
		sqStmt = sqStmt.Where("author=?", authorid)
	}

	sql, args, err := sqStmt.ToSql()

	if err != nil {
		return nil, err
	}
	fmt.Println(sql)
	fmt.Println(args)

	return nil, nil
}
