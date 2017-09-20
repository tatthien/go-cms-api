package database

import (
	"strconv"

	"github.com/Machiel/slugify"
	"github.com/tatthien/go-cms-api/model"
	"github.com/tatthien/go-cms-api/ultil"
)

// InsertPost insert new post into database
func (dbfactory *Database) InsertPost(post model.Post) (model.Post, error) {
	stmt, err := dbfactory.db.Prepare("INSERT INTO `posts` (title, content, post_type, slug, is_publish, author_id, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return model.Post{}, err
	}

	if post.Created == "" {
		post.Created = ultil.GetCurrentMySQLDate()
	}

	if post.Updated == "" {
		post.Updated = ultil.GetCurrentMySQLDate()
	}

	// Default post type is `post`
	if post.PostType == "" {
		post.PostType = "post"
	}

	// Sanitize slug
	post.Slug = slugify.Slugify(post.Title)
	post.Slug = dbfactory.GetUniquePostSlug(post.Slug)

	res, err := stmt.Exec(post.Title, post.Content, post.PostType, post.Slug, post.IsPublish, post.Author, post.Created, post.Updated)
	if err != nil {
		return model.Post{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return model.Post{}, err
	}

	post, err = dbfactory.GetPostByID(id)
	return post, err
}

// UpdatePost update post
func (dbfactory *Database) UpdatePost(id int64, post model.Post) (model.Post, error) {
	stmt, err := dbfactory.db.Prepare("UPDATE `posts` SET title=?, content=?, is_publish=?, updated_at=? WHERE id = ?")
	if err != nil {
		return model.Post{}, err
	}

	post.Updated = ultil.GetCurrentMySQLDate()

	_, err = stmt.Exec(post.Title, post.Content, post.IsPublish, post.Updated, id)
	if err != nil {
		return model.Post{}, err
	}

	post, err = dbfactory.GetPostByID(id)
	return post, err
}

// DeletePost delete post by ID
func (dbfactory *Database) DeletePost(id int64) error {
	stmt, err := dbfactory.db.Prepare("DELETE FROM `posts` WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

// GetPostByID get post data by id
func (dbfactory *Database) GetPostByID(id int64) (model.Post, error) {
	stmt, err := dbfactory.db.Prepare("SELECT * FROM `posts` WHERE id = ?")
	if err != nil {
		return model.Post{}, err
	}

	var post model.Post
	rows := stmt.QueryRow(id)
	if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.PostType, &post.Slug, &post.IsPublish, &post.Author, &post.Created, &post.Updated); err != nil {
		return model.Post{}, err
	}

	return post, nil
}

// GetPostBySlug get post data by slug
func (dbfactory *Database) GetPostBySlug(slug string) (model.Post, error) {
	stmt, err := dbfactory.db.Prepare("SELECT * FROM `posts` WHERE slug = ?")
	if err != nil {
		return model.Post{}, err
	}

	var post model.Post
	rows := stmt.QueryRow(slug)
	if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.PostType, &post.Slug, &post.IsPublish, &post.Author, &post.Created, &post.Updated); err != nil {
		return model.Post{}, err
	}

	return post, nil
}

// GetPosts get list of post
func (dbfactory *Database) GetPosts(args map[string]string) (model.Posts, error) {
	defaultArgs := map[string]string{
		"page":      "1",
		"limit":     "10",
		"post_type": "post",
	}

	for k, v := range defaultArgs {
		if _, ok := args[k]; !ok {
			args[k] = v
		}
	}

	postType := args["post_type"]
	page, err := strconv.Atoi(args["page"])
	if err != nil {
		return model.Posts{}, err
	}
	limit, err := strconv.Atoi(args["limit"])
	if err != nil {
		return model.Posts{}, err
	}

	offset := (page - 1) * limit

	stmt, err := dbfactory.db.Prepare("SELECT * FROM `posts` WHERE `post_type` = ? LIMIT ?, ?")
	if err != nil {
		return model.Posts{}, err
	}

	rows, err := stmt.Query(postType, offset, limit)
	if err != nil {
		return model.Posts{}, err
	}

	var posts model.Posts
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.PostType, &post.Slug, &post.IsPublish, &post.Author, &post.Created, &post.Updated); err != nil {
			continue
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// GetUniquePostSlug get the unique slug
func (dbfactory *Database) GetUniquePostSlug(slug string) string {
	post, _ := dbfactory.GetPostBySlug(slug)
	if post.ID == 0 {
		return slug
	}

	var newSlug string
	for i := 1; i < 10; i++ {
		newSlug = slug + "-" + strconv.Itoa(i)
		post, _ := dbfactory.GetPostBySlug(newSlug)
		if post.ID == 0 {
			break
		}
	}
	return newSlug
}
