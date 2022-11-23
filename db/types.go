package db

type Thread struct {
	Id        int
	Topic     string
	UserId    int
	CreatedAt string
}

type Post struct {
	Id        int
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt string
}

type User struct {
	Id        int
	UserName  string
	Password  string
	Email     string
	SessionId int 
	CreatedAt string
}

type Session struct {
	Id        int
	UUID      string
	Email     string
	UserId    int   
	CreatedAt string
}
