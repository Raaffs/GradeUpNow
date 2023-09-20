package models
import (
	"database/sql"
	"errors"
)
type Theory struct{
	TQ_id int
	TQ_num int
	TQ_type string
	TQ_question string
	TQ_keywords string
}

func (theory *Model)Get_Theory(q_num int, q_sub string)(*Theory,error){
	ques:=&Theory{}
	stmt:=`SELECT TQ_num, TQ_question,TQ_keywords
	FROM Theory
	WHERE TQ_qum=? and TQ_type=?`
	row:=theory.DB.QueryRow(stmt,q_num,q_sub)
	err:=row.Scan(&ques.TQ_num,&ques.TQ_question,&ques.TQ_keywords)
	if err!=nil{
		if errors.Is(err, sql.ErrNoRows){
			return nil, ErrNoRecord
		}else{
			return nil,err
		}
	}
	return ques,nil
}