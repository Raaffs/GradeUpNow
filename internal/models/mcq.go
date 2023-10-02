package models
import (
	//"database/sql"
	//"errors"
)

//Corresponds to the columns and their datatype in sql database except MQ_opt
type Mcq struct{
	MQ_id int
	MQ_num int
	MQ_type string
	MQ_question string
	MQ_ans int
	Option1     string // Separate columns for options
	Option2     string
	Option3     string
	Option4     string
	Options		[]string
	UserChoice int
}

func (mcq *Model)Get_Mcq(q_sub string)([]*Mcq,error){
	//ques:=&Mcq{}
	stmt:=`SELECT MQ_question, Option1, Option2, Option3, Option4, MQ_ans
	FROM Mcq
	WHERE MQ_type=?
	ORDER BY MQ_num ASC`
	rows,err:=mcq.DB.Query(stmt,q_sub)
	if err!=nil{
		return nil,err
	}
	defer rows.Close()
	
	mcq_list:=[]*Mcq{}
	
	for rows.Next(){
		m:=&Mcq{}
		err=rows.Scan(&m.MQ_question, &m.Option1, &m.Option2, &m.Option3, &m.Option4, &m.MQ_ans)
		if err!=nil{
			return nil,err
		}
		m.Options=[]string{m.Option1,m.Option2,m.Option3,m.Option4}
		mcq_list=append(mcq_list, m)
	}
	return mcq_list,nil
}
