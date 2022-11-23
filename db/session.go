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
			user.Id, time.Now().Format(time.ANSIC))

	if err != nil { 
		return "", fmt.Errorf("CreateSession: %v", err) 
	}
	return uuid, nil 
} 

func GetSession(uuid string) (*Session, error) {

	queryString:=`
	select 
		*
	from 
		sessions
	where uuid=?`

	rows, err := db.Query(queryString, uuid)
	if err != nil { 
		return nil, fmt.Errorf("GetSession: ", err)
	}
	
	s := Session{}

	defer rows.Close() 
	if rows.Next() { 
		rows.Scan(&s.Id, &s.UUID, &s.Email, &s.UserId, &s.CreatedAt)	
		fmt.Println("_________", s)
		return &s, nil
	} 
	
	return nil, fmt.Errorf("GetSession: No rows")
}

func ValidSession(uuid string) (bool, error) { 

	_, err := GetSession(uuid)
	if err != nil { 
		return false, fmt.Errorf("ValidSession: ", err)
	}
	return true, nil 
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
