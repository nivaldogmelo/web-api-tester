# Web Api Tester
Web API Tester is a tool that can be used to test multiples web applications. In it you can register many HTTP requests and execute then to test your application health.

## :warning: Disclaimer :warning:
Please be careful to not use this against application tat you do not have a permission to do so. Since it can be used to simulate many requests or trigger a sensive operation at the target.

## Usage
To register requests you'll need to make a http requests as the one below
```json
# http://<url>:<port>/
# Method: POST
{
  "Name":  "myServer",
  "Method":  "GET",
  "Headers":  [
    {
      "Header":  "Content-Type",
      "Content":  "application/json"
    }
  ],
  "Body":  "",
  "URL":  "http://myserver:3000/"
}
```
Then the app will it include in the next round of tests. The application will export metrics about the tests in the Prometheus format at `/metrics` as below
```
# HELP web_requests_latency Latency of registered web requests
# TYPE web_requests_latency summary
web_requests_latency{code="200",name="myServer",quantile="0.5"} 0.002459737
web_requests_latency{code="200",name="myServer",quantile="0.9"} 0.005159327
web_requests_latency{code="200",name="myServer",quantile="0.99"} 0.005159327
web_requests_latency_sum{code="200",name="myServer"} 0.027415872
web_requests_latency_count{code="200",name="myServer"} 8
# HELP web_requests_total How many registered web requests were made, partitioned by name and result
# TYPE web_requests_total counter
web_requests_total{code="200",name="myServer"} 8
```

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

```bash
┌─[<user>@<machine>] - [~/web-api-tester]
└─[$] <> ls
config.yaml
```

### 2. Then you can choose two ways to run the application
To run the application you have two options.

#### Local
To run things locally you can just run

```bash
┌─[<user>@<machine>] - [~/web-api-tester]
└─[$] <> go build

┌─[<user>@<machine>] - [~/web-api-tester]
└─[$] <> ./web-api-tester
2021/01/30 23:54:56 Starting to serve at port 8080...
```

#### Docker
To execute through docker first you'll need to execute the `build.sh` script to create the image:
```bash
┌─[<user>@<machine>] - [~/web-api-tester]
└─[$] <> ./build.sh
```

Then you run the docker imagem mounting the config folder that you create at the step one
```bash
┌─[<user>@<machine>] - [~/web-api-tester]
└─[$] <> docker run -d -p <port>:<port> -v "$PWD"/config:/config --name <name> web-api-tester
```

## Test
To test the application you can execute `test.sh` script at the root folder. Example
```bash
┌─[<user>@<machine>] - [~/web-api-tester]
└─[$] <> ./test.sh
```
