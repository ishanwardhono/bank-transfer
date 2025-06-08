package account

const (
	insertAccountQuery = `
		INSERT INTO account
			(id, balance)
		VALUES
			(:id, :balance);`
)
