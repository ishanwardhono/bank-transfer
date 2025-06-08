package transaction

const (
	insertTransactionQuery = `
		INSERT INTO transaction
			(source_account_id, destination_account_id, amount, reference_number)
		VALUES
			(:source_account_id, :destination_account_id, :amount, :reference_number)
		RETURNING id;`
)
