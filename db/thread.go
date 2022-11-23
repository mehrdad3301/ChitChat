package db

import ( 
	"fmt"
	"strconv"
	"html/template"
	md "github.com/shurcooL/github_flavored_markdown"
) 


func (user *User) CreateThread(topic string) (error) {

	queryString := `
	insert into 	
		threads(topic, user_id, created_at)
	values(?, ?, ?)` 

	_, err := db.Exec(queryString, topic, user.Id, getTime())
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
		t := Thread{}
		err = rows.Scan(&t.Id, &t.Topic,
						&t.UserId, &t.CreatedAt) 
		if err != nil {
			return nil, fmt.Errorf("Threads: ", err)
		}
		t.TopicHTML = renderMarkdown(t.Topic)
		threads = append(threads, t)
	}
	return threads, nil
}

func GetThread(id string) (*Thread, error) { 
	query := `
	select 
		* 
	from 
		threads 
	where id=?`

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("Thread: ", err)
	}
	
	var thread Thread 	

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&thread.Id, &thread.Topic,
						&thread.UserId, &thread.CreatedAt) 
		if err != nil {
			return nil, fmt.Errorf("Thread: ", err)
		}
		thread.TopicHTML = renderMarkdown(thread.Topic)
	}
	return &thread, nil
}

func (t *Thread) Posts() ([]Post, error) {

	query := `
	select * 
	from posts 
	where thread_id=?`
	
	rows, err := db.Query(query, t.Id) 
	if err != nil { 
		return nil, fmt.Errorf("Posts: ", err)
	}

	var posts []Post

	defer rows.Close()
	for rows.Next() { 
		var p Post 
		err =rows.Scan(&p.Id, &p.Body, &p.UserId, &p.ThreadId, &p.CreatedAt) 
		if err != nil { 
			return nil, fmt.Errorf("Posts: ", err)
		}
		p.BodyHTML = renderMarkdown(p.Body)
		posts = append(posts, p)
	}
	return posts, nil
}

func (t *Thread) User() (*User) { 

	user, err := GetUser("id", strconv.Itoa(t.UserId))
	if err != nil { 
		fmt.Println("User: ", err)
	} 
	return user 
}


func renderMarkdown(body string) template.HTML { 
	return template.HTML(md.Markdown([]byte(body)))
}
