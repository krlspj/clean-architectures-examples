package factory

import "github.com/samuelterra22/clean-architecture-go/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
