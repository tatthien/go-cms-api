package model

// Post model
type Post struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	PostType string `json:"post_type"`
	Slug     string `json:"slug"`
	Author   int64  `json:"author"`
	Created  string `json:"created_at"`
	Updated  string `json:"updated_at"`
}

// Posts model
type Posts []Post

// User model
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string
	Email    string `json:"email"`
}

// Users model
type Users []User
