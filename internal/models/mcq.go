package models
import (
	"database/sql"
	"errors"
)

//Corresponds to the columns and their datatype in sql database except MQ_opt
type Mcq struct{
	MQ_id int
	MQ_num int
	MQ_type string
	MQ_question string
	MQ_ans int
	MQ_opt Option
}
type Option struct{
	Opt1 int
	Opt2 int 
	Opt3 int
	Opt4 int
}
func (theory *Model)Get_Mcq(q_num int, q_sub string)(*Mcq,error){
	ques:=&Mcq{}
	stmt:=`SELECT MQ_num, MQ_question,MQ_ans
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