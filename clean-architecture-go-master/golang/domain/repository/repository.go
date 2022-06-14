package repository

type TransactionRepository interface {
	Insert(id string, accountId string, amount float64, status string, errorMessage string) error
}
