# 📦 Portfolio API

A RESTful API developed in Go to provide dynamic data for my personal portfolio. It manages information such as projects, experiences, skills, and contacts, allowing centralized updates and integration with the frontend.

## 🚀 Technologies Used

- **Go (Golang)**
- **Gin** – Lightweight and fast web framework
- **MongoDB** – NO-SQL database
- **Docker** – Containerization for consistent environments
- **Godotenv** & **Viper** – To config env vars and environments
- **Go-playground/validator** – Validate request payload
- **Sendgrid** – Provider to send email

## 📁 Project Structure

```
portfolio-api/
├── api/           # API handlers and controllers
├── config/        # Configuration and environment variables
├── infra/         # Infrastructure (e.g., database connection)
├── middlewares/   # Custom middlewares (e.g., authentication)
├── router/        # API route definitions
├── server/        # Server initialization and configuration
├── utils/         # Utility functions
├── main.go        # Application entry point
├── Dockerfile     # Dockerfile for containerization
└── .env.example   # Example environment variables
```

## ⚙️ Setup and Execution

### Prerequisites

- [Go](https://golang.org/doc/install) installed
- [Docker](https://www.docker.com/get-started) (optional, for containerization)

### Running Locally

1. Clone the repository:

```bash
git clone https://github.com/viniciustakedi/portfolio-api.git
cd portfolio-api
```

2. Copy the `.env.example` file to `.env` and configure the environment variables as needed.

3. Install dependencies and run the application:

```bash
go mod tidy
go run main.go
```

The API will be available at `http://localhost:8080`.

### Using Docker

1. Build the Docker image:

```bash
docker build -t portfolio-api .
```

2. Copy the `.env.example` file to `.env` and configure the environment variables as needed.

3. Run the container:

```bash
docker run -d -p 8080:8080 --env-file .env portfolio-api
```

The API will be available at `http://localhost:8080/api`.

## 🛠️ Main Endpoints

- `GET /health` – Just to check if API is running
- `GET /jobs` – List all jobs
- `POST /email/send/portfolio-message` – Send a email with a user's message

*Additional endpoints for experiences, skills, and contacts are available.*

### Flashcards

- `GET /api/flashcards?language=en|es&path=beginner|intermediate|advanced&limit=&skip=` – List cards (paginated)
- `GET /api/flashcards/paths?language=en|es` – Learning paths for a language
- `GET /api/flashcards/:id` – Single card

### API admin — Flashcards

Some flashcard routes are **admin-only**: creating and deleting cards. They are not tied to user login; the API uses a **shared secret** so only you (or trusted tooling) can mutate vocabulary data.

| Item | Purpose |
|------|--------|
| **`FLASHCARDS_ADMIN_KEY`** (environment variable) | Secret value stored on the server. If it is empty, `POST` and `DELETE` flashcard routes respond with **503** (“not configured”). |
| **`X-Admin-Key`** (HTTP request header) | On each admin request, the client must send this header with the **same value** as `FLASHCARDS_ADMIN_KEY`. If it is missing or wrong, the API responds with **401**. |

**Example** (create a card with `curl`):

```bash
curl -X POST http://localhost:8080/api/flashcards \
  -H "Content-Type: application/json" \
  -H "X-Admin-Key: YOUR_SECRET_FROM_ENV" \
  -d '{"word":"example","translation":"ejemplo","type":"noun","language":"en","path":"beginner","difficulty":2,"description":"A sample.","examples":["An example sentence."]}'
```

CORS is configured to allow the `X-Admin-Key` header from allowed origins, so a browser-based admin tool can call the API when the origin is permitted.

## 🤝 Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## 📄 License

This project is licensed under the [MIT License](LICENSE).
