package adapters

import "hexagonalgo"

// This implements all ports for mysql database
type RedisRepo struct {
}

func (r *RedisRepo) GetCustomers() []string {
	return nil
}

func (r *RedisRepo) GetUsers() []string {
	return nil
}

var _ hexagonalgo.UserRepo = &RedisRepo{}
