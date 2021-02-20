package main

type Post struct{
	Id int
	Content string
	Author string `"sql:not null"`
	Comment []Comment
	CreatedAt time.Time
}

func main() {
	
}