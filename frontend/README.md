# Frontend

## Run the Frontend Locally

Install Node.js, then run the following commands:

```bash
# Change directory to the frontend
$ cd path/to/data-visualization-demo/frontend

# Install dependencies specified in package.json
frontend $ npm install

# Run the frontend with Vite dev server
frontend $ npm run dev
```

The frontend now is running on http://localhost:5173

```bash
# Run the frontend and specify the host (your LAN IP address)
frontend $ npm run dev -- --host
```

The frontend now is running on http://192.168.0.101:5173/

## Run the Frontend in Docker Locally
```bash
# Change directory to the frontend
$ cd path/to/data-visualization-demo/frontend

# Buld the Docker image:
# name: data-visualization-demo-frontend, tag: latest (default)
# context: . (current directory)
frontend $ docker build -t data-visualization-demo-frontend:latest .

# Run the container:
# port mapping: 5173:80 (host:container)
frontend $ docker run -p 5173:80 data-visualization-demo-frontend:latest
```

The frontend now is running on http://localhost:5173
