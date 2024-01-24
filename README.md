# Microfrontend Server Project

This project is a microfrontend server setup, consisting of a frontend built
with [Single-SPA](https://single-spa.js.org/) and a backend server written in Go.

## Project Structure

```bash
my-mfe-server/
├── client/                           # Frontend part of the project
│ ├── src/                            # Frontend source files
│ │ ├── index.js                      # Entry point for your Single-SPA config
│ │ └── ...                           # Other frontend files
│ ├── public/                         # Public assets like HTML entry point
│ │ └── index.html                    # HTML file that serves as the entry point
│ ├── webpack.config.js               # Webpack configuration
│ ├── package.json
│ └── ...                             # Other frontend configuration files
└── server/                           # Backend (Golang) part of the project
├── cmd/
│ └── server/
│ └── main.go                         # The entry point of the server
├── internal/
│ ├── handler/                        # HTTP handlers
│ │ └── ...
│ ├── middleware/                     # HTTP middleware
│ │ └── ...
│ ├── model/                          # Data models
│ │ └── ...                           # Data model definitions
│ ├── repository/                     # Database interaction
│ │ └── ...                           # Database access logic
│ └── service/                        # Business logic
│ └── ...                             # Business logic implementations
├── pkg/                              # Shared libraries (if needed)
│ └── ...
├── static/                           # Static files directory for frontend assets
│ └── ...                             # Bundled frontend files
├── templates/                        # HTML templates (if needed)
│ └── ...
├── go.mod
├── go.sum
├── Dockerfile
└── Makefile                          # Makefile for task automation
```

## Getting Started

### Prerequisites

- Node.js and npm (for the frontend)
- Go (for the backend)
- Docker (optional, for containerization)

### Setting Up for Development

1. **Frontend Setup**:

   - Navigate to the `client` directory.
   - Install dependencies: `npm install`.
   - To build the frontend: `npm run build`.

2. **Backend Setup**:
   - Navigate to the `server` directory.
   - Build the Go server: `go build ./cmd/server`.
   - Run the server: `./bin/server`.

### Using the Makefile

- To build the project: `make build`.
- To run the project: `make run`.
- To clean up: `make clean`.
- To build and run with Docker: `make docker-build` and `make docker-run`.

## Docker Support

The project includes a Dockerfile for containerizing the application.

- Build the Docker image: `docker build -t mfe-server .`
- Run the Docker container: `docker run -p 8000:8000 mfe-server`

## Testing

- Run tests: `make test`.

## License

This project is licensed under the [MIT License](LICENSE).
