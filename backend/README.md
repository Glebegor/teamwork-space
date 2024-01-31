---
runme:
  id: 01HNDRCWH3EHTAVFV24HV239WW
  version: v2.2
---

# Backend

## Description

<p>
That is the backend that based on clean architecture and using REST and GRPC to communicate with clients or that provide open api for another server.  
</p>

## API

### Routes

<p>api/v1/ - public api routes</p>
<p>api/v2/ - protected api routes by jwt</p>
<p>api/v3/ - open api routes</p>

#### api/v1/;

api/v1/teams/        GET
api/v1/teams/:id     GET

api/v2/profile/      GET
api/v2/profile/:id   GET

#### api/v2/;

api/v2/teams/        POST
api/v2/teams/:id     PUT
api/v2/teams/:id     DELETE

api/v2/profile/:id   PUT
api/v2/profile/:id   DELETE

api/v2/teamTasks/    POST
api/v2/teamTasks/    GET
api/v2/teamTasks/:id GET
api/v2/teamTasks/:id PUT
api/v2/teamTasks/:id DELETE

#### api/v3/;

### Logic

## Enviroment

### Env

```sh {"id":"01HNFM6NAE7YRMZQ68ZP474DNV"}
DB_PASSWORD="123321"
SECRET="key"
```

### Config

```yaml {"id":"01HNFM6NAE7YRMZQ68ZQQMAS8Z"}
server:
  ENV:  "development"
  PORT: 8080

db:
  PORT:     5436
  HOST:     "localhost"
  NAME:     "team-work-space"
  SSLMODE:  "disabled"
  USERNAME: "admin"
```

## Entityes

### Domain

#### User

#### Team

#### Profile

#### Team tasks

Roles:

<ul>
<li>Admin</li>
<li>User</li>
<li>Subscriber</li>
</ul>