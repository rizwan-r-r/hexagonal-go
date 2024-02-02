package adapters

import (
	"database/sql"
	"hexagonalgo"
)

// This implements all ports for mysql database
// It can be seperate like
//
// type MysqlUsersRepo struct
// type MysqlCustomersRepo struct
//
// but as long it is same database connection behind we avoid that complexity

type MysqlRepo struct {
	DB *sql.DB
}

func (r *MysqlRepo) GetCustomers() []string {
	return nil
}

func (r *MysqlRepo) GetUsers() []string {
	return nil
}

var _ hexagonalgo.UserRepo = &MysqlRepo{}
var _ hexagonalgo.CustomerRepo = &MysqlRepo{}
