# Deployment Guide

This guide covers deploying Todo App to various environments.

---

## Prerequisites

- Docker & Docker Compose (for containerized deployment)
- Kubernetes cluster (for K8s deployment)
- MongoDB Atlas or MongoDB instance
- Domain name (for production)

---

## Local Development Deployment

### Using Make
```bash
# Build and run
make run

# Or with Docker Compose
make docker-up

# Stop
make docker-down
```

### Manual Steps
```bash
# 1. Install dependencies
go mod download

# 2. Build binary
go build -o server ./cmd/server/main.go

# 3. Set environment variables
export PORT=8080
export MONGO_URI=mongodb://localhost:27017
export DATABASE_NAME=todos
export JWT_SECRET=dev-secret

# 4. Run
./server
```

---

## Docker Deployment

### Build Docker Image
```bash
make docker-build
```

### Run Container
```bash
# With Docker Compose (recommended)
make docker-up

# Or with docker run
docker run -d \
  -p 8080:8080 \
  -e MONGO_URI=mongodb://mongo:27017 \
  -e DATABASE_NAME=todos \
  -e JWT_SECRET=your-secret \
  --name todo-app \
  todo-app:latest
```

### Push to Docker Registry
```bash
# Tag image
docker tag todo-app:latest myregistry/todo-app:1.0.0

# Push
docker push myregistry/todo-app:1.0.0
```

---

## Kubernetes Deployment

### Prerequisites
```bash
# Install kubectl
# Install helm (optional)
# Have a Kubernetes cluster running
```

### Deploy Using Kubectl

#### 1. Create Namespace
```bash
kubectl create namespace todo-app
```

#### 2. Create ConfigMap
```bash
kubectl create configmap todo-config \
  --from-literal=DATABASE_NAME=todos \
  --from-literal=PORT=8080 \
  -n todo-app
```

#### 3. Create Secret
```bash
kubectl create secret generic todo-secrets \
  --from-literal=MONGO_URI=mongodb://... \
  --from-literal=JWT_SECRET=$(openssl rand -base64 32) \
  -n todo-app
```

#### 4. Deploy Application
```bash
# Using provided manifests
kubectl apply -f deployments/kubernetes/ -n todo-app

# Or deploy step by step
kubectl apply -f deployments/kubernetes/deployment.yaml -n todo-app
kubectl apply -f deployments/kubernetes/service.yaml -n todo-app
```

### Example deployment.yaml
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: todo-app
  namespace: todo-app
spec:
  replicas: 3
  selector:
    matchLabels:
      app: todo-app
  template:
    metadata:
      labels:
        app: todo-app
    spec:
      containers:
      - name: app
        image: myregistry/todo-app:1.0.0
        ports:
        - containerPort: 8080
        env:
        - name: PORT
          value: "8080"
        - name: DATABASE_NAME
          valueFrom:
            configMapKeyRef:
              name: todo-config
              key: DATABASE_NAME
        - name: MONGO_URI
          valueFrom:
            secretKeyRef:
              name: todo-secrets
              key: MONGO_URI
        - name: JWT_SECRET
          valueFrom:
            secretKeyRef:
              name: todo-secrets
              key: JWT_SECRET
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
        resources:
          requests:
            memory: "64Mi"
            cpu: "100m"
          limits:
            memory: "256Mi"
            cpu: "500m"
---
apiVersion: v1
kind: Service
metadata:
  name: todo-app
  namespace: todo-app
spec:
  type: LoadBalancer
  ports:
  - port: 80
    targetPort: 8080
  selector:
    app: todo-app
```

### Check Deployment Status
```bash
# View pods
kubectl get pods -n todo-app

# View logs
kubectl logs -f pod/todo-app-xxx -n todo-app

# View service
kubectl get svc -n todo-app

# Scale deployment
kubectl scale deployment todo-app --replicas=5 -n todo-app
```

---

## Cloud Deployment

### AWS ECS (Elastic Container Service)

```bash
# Build and push image
docker build -t myregistry/todo-app:latest .
docker push myregistry/todo-app:latest

# Create ECS task definition (todo-app-task.json)
# Update with your image URI and environment variables

# Register task definition
aws ecs register-task-definition \
  --cli-input-json file://todo-app-task.json

# Create service
aws ecs create-service \
  --cluster my-cluster \
  --service-name todo-app \
  --task-definition todo-app:1 \
  --desired-count 3 \
  --load-balancers targetGroupArn=arn:aws:...,containerName=app,containerPort=8080
```

### Heroku

```bash
# Install Heroku CLI
# heroku login

# Create app
heroku create my-todo-app

# Set environment variables
heroku config:set \
  MONGO_URI=... \
  JWT_SECRET=$(openssl rand -base64 32) \
  -a my-todo-app

# Deploy
git push heroku main
```

### Google Cloud Run

```bash
# Build image
gcloud builds submit --tag gcr.io/my-project/todo-app

# Deploy to Cloud Run
gcloud run deploy todo-app \
  --image gcr.io/my-project/todo-app \
  --platform managed \
  --region us-central1 \
  --memory 512Mi \
  --set-env-vars MONGO_URI=... \
  --set-env-vars JWT_SECRET=...
```

---

## Database Setup

### MongoDB Atlas (Cloud)

```bash
# 1. Create cluster at mongodb.com/atlas
# 2. Create database user
# 3. Whitelist IP addresses
# 4. Copy connection string

# 4. Set MONGO_URI environment variable
MONGO_URI=mongodb+srv://user:password@cluster.mongodb.net/todo-app
```

### Self-Hosted MongoDB

```bash
# Using Docker
docker run -d \
  -e MONGO_INITDB_ROOT_USERNAME=admin \
  -e MONGO_INITDB_ROOT_PASSWORD=password \
  -p 27017:27017 \
  mongo:latest

# Or use Docker Compose in the project
docker-compose up -d mongo
```

### Create Indexes
```bash
# Connect to MongoDB
mongosh

# Switch to database
use todos

# Create indexes
db.todos.createIndex({ "user_id": 1 })
db.todos.createIndex({ "created_at": -1 })
db.users.createIndex({ "email": 1 }, { unique: true })
```

---

## SSL/TLS Configuration

### Using Let's Encrypt with Nginx

```nginx
server {
    listen 443 ssl;
    server_name api.example.com;

    ssl_certificate /etc/letsencrypt/live/api.example.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/api.example.com/privkey.pem;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

### In Kubernetes
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: todo-ingress
  annotations:
    cert-manager.io/cluster-issuer: "letsencrypt-prod"
spec:
  ingressClassName: nginx
  tls:
  - hosts:
    - api.example.com
    secretName: todo-tls
  rules:
  - host: api.example.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: todo-app
            port:
              number: 80
```

---

## Monitoring & Logging

### Health Check
```bash
curl http://localhost:8080/health
```

### View Logs
```bash
# Docker
docker logs -f todo-app

# Kubernetes
kubectl logs -f deployment/todo-app -n todo-app

# Direct application
# Logs are written to stdout and can be captured
```

### Set Up Monitoring
```bash
# Using Prometheus (example)
docker run -d \
  -p 9090:9090 \
  -v prometheus.yml:/etc/prometheus/prometheus.yml \
  prom/prometheus

# Scrape metrics from /metrics endpoint
```

---

## CI/CD Pipeline

### GitHub Actions Example

```yaml
name: Build and Deploy

on:
  push:
    branches: [main]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    
    - name: Set up Go
      uses: actions/setup-go@v2
      with:
        go-version: 1.25
    
    - name: Run tests
      run: make test
    
    - name: Build Docker image
      run: docker build -t myregistry/todo-app:${{ github.sha }} .
    
    - name: Push to registry
      run: |
        echo ${{ secrets.REGISTRY_PASSWORD }} | docker login -u ${{ secrets.REGISTRY_USER }} --password-stdin
        docker push myregistry/todo-app:${{ github.sha }}
    
    - name: Deploy to production
      run: |
        kubectl set image deployment/todo-app \
          app=myregistry/todo-app:${{ github.sha }} \
          -n todo-app
```

---

## Rollback Procedure

### Docker
```bash
# Run previous version
docker run -d \
  -p 8080:8080 \
  myregistry/todo-app:previous-tag
```

### Kubernetes
```bash
# View rollout history
kubectl rollout history deployment/todo-app -n todo-app

# Rollback to previous version
kubectl rollout undo deployment/todo-app -n todo-app

# Rollback to specific revision
kubectl rollout undo deployment/todo-app --to-revision=2 -n todo-app
```

---

## Backup Strategy

### Database Backup
```bash
# MongoDB Atlas - automated backups available
# Self-hosted MongoDB:
mongodump --uri="mongodb://..." --out=./backup

# Restore
mongorestore ./backup
```

### Backup Schedule
- Daily backups
- Retain for 30 days
- Test restore procedure monthly

---

## Performance Tuning

### MongoDB
- Create indexes on frequently queried fields
- Use connection pooling
- Monitor query performance

### Application
- Implement caching (Redis)
- Use load balancing
- Monitor resource usage

### Infrastructure
- Scale horizontally (multiple instances)
- Use CDN for static content
- Optimize network latency

---

## Troubleshooting Deployment

### Application won't start
```bash
# Check logs
docker logs todo-app

# Verify environment variables
docker inspect todo-app | grep -A 20 Env

# Check port availability
netstat -an | grep 8080
```

### Database connection failed
```bash
# Verify MONGO_URI
echo $MONGO_URI

# Test connection
mongosh $MONGO_URI

# Check network connectivity
nc -zv mongo-host 27017
```

### Health check failing
```bash
# Test health endpoint
curl -v http://localhost:8080/health

# Check database connectivity from app
```

---

## Maintenance Tasks

### Regular Tasks
- Monitor error rates
- Review logs
- Update dependencies
- Test backup/restore
- Update certificates

### Quarterly Tasks
- Performance analysis
- Security audit
- Capacity planning
- Update documentation

---

**Last Updated:** January 20, 2026
