# Backend

## Run the Backend Locally

Install Golang, then run the following commands:

```bash
# Change directory to the backend
$ cd path/to/data-visualization-demo/backend

# Install dependencies specified in go.mod and go.sum
backend $ go mod download

# Run the backend
backend $ go run main.go
```

```bash
# Or compile the backend into an executable binary:
backend $ go build -o backend
# Run the backend
backend $ ./backend
```

The backend now is running on http://localhost:8080
View the api/data on http://localhost:8080/api/data

## Run the Backend in Docker Locally

```bash
# Change directory to the backend
$ cd path/to/data-visualization-demo/backend

# Buld the Docker image:
# name: data-visualization-demo-backend, tag: latest (default)
# context: . (current directory)
backend $ docker build -t data-visualization-demo-backend:latest .

# Run the container:
# port mapping: 8080:8080 (host:container)
backend $ docker run -p 8080:8080 data-visualization-demo-backend:latest
```

The backend now is running on http://localhost:8080
View the api/data on http://localhost:8080/api/data
