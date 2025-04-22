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

## 🤝 Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## 📄 License

This project is licensed under the [MIT License](LICENSE).
