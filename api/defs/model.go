package defs

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Pwd string `json:"pwd"`
}
type CreateUser struct {
	Username string `json:"username"`
	Pwd string `json:"pwd"`
}
type VidioInfo struct{
	Id string `json:"id"`
	Authid int `json:"author_id"`
	Name string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}

type Comment struct{
	Id string `json:"id"`
	Vidio_id string `json:"vid"`
	Author_name string `json:"username"`
	Content string `json:"content"`
	Author_id int `json:"author_id"`
}
type SimpleSession struct {
	Username string
	TTL int64
}

type SignedUp struct{
	Success bool `json:"success"`
	Session string `json:"session_id"`
}
