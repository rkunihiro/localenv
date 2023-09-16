# Docker Local Environment

Local development environment with Docker Compose.

- AWS (localstack)
- DB (MySQL / MariaDB)
- MongoDB
- Redis

## Usage

### Start services

```sh
docker compose up -d && docker compose logs -f
```

### AWS (localstack)

```sh
docker compose exec aws bash
```

```sh
awslocal s3 ls
```

### DB (MySQL / MariaDB)

```sh
docker compose exec db mysql -u username -p dbname
> Enter password: password
```

### MongoDB

Connect from host

```sh
mongo -u username -p password localhost:27017/test
```

### Redis

Monitor

```sh
docker compose exec redis redis-cli monitor
```

Connect from host

```sh
redis-cli -h localhost -p 6379
```
