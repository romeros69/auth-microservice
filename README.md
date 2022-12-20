# auth-microservice
auth-microservice for cloud

### Как запустить (из корня проекта):
docker-compose up --build

### SWAGGER UI:
http://localhost:9000/swagger/index.html#

###
scp -i key dev/golang/cloud/auth-microservice/cmd/main/main romich-v2@51.250.108.95:~/


###
GOOS=linux GOARCH=amd64 go build
