package hexagonalgo

// ports defain what are ports for this package

// Generally THIS FILE SHOULD NOT EXISTS, as Interfaces and models should be defined where are used in your DOMAIN code

// This can be seperate file for visual

// Database interfaces as example

type CustomerRepo interface {
	GetCustomers() []string
}
type UserRepo interface {
	GetUsers() []string
}

// Compose to all but pass them as seperate in your domain code
type Repo interface {
	CustomerRepo
	UserRepo
}
