# Docker Local Environment

- AWS (localstack)

  ```
  aws configure --profile default
  aws --endpoint-url=http://localhost:4566 s3 ls
  ```

- DB (MySQL / MariaDB)

  ```
  mysql -u username -p dbname
  ```

- MongoDB

  ```
  mongo -u username -p password localhost:27017/test
  ```

- Redis

  ```
  redis-cli -h localhost -p 6379
  ```
