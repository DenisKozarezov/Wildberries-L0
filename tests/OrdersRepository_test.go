package database

import (
	"myapp/database"
	"myapp/services"
	"testing"
)

func TestOrdersRepository_SelectByUID(t *testing.T) {
	// Stage 1. Arrange
	uid := "o35fse23y63gs35test"
	database.ConnectToDatabase()
	services.RestoreCache()
	services.ConnectToNATS()

	// Stage 2. Act
	result, err := services.SelectOrderByUID(uid)

	// Stage 3. Assert
	if result == nil {
		t.Log("Result:", result)
		t.Errorf("Unable to query order. Expected: !nil. Result: %s", err)
	}
}
