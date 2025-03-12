package dao

import (
	"gofr.dev/pkg/gofr/datasource/mongo"
)

type UserDAO struct {
	db *mongo.Client
}

// NewUserRepository creates a new instance of UserRepository
func NewUserDAO(db *mongo.Client) *UserDAO {
	return &UserDAO{db: db}
}
