# go-aws-beanstalk


Zip the file contents and just upload it to AWS Beanstalk. 
You can choose to upload it to either the `worker` or `web-server` environment.

For `worker` environment, the `cron.yaml` will take into effect.

Remember to run `chmod +x build.sh`.

Carry out cross-compile to linux binary:

```bash
$ GOARCH=amd64 GOOS=linux go build -o bin/application application.go
```

## Folder structure

- __application.go__: The entry point of your application
- __build.sh__: -
- __Buildfile__: -
- __cron.yaml__: -
- __Procfile__: A Procfile is a mechanism for declaring what commands are run by your application's dynos on the Heroku platform. It follows the process model. You can use a Procfile to declare various process types, such as multiple types of workers, a singleton process like a clock, or a consumer of the Twitter streaming API
- __setup.sh__: -
- __.ebextensions__: -


## Limitation

- `go get ./...` does not work
- AWS Beanstalk only supports up to Go v1.7 (11 July 2017). Some packages that uses context might not work.
- AWS timezone always default to UTC.