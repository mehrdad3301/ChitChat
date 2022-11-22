package db 

import ( 
	"time"
	"fmt"
	UUID "github.com/google/uuid"
)


func CreateSession(user *User) (string, error) { 

	queryString := `
	insert into 
		sessions(uuid, email, user_id, created_at) 
	values(?, ?, ?, ?)`
	

	uuid := UUID.New().String()
	_, err := db.Exec(queryString, uuid, user.Email, 
			user.Id, time.Now().Format(time.UnixDate))

	if err != nil { 
		return "", fmt.Errorf("CreateSession: %v", err) 
	}
	return uuid, nil 
} 

func ValidSession(uuid string) (bool, error) { 

	queryString:=`
	select 
		*
	from 
		sessions
	where uuid=?`

	rows, err := db.Query(queryString, uuid)
	if err != nil { 
		return false, fmt.Errorf("ValidSessoin: ", err)
	}
	
	defer rows.Close() 
	if rows.Next() { 
		return true, nil
	}
	return false, nil
}

func DeleteSession(uuid string) (error) { 
	
	queryString :=`
	delete from 
		sessions
	where uuid=?`

	_, err := db.Exec(queryString, uuid)
	if err != nil { 
		return err 
	}
	return nil 
}
