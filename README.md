# dyrector.io assesment

You received a task of containerizing an application.

## Introduction

Application components:
 - service
 - web-client

MySQL is part of the infrastructure, but not the repository, since it not maintained by us.


### Service

A Golang simple CRUD service with a fuzzy-get feature to generate get cat ASCII arts. The service uses MySQL as persistence, connectivity is configured via environment variables.

Configuration

| Environment variable | Description              | Default   |
| -------------------- | ------------------------ | --------- |
| MYSQL_HOST           | Database host to connect | localhost |
| MYSQL_PORT           | Database port to use     | 3306      |
| MYSQL_USER           | Database user            | -         |
| MYSQL_PASSWORD       | Database password        | -         |
| MYSQL_DATABASE       | Database name            | catCRUD   |

For simplicity CORS requests are allowed without restriction.

In order to compile and run the service locally use the official to install go
https://go.dev/learn

For the service to work a configured MySQL is necessary (see tasks).

Run the service 
``` sh
# in the service directory
go run .
```


Compile the service
```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-extldflags '-static'" -o build/cats
```

In a linux environment you can now run the freshly created `cats` executable.


### Web client

The web client expects a name to calculate a matching cat.
It is built using HTML and VanillaJS.
It generates cats via the standard browser `fetch` call dynamically.

In the `web/assets/script.js` contains the service communaction data (url, port),
feel free to change it if need be.

One way to run/serve it
```
python -m http.server 3000
```
This command makes it available on port 3000, if the API is running it should work.
## Tasks

After creating the repository, your commits are also evaluated.
If you have completed a task create a commit in a format: `<task number>: <descriptive message>`

Example commit log:
```
3: Dockerfile created
4: cat-service added to compose
3: Extended Dockerfile 
```

1. Make a Github repository and upload all your work
2. Set up a MySQL service (hint: containerization) using docker-compose
3. Create a Dockerfile for the  `cats-service` (using Google is advised)
4. Add a cat-service to the docker-compose.yml file 
5. Add nginx to the docker-compose.yml
6. Serve static file content using using the previously added nginx on the 60080 port
7. Configure the web-server to proxy the request towards the API on the `/api/` subpath
8. Change URL in `script.js` where it is needed


If the instructions are unclear feel free to reach out with questions.

Please track your time and make clear and simple statistics on how much time did you spend on which task. Write the statistic to a new `Markdown` file eg.: `worklogs.md`.

Example track log:
```
3: 2h 20m - Investigate the Dockerfile
4: 3h 10m - Learn how docker-compose works
4: 1h 30m - Implement and test the new .yml file
```

## Optional tasks
For bonus points you can:
- add new cats into the db seed script `cat-service/db.go`
- enlist some extension ideas into repository root `ideas.md`
- create a dogs endpoint and ui section
