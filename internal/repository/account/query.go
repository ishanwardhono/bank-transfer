package account

const (
	insertAccountQuery = `
		INSERT INTO account
			(id, balance)
		VALUES
			(:id, :balance);`

	getAccountQuery = `
		SELECT
			id, balance, created_at, updated_at, deleted_at
		FROM
			account
		WHERE
			id = $1
			AND deleted_at IS NULL;
	`
	updateBalanceQuery = `
		UPDATE account
		SET 
			balance = $2,
			updated_at = NOW()
		WHERE
			id = $1
			AND deleted_at IS NULL;`
)
