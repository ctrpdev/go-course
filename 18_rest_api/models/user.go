package models

import "example.com/db"

type User struct {
	ID       int64
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (u *User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	r, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}
	u.ID, err = r.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}
