package db

import (
	"database/sql"

	"data-management-app/models"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Initialize database connection
func InitDB() error {
	var err error
	dsn := "<USERNAME HERE (typically root)>:PASSWORDHERE@tcp(<SERVER IP (or localhost)>:3306)/data_management?parseTime=true"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return db.Ping()
}

// Get all users from the database
func GetAllUsers() ([]models.User, error) {
	rows, err := db.Query("SELECT id, first_name, last_name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Create a new user in the database
func CreateUser(user *models.User) error {
	query := `INSERT INTO users (first_name, last_name, email) VALUES (?, ?, ?)`
	result, err := db.Exec(query, user.FirstName, user.LastName, user.Email)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(id)
	return nil
}

// Retrieve a user by ID from the database
func GetUser(id int) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, first_name, last_name, email FROM users WHERE id = ?`
	err := db.QueryRow(query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update a user in the database
func UpdateUser(user *models.User) error {
	query := `UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?`
	_, err := db.Exec(query, user.FirstName, user.LastName, user.Email, user.ID)
	return err
}

// Delete a user from the database
func DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}
