package database

import "github.com/tatthien/go-cms-api/model"
import "errors"

// InsertUser - insert a user to database
func (dbfactory Database) InsertUser(user model.User) (model.User, error) {
	// Check if username already existed
	exist, err := dbfactory.CheckUsernameExists(user.Username)
	if exist == true {
		return model.User{}, err
	}

	stmt, err := dbfactory.db.Prepare("INSERT INTO `users` (username, password, email) VALUES(?, ?, ?)")
	if err != nil {
		return model.User{}, err
	}

	res, err := stmt.Exec(user.Username, user.Password, user.Email)
	if err != nil {
		return model.User{}, err
	}

	id, _ := res.LastInsertId()

	lastInsertUser, err := dbfactory.GetUser(id)
	return lastInsertUser, err
}

// GetUser - get user from database by user id
func (dbfactory Database) GetUser(id int64) (model.User, error) {
	stmt, err := dbfactory.db.Prepare("SELECT id, username, email FROM `users` WHERE id = ?")
	if err != nil {
		return model.User{}, err
	}

	var user model.User
	rows := stmt.QueryRow(id)
	if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
		return model.User{}, err
	}
	return user, nil
}

// CheckUsernameExists - check if the username already existed
func (dbfactory Database) CheckUsernameExists(username string) (bool, error) {
	stmt, err := dbfactory.db.Prepare("SELECT EXISTS(SELECT 1 FROM `users` WHERE username = ?)")
	if err != nil {
		return true, errors.New("an error while checking the username")
	}

	var exist bool
	rows := stmt.QueryRow(username)
	if err := rows.Scan(&exist); err != nil {
		return true, errors.New("an error while checking the username")
	}

	if exist == true {
		return true, errors.New("this username already exists")
	}

	return false, nil
}
