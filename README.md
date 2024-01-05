# BookWise - Library Management System

BookWise is a comprehensive Library Management System designed to streamline the management of books, authors, genres, and user transactions in libraries. It provides a robust API with endpoints for various actions, ensuring efficient library operations.

## Features

- **Book Management:** Create, update, and delete book records. Retrieve details of all books or a specific book by ID.

- **Author Management:** Register authors, update their details, and delete author records.

- **Genre Management:** Create and retrieve details of genres, facilitating efficient categorization of books.

- **User Transactions:** Allow users to borrow and return books. Keep track of user transactions.

- **Token Generation:** Secure token generation for authentication and password reset functionality.

- **Health Check:** Monitor the application's health and version information.

- **Debugging:** Access application metrics using the `/debug/vars` endpoint.

## API Endpoints

- **GET /v1/healthcheck:** Show application health and version information.

- **GET /v1/books:** Retrieve details of all books.

- **POST /v1/books:** Create a new book.

- **GET /v1/books/:id:** Retrieve details of a specific book.

- **PATCH /v1/books/:id:** Update the details of a specific book.

- **DELETE /v1/books/:id:** Delete a specific book.

- **POST /v1/authors:** Register a new author.

- **PUT /v1/authors/:id:** Update the details of a specific author.

- **DELETE /v1/authors/:id:** Delete a specific author.

- **POST /v1/genres:** Create a new genre.

- **GET /v1/genres:** Retrieve details of all genres.

- **GET /v1/genres/:id:** Retrieve details of a specific genre.

- **POST /v1/users/borrow:** Register a book borrowing transaction for a user.

- **PUT /v1/users/return/:transaction_id:** Mark a book as returned for a specific user transaction.

- **GET /v1/users/:id/transactions:** Show all book borrowing transactions for a specific user.

- **POST /v1/tokens/authentication:** Generate a new authentication token.

- **POST /v1/tokens/password-reset:** Generate a new password-reset token.

- **GET /debug/vars:** Display application metrics.

## Getting Started

1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/bookwise.git
   cd bookwise

2.  **Install dependencies:**
    ```bash
    go mod download
3. **Run the application**
    ```bash
    go mod download
