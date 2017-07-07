# Another option for deploying the application is by creating a binary
# Since AWS is running linux, we need to create a linux executable for go
GOARCH=amd64 GOOS=linux go build -o bin/application application.go
