# Book Store API

This is a simple Book Store API built using Go language, Fiber framework, and PostgreSQL as the database. The API handles basic CRUD operations for managing books.

## Prerequisites

Before you start, ensure you have the following installed:

- [Go](https://github.com/golang/go)
- PostgreSQL
- [Fiber](https://github.com/gofiber/fiber)

## Getting Started

1. **Clone the repository:**

    ```bash
    git clone https://github.com/nainglinnphyo/book-store-api.git
    cd book-store-api
    ```

2. **Install Dependencies:**

    ```bash
    go mod tidy
    ```

3. **Set up PostgreSQL:**

    - Create a new database for the project.
    - Update the `DATABASE_URL` in the `.env` file with your PostgreSQL connection details.

4. **Environment Variables:**

    Create a `.env` file in the root of the project and add the following variables:

    ```env
     DB_HOST=
     DB_PORT=
     DB_USER=
     DB_PASSWORD=
     DB_NAME=
    ```

    Replace `user`, `password`, `localhost`, `port`, and `database` with your PostgreSQL credentials and database name.

5. **Generate Swagger Docs:**

    ```bash
    swag init
    ```

    The Docs will be generate in `docs` folder.

6. **Run the Application:**

    ```bash
    go run main.go
    ```

    The API will be running at `http://localhost:3000`.

## API Endpoints

- **GET /books**: Get all books
- **GET /books/{id}**: Get a single book by ID
- **POST /books**: Add a new book
- **PUT /books/{id}**: Update a book by ID
- **DELETE /books/{id}**: Delete a book by ID

## Example Usage

```bash
# Get all books
curl http://localhost:3000/books

# Get a single book by ID
curl http://localhost:3000/books/1

# Add a new book
curl -X POST -H "Content-Type: application/json" -d '{"title":"New Book","author":"Author Name"}' http://localhost:3000/books

# Update a book by ID
curl -X PUT -H "Content-Type: application/json" -d '{"title":"Updated Book"}' http://localhost:3000/books/1

# Delete a book by ID
curl -X DELETE http://localhost:3000/books/1
