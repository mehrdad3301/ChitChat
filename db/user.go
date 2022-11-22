package db 

import ( 
	"fmt"
	"time"
	"crypto/sha1"
)

func CreateUser(name, password, email string) (error) {  
	
	queryString := `
	insert into 
		users(name, email, password, created_at)
	values(?, ?, ?, ?)`

	_, err := db.Exec(queryString, name, email, 
	encryptPassword(password), time.Now().Format(time.UnixDate))
	
	if err != nil { 
		return fmt.Errorf("CreateUser: %v", err) 
	}
	return nil 
}

func GetUser(email string) (*User, error) { 
	queryString := `
	select 
		id,
		name,
		email,
		password, 
		created_at
	from 
		users 
	where email=?`
	
	rows , err := db.Query(queryString, email)
	if err != nil { 
		return nil, fmt.Errorf("getUser: %v", err) 
	}

	defer rows.Close()
	if rows.Next() { 
		
		user := User{}
		err = rows.Scan(&user.Id, &user.UserName, &user.Email, 
						&user.Password, &user.CreatedAt)	
		if err != nil { 
			return nil, fmt.Errorf("getUser: %v", err) 
		}

		return &user, nil
	}  
		return nil, fmt.Errorf("getUser: %v", err) 
}

func CheckPassword(user *User, password string) (bool) { 
	
	if encryptPassword(password) == user.Password {
		return true 
	}
	return false
}


func encryptPassword(password string) string { 
	encrptPass := sha1.Sum([]byte(password))
	return string(encrptPass[:])
}
