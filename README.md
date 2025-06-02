# MindShelf

MindShelf is a personal link management system that allows users to save, organize, and search bookmarks.

## Features

- User authentication (login/register)
- Save links with titles, descriptions, and tags
- Automatic metadata extraction when adding links
- Full-text search functionality for finding links quickly
- Tag-based organization and filtering
- Responsive web interface with dark mode support
- Keyboard shortcuts for power users
- RESTful API with JWT authentication
- Multi-container Docker orchestration

## Architecture

MindShelf uses a modern, containerized architecture:

- **Frontend**: Minimalistic UI built with Alpine.js - a lightweight JavaScript framework with no build step
- **Backend API**: Go (Chi router) providing RESTful endpoints
- **Database**: MongoDB for data storage
- **Reverse Proxy**: Nginx for routing, SSL termination, and serving static files

## Development Setup

### Prerequisites

- Docker and Docker Compose
- Go 1.24+ (for local development without Docker)
- OpenSSL (for generating self-signed certificates)

### Quick Start

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/mindshelf.git
   cd mindshelf
   ```

2. Create a `.env` file:
   ```
   cp .env.example .env
   ```
   Edit the `.env` file and set your desired configuration values.

3. Run the setup script:
   ```
   ./mindshelf.sh setup
   ```

4. Access the application at https://localhost

### Development Scripts

- `./mindshelf.sh start` - Start the development environment
- `./mindshelf.sh stop` - Stop the development environment
- `./mindshelf.sh logs` - View logs from all services
- `./mindshelf.sh logs api` - View logs only from the API service
- `./mindshelf.sh build` - Rebuild the Docker images

### Alternative: Using Make

- `make dev` - Start development environment
- `make dev-stop` - Stop development environment
- `make logs` - View logs
- `make test` - Run tests
- `make clean` - Clean up build artifacts

### Manual Development

If you prefer to run the Go server directly:

1. Install MongoDB locally or use a remote instance
2. Update your `.env` file with the correct MongoDB connection URI
3. Run the server:
   ```
   make run
   ```

## Production Deployment

### Using Docker Compose

1. Ensure you have proper SSL certificates for your domain
2. Place your SSL certificates in `docker/nginx/certs/`:
   - Certificate: `cert.pem`
   - Private key: `privkey.pem`

3. Update the Nginx configuration in `docker/nginx/nginx.prod.conf` with your domain name

4. Run the deployment script:
   ```
   sudo ./deploy-prod.sh
   ```
   
### Manual Deployment

1. Build for production:
   ```
   make build-prod
   ```

2. Start the production stack:
   ```
   docker-compose -f docker-compose.prod.yml up -d
   ```

### Managing Production

- View production logs:
  ```
  make logs-prod
  ```

- Stop production services:
  ```
  docker-compose -f docker-compose.prod.yml down
  ```

- Backup the database:
  ```
  make db-backup
  ```

- Restore from a backup:
  ```
  make db-restore BACKUP=./backups/mongodb-20250601-120000
  ```

## API Documentation

### Authentication

#### Register a new user
```
POST /api/auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "securepassword"
}
```

#### Login
```
POST /api/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "securepassword"
}
```

Response:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Links

All link endpoints require authentication with a JWT token in the Authorization header:
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

#### Get all links
```
GET /api/links
```

#### Get a specific link
```
GET /api/links/{id}
```

#### Create a new link
```
POST /api/links
Content-Type: application/json

{
  "url": "https://example.com",
  "title": "Example Website",
  "description": "An example website",
  "tags": ["example", "website"]
}
```

#### Update a link
```
PUT /api/links/{id}
Content-Type: application/json

{
  "url": "https://updated-example.com",
  "title": "Updated Example",
  "description": "An updated example website",
  "tags": ["updated", "example"]
}
```

#### Delete a link
```
DELETE /api/links/{id}
```

#### Search links
```
GET /api/links/search?q=searchterm
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
