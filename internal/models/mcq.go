package models
import (
	"database/sql"
	"errors"
)
type Mcq struct{
	MQ_id int
	MQ_num int
	MQ_type string
	MQ_question string
	MQ_ans int
}

func (theory *Model)Get_Mql(q_num int, q_sub string)(*Mcq,error){
	ques:=&Mcq{}
	stmt:=`SELECT TQ_num, TQ_question,TQ_keywords
	FROM Theory
	WHERE MQ_qum=? and MQ_type=?`
	row:=theory.DB.QueryRow(stmt,q_num,q_sub)
	err:=row.Scan(&ques.MQ_num,&ques.MQ_question,&ques.MQ_ans)
	if err!=nil{
		if errors.Is(err, sql.ErrNoRows){
			return nil, ErrNoRecord
		}else{
			return nil,err
		}
	}
	return ques,nil
}