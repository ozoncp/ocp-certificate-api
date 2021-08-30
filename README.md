# Ozon Code Platform Certificate API

[![build-and-test](https://github.com/ozoncp/ocp-certificate-api/actions/workflows/build-and-test.yml/badge.svg?branch=main)](https://github.com/ozoncp/ocp-certificate-api/actions/workflows/build-and-test.yml)
[![codecov](https://codecov.io/gh/ozoncp/ocp-certificate-api/branch/main/graph/badge.svg?token=2649a463-f405-4622-8624-c91aa3dd7d5f)](https://codecov.io/gh/ozoncp/ocp-certificate-api)

OCP Certificate Api - service for work and management of the certificate platform.

### The service supports management methods:

| Field | Type | Description |
| ------ | ------ | ------ |
| Id | Number | Unique id certificate |
| UserId | Number | Id user to whom the certificate belongs |
| Created | Timestamp | Certificate creation time |
| Link | String | Link to certificate |

### To start and build the service, you need to do the following

#### 1 Step [Clone]
- git clone https://github.com/ozoncp/ocp-certificate-api.git
- cd ocp-certificate-api
#### 2 Step [Dependence]
```sh 
    make deps
```
#### 3 Step [Build]
```sh 
    make build
```
#### 4 Step [Run]
```sh 
    make start
```
#### 5 Step [Migrate]
```sh 
    make migrate
```
- or run in container cli interface
```sh 
    ./ocp-certificate-api -migrate up
    ./ocp-certificate-api -migrate down
```
### OR
#### All steps [Build+Deploy+Run+Migrate]
```sh 
    make deploy
```
-----
### Stop container
```sh 
    make stop
```
### Remove containers and images
```sh 
    docker rm -vf $(docker ps -a -q)
    docker rmi -f $(docker images -a -q)
```
-----
## Load testing with Pandora:
### run:
```sh 
 test/load/load pandora.yaml
```
### set config:
```sh 
 pandora.yaml
```
-----

## Services:
### [Swagger UI](http://localhost:8080)
- http://localhost:8080
### [REST](http://localhost:8081)
- http://localhost:8081/v1/{methods}
### [gRPC](http://localhost:8082)
- http://localhost:8082
### [Prometheus](http://localhost:9090)
- http://localhost:9090
### [Grafana](http://localhost:3000)
- http://localhost:3000
- Auth
  - admin/MYPASSWORT
- Set dashboard:
  - Configuration -> DataSources -> Add datasource -> Prometheus
  - Set Url: http://localhost:9090
  - Save & test
  - Run to Explore page
### [Metrics](http://localhost:9100/metrics)
- http://localhost:9100/metrics
### [Jaeger UI](http://localhost:16686)
- http://localhost:16686
### [Graylog](http://localhost:9000)
- http://localhost:9000
- Auth
  - admin/admin
### [Kafka](http://localhost:9094)
- http://localhost:9094
### [Kafka UI](http://localhost:9001)
- http://localhost:9001
### [Zookeeper](http://localhost:2181)
- http://localhost:2181