# Nginx Configuration on a Linux Server

Use Nginx as a reverse proxy to serve a TypeScript/React frontend and a Go/Gin backend on the same domain.

## 1. Create a new file data-visualization-demo in the /etc/nginx/sites-available directory:
```bash
$ sudo vim /etc/nginx/sites-available/data-visualization-demo
```

## 2. Add the following configuration:
```conf
# /etc/nginx/sites-available/data-visualization-demo

# Define a server block.
server {
    # Listen on port 80 (default HTTP port).
    listen 80;

    # Define the domain name.
    server_name data-visualization-demo.qingquanli.com csv-data-visualization.qingquanli.com;

    # Define the location for the API endpoint (Go/Gin backend).
    location /api {
        # Deploy with Docker Compose. 8003 is set in the docker-compose.prod.yml.
        # proxy_pass http://localhost:8003;
        # Deploy with Kubernetes. 30010 is the Kubernetes nodePort set in a Kubernetes yaml file.
        proxy_pass http://localhost:30010;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }

    # Define the location for the root URL (TypeScript/React frontend).
    location / {
        # Deploy with Docker Compose. 5174 is set in the docker-compose.prod.yml.
        # proxy_pass http://localhost:5174;
        # Deploy with Kubernetes. 30011 is the Kubernetes nodePort set in a Kubernetes yaml file.
        proxy_pass http://localhost:30011;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}
```

## 3. Create a symbolic link to the file in the /etc/nginx/sites-enabled directory:
```bash
$ sudo ln -s /etc/nginx/sites-available/data-visualization-demo /etc/nginx/sites-enabled/
```


## 4. Test the Nginx configuration:
```bash
$ sudo nginx -t
```


## 5. Reload Nginx to apply the changes:
```bash
$ sudo systemctl reload nginx
```


## 6. Optional: Configure HTTPS

- Use Cloudflare DNS proxy to enable HTTPS (it can also enable to hide the IP address of the server)
- Or: Configure Let's Encrypt (certbot) in Nginx to enable HTTPS
