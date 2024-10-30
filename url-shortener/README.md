Certainly! Here’s a sample `README.md` file for your URL shortener project. You can customize it further based on your needs.

```markdown
# URL Shortener

A simple URL shortener built in Go. This application allows users to shorten long URLs and redirect them to their original destinations. It uses a MySQL database for data persistence.

## Project Structure
```

url-shortener/
├── cmd/
│ └── shortener/
│ └── main.go # Entry point, server setup
├── internal/
│ ├── handler/
│ │ └── handler.go # HTTP request handlers
│ ├── shortener/
│ │ └── shortener.go # URL shortening logic
│ ├── storage/
│ │ └── storage.go # Data persistence
│ └── model/
│ └── url.go # URL data structures
├── config/
│ └── config.go # Configuration management
├── go.mod
└── README.md

````

## Features

- Shorten long URLs
- Redirect to original URLs
- Data persistence using MySQL
- Simple HTTP API

## Requirements

- Go (version 1.16 or higher)
- MySQL (version 5.7 or higher)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/url-shortener.git
   cd url-shortener
````

2. Set up a MySQL database:

   ```sql
   CREATE DATABASE urlShortener;
   ```

3. Create a `.env` file in the root directory with the following content:

   ```plaintext
   DB_USER=root
   DB_PASSWORD=yourpassword
   DB_HOST=localhost
   DB_PORT=3306
   DB_NAME=urlShortener
   ```

4. Install the required Go dependencies:

   ```bash
   go mod tidy
   ```

## Running the Application

1. Start your MySQL server.
2. Run the application:

   ```bash
   go run cmd/shortener/main.go
   ```

3. The server will start on `http://localhost:8080`.

## API Endpoints

### Shorten URL

- **Endpoint:** `POST /shorten`
- **Request Body:**

  ```json
  {
    "original_url": "https://www.example.com"
  }
  ```

- **Response:**

  ```json
  {
    "short_url": "https://short.ly/abc123"
  }
  ```

### Redirect URL

- **Endpoint:** `GET /u/{short_url}`
- **Example:** `GET /u/abc123`
- **Response:** Redirects to the original URL.

## Testing

You can use tools like [Postman](https://www.postman.com) or `curl` to test the API endpoints.

### Example using curl

- Shorten a URL:

  ```bash
  curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"original_url": "https://www.example.com"}'
  ```

- Redirect to the original URL:

  ```bash
  curl -L http://localhost:8080/u/abc123
  ```

```

```
