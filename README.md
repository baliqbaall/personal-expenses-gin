# Personal Expenses Tracker

A simple Go application using the Gin framework to track personal expenses.

## Installation

1. Clone the repository.
2. Install dependencies with `go mod tidy`.
3. Run the application with `go run main.go`.

## Usage

- Add expenses via API.
- Retrieve and manage expenses.
- Update or delete expenses.

## API Endpoints

- GET `/expenses`: Retrieve all expenses.
- POST `/expenses`: Add a new expense.
- PUT `/expenses/{id}`: Update an existing expense.
- DELETE `/expenses/{id}`: Delete an expense by ID.

## Technologies Used

- Go programming language
- Gin web framework
- MySQL database

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
