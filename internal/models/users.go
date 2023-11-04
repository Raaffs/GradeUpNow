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
	Java_score  int
    DBMS_score  int
    DSA_score   int
    FSE_score   int
}
//inserts a new user in database and initializes the score to 0
func (user *Model)SignUp(username string,email_id string,password string)(int,error){
	//? works as a place holder for the data we want to insert
	stmt:=`INSERT INTO User_profile(username,email_id,password,total_score,mcq_score,theory_score,Java_score,DBMS_score,DSA_score,FSE_score)
	VALUES(?,?,?,0,0,0,0,0,0,0)`
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
func (user *Model)Check_if_exist(username string,email_id string)(bool,error){
	stmt:="SELECT COUNT(*) FROM User_profile WHERE Username = ?"
    row := user.DB.QueryRow(stmt, username)
    var count int
    err := row.Scan(&count)
    if err != nil {
        return false, err
    }
    return count > 0, nil
}

func (user *Model)Get(username string)(*User, error){
	
	user_query:=&User{}
	stmt:=`SELECT username,email_id,total_score,mcq_score,theory_score,Java_score,DBMS_score,DSA_score,FSE_score,password
	FROM User_profile
	WHERE username=?`

	row:=user.DB.QueryRow(stmt,username)

	err := row.Scan(&user_query.Username,&user_query.Email_id,&user_query.Total_score,&user_query.Mcq_score,&user_query.Theory_score,&user_query.Java_score,&user_query.DBMS_score,&user_query.DSA_score,&user_query.FSE_score,&user_query.Password)
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
	ORDER BY total_score DESC`
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

func (user *Model) Update_score(subject string, score int) error {

	current_score,err:=user.Get(G_CurrentUserSession)
	if err!=nil{
		if errors.Is(err,ErrNoRecord){
			return ErrNoRecord
		}
		return err
	}
	

    switch subject {
    case "java":
		
		current_score.Java_score+=score
		
		stmt := `UPDATE User_profile 
		SET Java_score=? 
		WHERE username=?`
    	_, err = user.DB.Exec(stmt,current_score.Java_score,G_CurrentUserSession)
   		if err != nil {
       	 	return err
    	}
    case "dbms":
		current_score.DBMS_score+=score
		stmt := `UPDATE User_profile 
		SET DBMS_score=? 
		WHERE username=?`
    	_, err := user.DB.Exec(stmt, current_score.DBMS_score,G_CurrentUserSession)

   		if err != nil {
       	 	return err
    	}
    case "dsa":
		current_score.DSA_score+=score
		stmt := `UPDATE User_profile 
		SET DSA_score=? 
		WHERE username=?`
    	_, err := user.DB.Exec(stmt, current_score.DSA_score,G_CurrentUserSession)
   		if err != nil {
       	 	return err
    	}
    case "fse":
		current_score.FSE_score+=score
		stmt := `UPDATE User_profile 
		SET FSE_score=? 
		WHERE username=?`
    	_, err := user.DB.Exec(stmt, current_score.FSE_score,G_CurrentUserSession)
   		if err != nil {
       	 	return err
    	}
    default:
        return errors.New("Invalid subject")
    }
	current_score.Total_score=current_score.Java_score+current_score.DBMS_score+current_score.FSE_score+current_score.DSA_score
	stmt:=`UPDATE User_profile
	SET total_score=?
	WHERE username=?`
	_,err=user.DB.Exec(stmt,current_score.Total_score,G_CurrentUserSession)
	if err!=nil{
		fmt.Println("total score:" ,current_score.Total_score,"dbms score:",current_score.DBMS_score)
		return err
	}
	return nil
}
    

