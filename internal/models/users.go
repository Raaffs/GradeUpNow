package models

import (
	"database/sql"
	"errors"
	"fmt"
)
type Model struct{
	DB *sql.DB
}
//Corresponds to the columns and their datatype in sql database
type User struct{
	User_id int			//gets auto incremented in the database
	Username string		//NOT NULL UNIQUE
	Email_id string		//NOT NULL UNIQUE
	Password string		//NOT NULL
	Total_score int
	Mcq_score int
	Theory_score int
}
//inserts a new user in database and initializes the score to 0
func (user *Model)SignUp(username string,email_id string,password string)(int,error){
	//? works as a place holder for the data we want to insert
	stmt:=`INSERT INTO User_profile(username,email_id,password,total_score,mcq_score,theory_score)
	VALUES(?,?,?,0,0,0)`
	result,err:=user.DB.Exec(stmt,username,email_id,password)
	if err!=nil{
		return 0,err
	}
	// Use the LastInsertId() method on the result to get the ID of our
	// newly inserted record in the snippets table.
	id,err:=result.LastInsertId()
	if err!=nil{ 
		return 0,err
	}
	// The ID returned has the type int64, so we convert it to an int type
	// before returning.
	return int(id),nil
}
func (user *Model)Get(username string)(*User, error){
	user_query:=&User{}
	fmt.Println(username)
	stmt:=`Select username,total_score,mcq_score,theory_score,password
	FROM User_profile
	WHERE username=?`
	row:=user.DB.QueryRow(stmt,username)
	err := row.Scan(&user_query.Username,&user_query.Total_score, &user_query.Mcq_score, &user_query.Theory_score,&user_query.Password)
	if err!=nil{
		if errors.Is(err, sql.ErrNoRows){
			return nil, ErrNoRecord
		}else{
			return nil,err
		}
	}
	return user_query,nil
}

func(user *Model)Leader_board()([] *User,error){
	stmt:=`SELECT username,mcq_score,theory_score,total_score
	FROM User_profile
	ORDER BY total_score DESC
	LIMIT 10`
	rows,err:=user.DB.Query(stmt)
	if err!=nil{
		return nil,err
	}

	// We defer rows.Close() to ensure the sql.Rows resultset is
	// always properly closed before the Latest() method returns. This defer
	// statement should come *after* you check for an error from the Query()
	// method. Otherwise, if Query() returns an error, you'll get a panic
	// trying to close a nil resultset.
	defer rows.Close()
	leaderboard:=[]*User{}

	for rows.Next(){
		u:=&User{}

		err=rows.Scan(&u.Username,&u.Mcq_score,&u.Theory_score,&u.Total_score)
	
		if err!=nil{
			return nil,err
		}
		leaderboard=append(leaderboard, u)
	}
	return leaderboard,nil
}