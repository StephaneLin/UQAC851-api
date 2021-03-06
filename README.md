[![Integration Backend](https://github.com/ethicnology/uqac-851-software-engineering-api/actions/workflows/continuous-integration.yaml/badge.svg)](https://github.com/ethicnology/uqac-851-software-engineering-api/actions/workflows/continuous-integration.yaml)
[![Deployment Backend](https://github.com/ethicnology/uqac-851-software-engineering-api/actions/workflows/continuous-deployment.yaml/badge.svg)](https://github.com/ethicnology/uqac-851-software-engineering-api/actions/workflows/continuous-deployment.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ethicnology/uqac-851-software-engineering-api)](https://goreportcard.com/report/github.com/ethicnology/uqac-851-software-engineering-api)
[![BCH compliance](https://bettercodehub.com/edge/badge/ethicnology/uqac-851-software-engineering-api?branch=develop)](https://bettercodehub.com/)


## About The Project
![](https://github.com/ethicnology/uqac-851-software-engineering-api/blob/develop/docs/logo.png "Screenshot")  
This is a school project.  
All the specifications are specified in [docs/Projet_Pratique.pdf](https://github.com/ethicnology/uqac-851-software-engineering-api/blob/develop/docs/Projet_Pratique.pdf)

* [Docker](https://www.docker.com)
* [Golang](https://golang.org)  
* [Goyave](https://goyave.dev) 
* [MariaDB](https://mariadb.org) 

### :open_file_folder: Directory Structure
```
.
├── docs
|
├── database
│   ├── model                // ORM models and Generators for database testing
│       └── ...
├── http
│   ├── controller           // Business logic of the application
│   │   └── ...
│   ├── middleware           // Logic executed before or after controllers
│   │   └── ...
│   └── route
│       └── route.go         // Routes definition
│
├── test                     // Unit and Functional tests
|   └── func_test.go
|   └── unit_test.go
|
├── .gitignore
├── config.example.json      // Example config for local development
├── config.test.json         // Config file used for tests
├── docker-compose.yml       // Build local architecture file used for tests
├── docker-compose.test.yml  // Build architecture to execute tests
├── Dockerfile               // API Dockerfile
├── go.mod
└── main.go                  // Application entrypoint
```


## :rocket: Getting Started
To get a local copy up and running follow these simple example steps.
### :page_with_curl: Prerequisites
All the project architecture is dockerized : [Docker](https://www.docker.com/products/docker-desktop).  
You only need docker installed in your OS in order to host all the components without any extra-installation.  

### :construction_worker: Installation
#### Clone
```sh
git clone https://github.com/ethicnology/uqac-851-software-engineering-api.git
```
#### Configuration
Copy **config.example.json** as **config.json** and add **the SMTP password (ask to students)**.  
```sh
cp config.example.json config.json
```
If you want to run the tests, you need to **specify the SMTP password in the config.test.json too** :
```json
{ ...
      "prix-banque":{
        "smtp_host":"smtp.gmail.com",
        "smtp_port":"587",
        "smtp_user":"prixbanque@gmail.com",
        "smtp_pass":"" //Ask Students For The Password
    }
}
```


### :whale: Run with docker
```sh
docker-compose up
# Available on 172.x.x.x:1984
```

### :pray: Tests
```sh
docker-compose -f docker-compose.test.yml up --abort-on-container-exit --remove-orphans
# Exit when tests are finished
```
### :books: Documentation
I built a website with Github Pages which contains [**API documentation**](https://ethicnology.github.io/uqac-851-software-engineering-api/)  
Also, you can find a markdwown version in [/docs/README.md](https://github.com/ethicnology/uqac-851-software-engineering-api/tree/develop/docs#readme)

### :runner: Usage
#### Postman
To play with your API, import the collection in /docs in [Postman](https://www.postman.com/).  
All available endpoints are documented with examples.

#### cURL
Instead you can try it with your CLI using **cURL**.  
On following examples URL is : **https://dissidence.dev:9999** you can change this by your docker IP using **http** without ssl

##### Windows
Windows force you to mask doublequotes with backslash like this : 
```powershell
curl -d "{\"email\":\"existing@email.pls\",\"password\":\"Str0ng\",\"first_name\":\"Paul\",\"last_name\":\"Lefevbre\"}" -H "Content-Type: application/json" -X POST https://dissidence.dev:9999/auth/register
```
##### Linux
```sh
curl -d '{"email":"existing@email.pls","password":"Str0ng","first_name":"Paul","last_name":"Lefevbre"}' -H 'Content-Type: application/json' -X POST https://dissidence.dev:9999/auth/register
# {"id":9}
```

```sh
curl -d '{"email":"existing@email.pls","password":"Str0ng"}' -H 'Content-Type: application/json' -X POST https://dissidence.dev:9999/auth/login
# {"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTM0NzkyNDQsIm5iZiI6MTYxNzQ3OTI0NCwidXNlcmlkIjoic2Vuc2VpQHVxYWMuY2EifQ.aMRWeebCTfJyUPfsUz5H8Ng1x1L1T10hSKpXoVdyPUY"}
```

```sh
TOKEN=$YourToken;curl -H 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" -X GET https://dissidence.dev:9999/users/existing@email.pls
# {"created_at":"2021-04-03T19:47:19.718Z","updated_at":"2021-04-03T19:47:19.718Z","deleted_at":null,"id":10,"email":"existing@email.pls","first_name":"Paul","last_name":"Lefebvre"}
```

### :hammer: Build from scratch
#### Prerequisites
If you want to build the project from scratch without docker you need to install few tools.
* Go 1.13+
* mariadb-server

Once you've installed Go and MariaDB.
You need to create a MariaDB user and a database:
```sql
mysql> CREATE USER 'goyave'@'localhost' IDENTIFIED BY 'secret';
mysql> CREATE DATABASE goyave;
mysql> GRANT ALL PRIVILEGES ON 'goyave'.* TO 'goyave'@'localhost';
mysql> FLUSH PRIVILEGES;
```
You can change theses values but you will need to update **config.json**.

#### Build
From root directory, execute :
```sh
go build
```
It will output an executable which have the same name as the project, make it executable :
```sh
chmod +X uqac-851-software-engineering-api
```
Then, run :
```sh
./uqac-851-software-engineering-api
```
## :mag: Contact
[@ethicnology](https://twitter.com/ethicnology)  
Project Link: [https://github.com/851-software-engineering](https://github.com/851-software-engineering)
