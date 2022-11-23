package db

import ( 
	"fmt"
	"time"
	"strconv"
) 


func CreateThread(topic string, userId int) (error) {

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

func (thread *Thread) NumReplies() (int) {

	queryString := `
	select 
		count(*) 
	from 
		posts
	where thread_id = ?`

	rows, err := db.Query(queryString, thread.Id)
	if err != nil {
		fmt.Println("NumReplies: ", err)
		return 0
	}

	var count int 

	defer rows.Close()
	if rows.Next() {
		if err = rows.Scan(&count); err != nil {
			fmt.Println("NumReplies: ", err)
			return 0
		}
	}
	return count 
}


func GetThreads() ([]Thread, error) {
	query := `
	select 
		* 
	from 
		threads 
	order by created_at DESC`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Threads: ", err)
	}

	threads := []Thread{}
	
	defer rows.Close()
	for rows.Next() {
		conv := Thread{}
		err = rows.Scan(&conv.Id, &conv.Topic,
						&conv.UserId, &conv.CreatedAt) 
		if err != nil {
			return nil, fmt.Errorf("Threads: ", err)
		}
		threads = append(threads, conv)
	}
	return threads, nil
}

func (t Thread) User() (*User) { 

	user, err := GetUser("id", strconv.Itoa(t.UserId))
	if err != nil { 
		fmt.Println("User: ", err)
	} 
	return user 
}

