package models

import (
	"database/sql"
	"errors"
	"fmt"
)
type Model struct{
	DB *sql.DB
}
type User struct{
	User_id int
	Username string
	Email_id string
	Password string
	Total_score int
	Mcq_score int
	Theory_score int
}

func (user *Model)SignIn(username string,email_id string,password string)(int,error){
	stmt:=`INSERT INTO User_profile(username,email_id,password,total_score,mcq_score,theory_score)
	VALUES(?,?,?,78,82,160)`
	result,err:=user.DB.Exec(stmt,username,email_id,password)
	if err!=nil{
		return 0,err
	}
	id,err:=result.LastInsertId()
	if err!=nil{ 
		return 0,err
	}
	return int(id),nil
}
func (user *Model)Get(username string)(*User, error){
	user_query:=&User{}
	fmt.Println(username)
	stmt:=`Select username,total_score,mcq_score,theory_score
	FROM User_profile
	WHERE username=?`
	row:=user.DB.QueryRow(stmt,username)
	err := row.Scan(&user_query.Username,&user_query.Total_score, &user_query.Mcq_score, &user_query.Theory_score)
	if err!=nil{
		if errors.Is(err, sql.ErrNoRows){
			return nil, ErrNoRecord
		}else{
			return nil,err
		}
	}
	return user_query,nil
}
//func(user *UserModel)Leader_board()([] *User,error){}