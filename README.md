# Money Transfer System

A money transfer system built with Go, featuring account management and secure money transfers between accounts.

## Prerequisites

### Docker & Docker Compose
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## How to Run

### Installation

The application can be quickly set up using Make commands:

```bash
# Start the application and database
make install

# To rebuild and deploy changes
make deploy

# To stop and remove all containers and volumes
make clean
```

### Testing the API

You can test the API in two ways:

1. **Using Postman**:
   - Import the provided Postman collection: `transfer-system.postman_collection.json`
   - The collection contains ready-to-use requests for all available endpoints
   - Sample responses for all endpoints are included in the collection for reference

2. **Direct API Calls**:
   - The API is available at `http://localhost:8080`
   - See the API documentation section below for endpoint details

## Project Overview

The Money Transfer System is a microservice designed to handle money transfers between accounts in a secure and reliable manner. It provides RESTful APIs for account management and transaction processing.

> **Note:** You can follow my step-by-step development process through the commit history of this repository. Each commit is structured to show my thought process, implementation decisions, and iterative improvements to the codebase.

### Key Features

- **Account Management**: Create and retrieve account information
- **Money Transfers**: Securely transfer funds between accounts
- **Transaction Management**: Track all financial transactions
- **Graceful Shutdown**: Ensure ongoing transactions are completed before shutdown

### System Architecture

The application follows a clean architecture pattern with clear separation of concerns:

- **HTTP Handlers**: Process incoming API requests
- **Services**: Implement business logic
- **Repositories**: Handle data persistence
- **Entities**: Define data structures

### Data Model

The system uses two main database tables:

1. **Account Table**:
   - Stores account information including current balance
   - Uses soft deletion for audit purposes

2. **Transaction Table**:
   - Records all transfers between accounts
   - Maintains source and destination account IDs, amount, and reference numbers
   - Indexes for efficient querying

## API Documentation

### Accounts

#### Register a New Account
- **Endpoint**: `POST /accounts`
- **Description**: Creates a new account with an initial balance
- **Request Body**:
  ```json
  {
    "account_id": 1001,
    "initial_balance": "1000.50"
  }
  ```
- **Response**: 201 Created

#### Get Account Information
- **Endpoint**: `GET /accounts/{accountId}`
- **Description**: Retrieves account information by ID
- **Response**:
  ```json
  {
    "account_id": 1001,
    "balance": "1000.50"
  }
  ```

### Transactions

#### Transfer Money
- **Endpoint**: `POST /transactions`
- **Description**: Transfers money from one account to another
- **Request Body**:
  ```json
  {
    "source_account_id": 1001,
    "destination_account_id": 1002,
    "amount": "50.25"
  }
  ```
- **Response**: 200 OK

## Technical Implementation Details

### Transaction Integrity

The system uses database transactions to ensure data integrity during money transfers. If any step of the transfer process fails (e.g., account validation, balance updates), the entire transaction is rolled back, maintaining consistency across accounts.

### Graceful Shutdown

The application implements a graceful shutdown mechanism that handles server stop signals properly. When shutdown is initiated, the server stops accepting new requests but allows in-progress transfers to complete, preventing data corruption and dangling database transactions (which could cause db locking) during system restarts.

### Context-Based Logging

The application implements a sophisticated context-based logging system that propagates request and account identifiers throughout the entire request lifecycle. This makes debugging and request tracing straightforward:

- Each request receives a unique `request_id` that flows through all log entries
- The `account_id` is extracted from requests and attached to the context
- All log entries include these contextual identifiers for easy correlation

Example log output:
```json
{"context":{"account_id":95,"request_id":"57caf02c18e6/0arzSQj4Go-000006"},"level":"error","msg":"failed to insert account, err: dial tcp: lookup postgres on 127.0.0.11:53: no such host","source":"/app/internal/service/account/account.go:23","time":"2025-06-09T00:09:39Z"}
{"context":{"account_id":95,"request_id":"57caf02c18e6/0arzSQj4Go-000006"},"level":"error","msg":"Failed to register account, err: dial tcp: lookup postgres on 127.0.0.11:53: no such host","source":"/app/internal/handler/http/handler.go:63","time":"2025-06-09T00:09:39Z"}
{"context":{"request_id":"57caf02c18e6/0arzSQj4Go-000006"},"level":"error","msg":"Request: GET /accounts/95 - 500 | 109.335417ms | 192.168.65.1:57251","source":"/app/pkg/httphelper/helper.go:76","time":"2025-06-09T00:09:39Z"}
```

This approach enables efficient debugging and monitoring, as related log entries can be easily identified and filtered by their context identifiers.

## Development

### Testing

The service layer is thoroughly tested with unit tests, achieving 100% code coverage in service layer. This ensures reliability and correctness of the business logic handling account operations and money transfers.

- All service components have accompanying `_test.go` files
- Test mocks are generated with mockgen in the `test/mock` directory
- Database operations are mocked to isolate service layer logic

To run the tests:

```bash
# Run all tests
make run-test

# Run tests with coverage report
make test-coverage
```

### Project Structure

```
transfer-system/
├── cmd/                   # Application entry point
├── data/                  # Database scripts
├── internal/
│   ├── entity/            # Domain models and DTOs
│   ├── handler/           # HTTP handlers
│   ├── repository/        # Data access layer
│   └── service/           # Business logic
│       ├── account/       # Account service
│       └── transaction/   # Transaction service
├── pkg/                   # Reusable packages
│   ├── config/            # Configuration
│   ├── context/           # Context utilities
│   ├── db/                # Database connection
│   ├── errors/            # Error handling
│   ├── httphelper/        # HTTP utilities
│   ├── logger/            # Logging
│   └── utils/             # Utility functions
├── docker-compose.yaml    # Docker Compose configuration
├── Dockerfile             # Docker image definition
├── go.mod                 # Go modules
└── Makefile               # Build automation
```

