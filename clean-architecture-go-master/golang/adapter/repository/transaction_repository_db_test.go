package repository

import (
	"github.com/samuelterra22/clean-architecture-go/adapter/repository/fixture"
	"github.com/samuelterra22/clean-architecture-go/domain/entity"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTransactionRepositoryDb_Insert(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")

	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	repository := NewTransactionRepositoryDb(db)
	err := repository.Insert("1", "1", 2, entity.APPROVED, "")
	assert.Nil(t, err)
}
