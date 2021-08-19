# Ozon Code Platform Certificate API

OCP Certificate Api - service for work and management of the certificate platform.

###The service supports management methods:

| Field | Type | Description |
| ------ | ------ | ------ |
| Id | Number | Unique id certificate |
| UserId | Number | Id user to whom the certificate belongs |
| Created | Timestamp | Certificate creation time |
| Link | String | Link to certificate |

###To start and build the service, you need to do the following

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
###OR
#### All steps [Build+Deploy+Run+Migrate]
```sh 
    make deploy
```
-----
##Stop container
```sh 
    make stop
```
##Remove containers and images
```sh 
    docker rm -vf $(docker ps -a -q)
    docker rmi -f $(docker images -a -q)
```