package db

import ( 
	"fmt"
	"time"
) 


func CreateThread(topic, userId string) (error) {

	queryString := `
	insert into 	
		threads(topic, user_id, created_at)
	values(?, ?, ?)` 

	_, err := db.Exec(queryString, topic, userId, 
					time.Now().Format(time.UnixDate))
	if err != nil { 
		return fmt.Errorf("CreateThread: ", err)
	}

	return nil
}
