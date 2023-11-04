package models

//Corresponds to the columns and their datatype in sql database
type Theory struct{
	TQ_id int
	TQ_num int
	TQ_type string
	TQ_question string
	TQ_keywords string
}

func (theory *Model)Get_Theory(q_sub string)([]*Theory,error){
	stmt:=`SELECT *
	FROM Theory
	WHERE TQ_type=?
	ORDER BY TQ_num ASC`
	rows,err:=theory.DB.Query(stmt,q_sub)
	if err!=nil{
		return nil,err
	}
	defer rows.Close()
	theory_list:=[]*Theory{}
	for rows.Next(){
		t:=&Theory{}
		err=rows.Scan(
			&t.TQ_id,
			&t.TQ_num,
			&t.TQ_type,
			&t.TQ_question,
			&t.TQ_keywords,
		)
		if err!=nil{
			return nil,err
		}
		theory_list=append(theory_list, t)
	}
	return theory_list,nil
}