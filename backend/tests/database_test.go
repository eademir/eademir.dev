package test

import (
	"testing"

	"eademir.dev/database"
	"github.com/stretchr/testify/assert"
)

func TestRunDatabase(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	db, err := database.DatabaseConnection()
	if err == nil {
		defer db.Close()
	}
	assert.True(t, err != nil, err)
}
