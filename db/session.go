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
