# Curio QA API

## Run Manually Without docker-compose

### Curio QA

Build Docker image
```
docker build -f Dockerfile.dev --tag curio_qa_api:1.0 .
```

Create Container
```
docker container create --name curio_qa_api -p 8080:8080 curio_qa_api:1.0
```

Run Container
```
docker container start curio_qa_api
```

See Logs
```
docker container logs curio_qa_api
```

### Databases

Create docker volumes
```
docker volume create mongo_data
docker volume create mongo_config
docker volume create redis_cache
```

Create and run MongoDB with created volumes
```
docker container create --name mongo -p -v 27017:27017 mongo_data:/data/db mongo_config:/data/configdb mongo:7.0.3

docker run -d --name mongo \
    -v mongo_data:/data/db \
    -v mongo_config:/data/configdb \
    -p 27017:27017 \
    mongo:7.0.3
```

Create and run redis with created volumes
```
docker container create --name redis -p -v 6379:6379 redis_cache:/data redis:7.2.3

docker run -d --name redis \
    -v redis_cache:/data \
    -p 6379:6379 \
    redis:7.2.3
```

### Network

Create a network
```
docker network create curio_qa_api_network
```

Connect existing containers to the created network
```
docker network connect curio_qa_api_network mongo
docker network connect curio_qa_api_network redis
docker network connect curio_qa_api_network curio_qa_api
```
