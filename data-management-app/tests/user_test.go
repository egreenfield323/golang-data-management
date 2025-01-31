package tests

import (
	"data-management-app/db"
	"data-management-app/models"
	"testing"
)

var createdUserID int

// TestDatabaseConnection checks if the application can connect to the database
func TestDatabaseConnection(t *testing.T) {
	err := db.InitDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	t.Log("Successfully connected to the database!")
}

// TestCreateUser tests the CreateUser function
func TestCreateUser(t *testing.T) {
	// Initialize the database connection
	if err := db.InitDB(); err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}

	user := &models.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@email.com",
	}
	err := db.CreateUser(user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}
	createdUserID = user.ID
	t.Logf("Created user with ID: %d", user.ID)
}

// TestGetUser tests the GetUser function
func TestGetUser(t *testing.T) {
	if err := db.InitDB(); err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}

	user, err := db.GetUser(createdUserID)
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}
	if user.FirstName != "John" {
		t.Errorf("Expected first name 'John', got '%s'", user.FirstName)
	}
	t.Logf("Retrieved user: %+v", user)
}

// TestCleanup deletes the test user from the database
func TestCleanup(t *testing.T) {
	if err := db.InitDB(); err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}

	err := db.DeleteUser(createdUserID)
	if err != nil {
		t.Fatalf("Failed to delete user: %v", err)
	}
	t.Log("Deleted test user from the database")
}
