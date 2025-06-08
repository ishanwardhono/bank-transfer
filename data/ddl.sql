CREATE TABLE account (
    id INTEGER PRIMARY KEY,
    balance NUMERIC(15,5) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);
CREATE INDEX idx_account_deleted_at ON account(deleted_at) WHERE deleted_at IS NOT NULL;

CREATE TABLE transaction (
    id SERIAL PRIMARY KEY,
    source_account_id INTEGER NOT NULL,
    destination_account_id INTEGER NOT NULL,
    amount NUMERIC(10,5) NOT NULL,
    reference_number VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (source_account_id) REFERENCES account(id),
    FOREIGN KEY (destination_account_id) REFERENCES account(id)
);
CREATE INDEX idx_transaction_reference_number ON transaction(reference_number);
CREATE INDEX idx_transaction_source_account ON transaction(source_account_id);
CREATE INDEX idx_transaction_destination_account ON transaction(destination_account_id);
CREATE INDEX idx_transaction_created_at ON transaction(created_at);
CREATE INDEX idx_transaction_deleted_at ON transaction(deleted_at) WHERE deleted_at IS NOT NULL;