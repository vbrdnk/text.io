package database

import (
	"testing"

	"text.io/configs"
)

// Setup test database
func setupTestDB(t *testing.T) {
	// Create test configuration
	config := configs.Config{
		DBHost:     "localhost",
		DBPort:     5432,
		DBUser:     "postgres",
		DBPassword: "postgres",
		DBName:     "test_goapi", // Use a different database for tests
	}

	// Initialize test database
	if err := InitDB(config); err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}

	// Clean up existing test data
	_, err := DB.Exec("DELETE FROM links")
	if err != nil {
		t.Fatalf("Failed to clean up test database: %v", err)
	}
}

func TestCreateAndGetItem(t *testing.T) {
	setupTestDB(t)
	defer CloseDB()

	// Test item
	itemID := "test123"
	itemName := "Test Item"

	// Insert test item
	_, err := DB.Exec("INSERT INTO links (id, name) VALUES ($1, $2)",
		itemID, itemName)
	if err != nil {
		t.Fatalf("Failed to insert test item: %v", err)
	}

	// Retrieve the item
	var name string
	err = DB.QueryRow("SELECT name FROM links WHERE id = $1", itemID).
		Scan(&name)
	if err != nil {
		t.Fatalf("Failed to retrieve test item: %v", err)
	}

	// Verify the item
	if name != itemName {
		t.Errorf("Retrieved item name doesn't match: got %v want %v",
			name, itemName)
	}
}
