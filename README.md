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
#### 2 Step [Deploy]
```sh 
    make deploy
```