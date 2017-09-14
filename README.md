# go-cms-api

This is a small api service provides getting, creating posts via REST. It also includes the JWT authentication to prevent anonymous put the data into database.

## Table of contents

[How to use](https://github.com/tatthien/go-cms-api#how-to-use)
- [Get the project](https://github.com/tatthien/go-cms-api#get-the-project)
- [Configure app information](https://github.com/tatthien/go-cms-api#configure-app-information)
- [Configure token](https://github.com/tatthien/go-cms-api#configure-token)
- [Install database](https://github.com/tatthien/go-cms-api#install-database)
- [Run in local](https://github.com/tatthien/go-cms-api#run-in-local)
- [Run in server with custom domain](https://github.com/tatthien/go-cms-api#run-in-server-with-custom-domain)

## How to use

### Get the project 

Assume that you've already configured your GOPATH.

```
go get github.com/tatthien/go-cms-api
```


### Configure app information
Copy or rename the file `.env.example` to `.env`. Then enter your own server ip, application port, database information.

```
APP_IP=your_server_ip
APP_PORT=your_application_port

DATABASE_NAME=your_database_name
DATABASE_USERNAME=your_datatase_username
DATABASE_PASSWORD=your_database_password

ADMIN_USERNAME=your_admin_username
ADMIN_PASSWORD=your_admin_password
ADMIN_EMAIL=your_admin_email
```

If you folk this project, make sure you do not push this `.env` file to Github. I think you know why ;)

### Configure token

This project is using `jwt-golang` package to generate and validate the token.

To use the token you need to create folder `keys` then put your key files into that folder.

**Create folder and generate key files**

```
mkdir keys
openssl genrsa -out app.rsa 1024
openssl rsa -in app.rsa -pubout -out app.rsa.pub
```

### Install database

```
go run install/install.go
```

### Run in local

```
go run go-cms-api.go
```

### Run in server with custom domain

**Build and run in background**

```
go build go-cms-api.go
nohup ./go-cms-api &
```

**Configure nignx proxy**

Assume that you run the application on port `3000`.

Follow these steps below:

Open config file in `vim`
```
vi yourdomain.com.conf
```

Add this code at the bottom of the file.

```
 location /app/ {
  proxy_set_header X-Real-IP $remote_addr;
  proxy_set_header X-Forwarded-For $remote_addr;
  proxy_set_header Host $host;
  proxy_pass http://127.0.0.1:3000/;
}
``` 

Restart nginx

```
systemctl restart nginx
```

After the configuration, you can access your api at `yourdomain.com/app/api/v1`
