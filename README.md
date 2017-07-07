# go-aws-beanstalk


Zip the file contents and just upload it to AWS Beanstalk. 
You can choose to upload it to either the `worker` or `web-server` environment.

For `worker` environment, the `cron.yaml` will take into effect.

Remember to run `chmod +x build.sh`.

Carry out cross-compile to linux binary:

```bash
$ GOARCH=amd64 GOOS=linux go build -o bin/application application.go
```