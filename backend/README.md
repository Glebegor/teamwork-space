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
api/v1/teams/        GET<br>
api/v1/teams/:id     GET<br>

api/v2/profile/      GET<br>
api/v2/profile/:id   GET<br>

#### api/v2/;
api/v2/teams/        POST<br>
api/v2/teams/:id     PUT<br>
api/v2/teams/:id     DELETE<br>

api/v2/profile/:id   PUT<br>
api/v2/profile/:id   DELETE<br>

api/v2/teamTasks/    POST<br>
api/v2/teamTasks/    GET<br>
api/v2/teamTasks/:id GET<br>
api/v2/teamTasks/:id PUT<br>
api/v2/teamTasks/:id DELETE<br>

#### api/v3/;

### Logic

## Enviroment

## Entityes

Roles:
<ul>
<li>Admin</li>
<li>User</li>
<li>Subscriber</li>
</ul>