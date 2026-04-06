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

---

## 📚 Flashcards — data model, seeding, and adding content

This section is for you (or another developer, or an LLM-assisted workflow) to **add vocabulary**, **tags**, and **paths** without digging through the codebase.

### MongoDB collections

| Collection | Purpose |
|------------|---------|
| **`flashcards`** | Individual study cards (word, translation, examples, etc.). |
| **`flashcard_paths`** | Metadata for each **language + level** row shown on the path map (title, description, order, icon). |

Indexes are created automatically on API startup (compound index on `language` + `path` + `difficulty`, text index on `word` + `description`).

### Card document shape (`flashcards`)

Fields the API expects (aligned with `POST /api/flashcards`):

| Field | Type | Rules |
|-------|------|--------|
| `word` | string | Required. The term or phrase on the front of the card. |
| `translation` | string | Required. Short gloss (can include notes, e.g. “NOT …”). |
| `type` | string | One of: `verb`, `noun`, `adjective`, `adverb`, `phrasal_verb`, `expression`. |
| `language` | string | `en` or `es`. |
| `path` | string | `beginner`, `intermediate`, or `advanced` (must match a path **level** for that language). |
| `difficulty` | int | **1–5** (used for sorting within a list). |
| `description` | string | Learner-friendly definition or explanation. |
| `examples` | string[] | **At least one** non-empty example sentence. |
| `tags` | string[] | **Optional.** Free-form labels (e.g. `daily-use`, `false-friend`, `travel`). Omit or use `[]` if none. |

`createdAt` / `updatedAt` are set by the server on create.

### Path document shape (`flashcard_paths`)

Paths are **not** created via the public REST API today. They are loaded from this collection for `GET /api/flashcards/paths`. Each row is one row on the UI map.

| Field | Type | Notes |
|-------|------|--------|
| `language` | string | `en` or `es`. |
| `level` | string | Must be exactly `beginner`, `intermediate`, or `advanced` (this is what cards use in their `path` field). |
| `title` | string | Short name (e.g. “First Steps”). |
| `description` | string | Subtitle on the path map. |
| `order` | int | Sort order **within the language** (1 = top of path). |
| `totalCards` | int | Informational; the API also **counts** live cards per `language` + `path` when possible. |
| `icon` | string | Emoji or single character shown in the bubble (e.g. `🌱`). |

**Adding a new path** (e.g. a fourth track):

1. Insert a document into **`flashcard_paths`** with a **new** `level` value only if you also change the **frontend** to understand it (today it only knows the three levels above). For the same three tiers in a new language, add rows with that `language` and `order` 1…3.
2. Or edit **`api/flashcards/seed_data.go`** (`seedPaths()`), then run the seed command with **`-force`** (see below).

### How to add more words (cards)

**Option A — HTTP API (recommended for scripts / LLM output)**

1. Set `FLASHCARDS_ADMIN_KEY` in `.env`.
2. `POST /api/flashcards` with header `X-Admin-Key: <same value>` and JSON body.

Minimal example **with tags**:

```bash
curl -X POST http://localhost:8080/api/flashcards \
  -H "Content-Type: application/json" \
  -H "X-Admin-Key: YOUR_SECRET" \
  -d '{
    "word": "carry on",
    "translation": "continuar / seguir em frente",
    "type": "phrasal_verb",
    "language": "en",
    "path": "intermediate",
    "difficulty": 3,
    "description": "To continue doing something, especially after an interruption.",
    "examples": [
      "We can carry on with the plan tomorrow.",
      "She carried on working despite the noise."
    ],
    "tags": ["daily-use", "work"]
  }'
```

**Option B — Code + seed**

Add entries in **`api/flashcards/seed_data.go`** (e.g. `enIntermediate()`, `esBeginner()`, …) using the same shape as existing `mkCard(...)` calls, then run:

```bash
go run ./cmd/seed-flashcards -e development -force
```

`-force` **deletes all documents** in `flashcards` and `flashcard_paths` and re-inserts from code. Use only when you accept a full reset, or maintain seed data as the source of truth.

**Option C — MongoDB directly**

Insert into **`flashcards`** with fields matching the table above, plus `createdAt` and `updatedAt` as BSON dates (UTC). Ensure `language` + `path` match an existing path row’s `language` + `level`.

### How to delete a card

```bash
curl -X DELETE http://localhost:8080/api/flashcards/CARD_OBJECT_ID_HEX \
  -H "X-Admin-Key: YOUR_SECRET"
```

### Using an LLM to bulk-generate cards

1. Share this README section (or the **Card document shape** table) as context.
2. Ask for **valid JSON only**: an array of objects, each object having exactly the POST fields (`word`, `translation`, `type`, `language`, `path`, `difficulty`, `description`, `examples`, and optionally `tags`).
3. Validate enums: `type`, `language`, `path` must match the allowed values above; `examples` must be a non-empty array of strings.
4. Pipe each object to `curl` or a small script that calls `POST /api/flashcards` with `X-Admin-Key`.

Example prompt fragment for an LLM:

> Generate 10 JSON objects for English beginner phrasal verbs. Each object: word, translation (Portuguese gloss OK), type `phrasal_verb`, language `en`, path `beginner`, difficulty 1–5, description, examples (array of 2 strings), tags (array, e.g. `["travel"]`). Output a single JSON array, no markdown.

### Quick reference — allowed values

- **`type`:** `verb` · `noun` · `adjective` · `adverb` · `phrasal_verb` · `expression`
- **`language`:** `en` · `es`
- **`path` (same as path `level`):** `beginner` · `intermediate` · `advanced`

### Project files involved

| What | Where |
|------|--------|
| HTTP handlers, list filters | `api/flashcards/flashcards_controller.go`, `flashcards_service.go` |
| Request validation (POST body) | `api/flashcards/flashcards.go` → `CreateFlashcardPayload` |
| Default paths + starter cards | `api/flashcards/seed_data.go`, `seed.go` |
| Seed CLI | `cmd/seed-flashcards/main.go` |
| Routes | `router/flashcards_routes.go` |

## 🤝 Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## 📄 License

This project is licensed under the [MIT License](LICENSE).
