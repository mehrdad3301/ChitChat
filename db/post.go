package db

import (
	"strconv"
	"fmt"
) 



func (p Post) User() (*User) { 

	user, err := GetUser("id", strconv.Itoa(p.UserId))
	if err != nil { 
		fmt.Println("User: ", err)
	} 
	return user 
}

func CreatePost(threadId, userId int, body string) (error) { 

	query := `
	insert into 
		posts (body, user_id, thread_id, created_at) 
	values(?, ?, ?, ?)`

	
	_, err := db.Exec(query, body, userId, threadId, getTime())
	if err != nil { 
		return fmt.Errorf("CreatePost: ", err)
	}	
	return nil 
}
