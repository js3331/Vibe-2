# Vibe - Simple Text Post Web Application

A lightweight web application that allows users to create and search text posts. Built with Go, PostgreSQL, and vanilla JavaScript.

## Features

- Create text posts
- Search through existing posts
- Real-time updates
- Simple and responsive UI
- Docker containerized deployment

## Tech Stack

- **Backend:**
  - Go (Golang)
  - Gorilla Mux (Router)
  - PostgreSQL (Database)
  - `lib/pq` (PostgreSQL driver)

- **Frontend:**
  - HTML5
  - CSS3
  - Vanilla JavaScript

- **Infrastructure:**
  - Docker
  - Docker Compose

## Project Structure

```
.
├── cmd/
│   └── main.go           # Application entry point
├── internal/
│   ├── handlers/         # HTTP request handlers
│   └── models/          # Data models
├── static/
│   ├── css/             # Stylesheets
│   └── js/              # JavaScript files
├── templates/           # HTML templates
├── docker-compose.yml   # Docker Compose configuration
├── Dockerfile          # Docker build instructions
├── init.sql           # Database initialization
└── go.mod             # Go module definition
```

## Prerequisites

- Docker
- Docker Compose

## Getting Started

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd vibe
   ```

2. Start the application:
   ```bash
   docker-compose up --build
   ```

3. Access the application:
   - Open your web browser and navigate to `http://localhost:8080`

## Development

### Local Development Setup

1. Install Go (1.21 or later)
2. Install PostgreSQL
3. Install dependencies:
   ```bash
   go mod download
   go mod tidy
   ```

### Database Configuration

The application uses environment variables for database configuration:
- `DB_HOST`: Database host (default: "db")
- `DB_USER`: Database user (default: "postgres")
- `DB_PASSWORD`: Database password (default: "postgres")
- `DB_NAME`: Database name (default: "vibe")
- `DB_PORT`: Database port (default: "5432")

### Making Changes

1. **Backend Changes:**
   - Edit files in `cmd/` or `internal/`
   - Rebuild and restart the containers:
     ```bash
     docker-compose down
     docker-compose up --build
     ```

2. **Frontend Changes:**
   - Modify files in `static/` or `templates/`
   - No rebuild needed, just refresh the browser

3. **Database Changes:**
   - Update `init.sql`
   - Rebuild and restart:
     ```bash
     docker-compose down -v  # Remove volumes to reset database
     docker-compose up --build
     ```

## Maintenance

### Logs

View application logs:
```bash
docker-compose logs -f app
```

View database logs:
```bash
docker-compose logs -f db
```

### Backup

Backup the database:
```bash
docker exec vibe-db pg_dump -U postgres vibe > backup.sql
```

### Monitoring

1. Check container status:
   ```bash
   docker-compose ps
   ```

2. Monitor container resources:
   ```bash
   docker stats vibe-app vibe-db
   ```

## Troubleshooting

1. **Database Connection Issues:**
   - Check if the database container is running:
     ```bash
     docker-compose ps
     ```
   - Verify database logs:
     ```bash
     docker-compose logs db
     ```

2. **Application Errors:**
   - Check application logs:
     ```bash
     docker-compose logs app
     ```

3. **Reset Everything:**
   ```bash
   docker-compose down -v
   docker-compose up --build
   ```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
