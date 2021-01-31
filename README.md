# Web Api Tester
Web API Tester is a tool that can be used to test multiples web applications. In it you can register many HTTP requests and execute then to test your application health.

## :warning: Disclaimer :warning:
Please be careful to not use this against application tat you do not have a permission to do so. Since it can be used to simulate many requests or trigger a sensive operation at the target.

## Run
To run the application you have to follow these steps:

### 1. First you'll need to create a configuration file on `config/`
This is a simple `yaml` file that will indicate the port to be used and point to the database file that will be used **(if you don't have one don't worry, the application will create one for you)**, for example:
```yaml
# config/config.yaml
server:
  port: 8080

database:
  filename: config/database.db
```

Here how you config folder should look like:

```console
┌─[<user>@<machine>] - [~/web-api-tester]
└─[$] <> ls
config.yaml
```

### 2. Then you can choose two ways to run the application
To run the application you have two options.

#### Local
To run things locally you can just run

```console
┌─[<user>@<machine>] - [~/web-api-tester]
└─[$] <> go build

┌─[<user>@<machine>] - [~/web-api-tester]
└─[$] <> ./web-api-tester
2021/01/30 23:54:56 Starting to serve at port 8080...
```

#### Docker
To execute through docker first you'll need to execute the `build.sh` script to create the image:
```console
┌─[<user>@<machine>] - [~/web-api-tester]
└─[$] <> ./build.sh
```

Then you run the docker imagem mounting the config folder that you create at the step one
```console
┌─[<user>@<machine>] - [~/web-api-tester]
└─[$] <> docker run -d -p <port>:<port> -v "$PWD"/config:/config --name <name> web-api-tester
```

## Test
To test the application you can execute `test.sh` script at the root folder. Example
```console
┌─[<user>@<machine>] - [~/web-api-tester]
└─[$] <> ./test.sh
```
