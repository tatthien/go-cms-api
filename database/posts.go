package database

import (
	"github.com/tatthien/go-cms-api/model"
	"github.com/tatthien/go-cms-api/ultil"
)

// InsertPost insert new post into database
func (dbfactory *Database) InsertPost(post model.Post) (model.Post, error) {
	stmt, err := dbfactory.db.Prepare("INSERT INTO `posts` (title, content, author_id, created_at, updated_at) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return model.Post{}, err
	}

	if post.Created == "" {
		post.Created = ultil.GetCurrentMySQLDate()
	}

	if post.Updated == "" {
		post.Updated = ultil.GetCurrentMySQLDate()
	}

	res, err := stmt.Exec(post.Title, post.Content, post.Author, post.Created, post.Updated)
	if err != nil {
		return model.Post{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return model.Post{}, err
	}

	lastInsertPost, err := dbfactory.GetPost(id)
	return lastInsertPost, err
}

// GetPost get post data by id
func (dbfactory *Database) GetPost(id int64) (model.Post, error) {
	stmt, err := dbfactory.db.Prepare("SELECT * FROM `posts` WHERE id = ?")
	if err != nil {
		return model.Post{}, err
	}

	var post model.Post
	rows := stmt.QueryRow(id)
	if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author, &post.Created, &post.Updated); err != nil {
		return model.Post{}, err
	}

	return post, nil
}

// GetPosts get list of post
func (dbfactory *Database) GetPosts(offset, limit int) (model.Posts, error) {
	stmt, err := dbfactory.db.Prepare("SELECT * FROM `posts` LIMIT ?, ?")
	if err != nil {
		return model.Posts{}, err
	}

	rows, err := stmt.Query(offset, limit)
	if err != nil {
		return model.Posts{}, err
	}

	var posts model.Posts
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author, &post.Created, &post.Updated); err != nil {
			continue
		}
		posts = append(posts, post)
	}
	return posts, nil
}
