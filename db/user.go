package db 

import ( 
	"log"
	"time"
	"crypto/sha1"
)

func CreateUser(name, password, email string) {  
	
	queryString := `
	insert into 
		users(name, email, password, created_at)
	values(?, ?, ?, ?)`

	if db == nil { 
		return 
	}
	_, err := db.Exec(queryString, name, email, 
	encryptPassword(password), time.Now().Format(time.UnixDate))
	
	if err != nil { 
		log.Println("CreateUser: ", err) 
	}

}

func encryptPassword(password string) string { 
	encrptPass := sha1.Sum([]byte(password))
	return string(encrptPass[:])
}
