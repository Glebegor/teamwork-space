# Backend

## Description
<p>That is the backend that based on clean architecture and using REST and GRPC to communicate with clients or that provide open api for another server.</p>

## API

### Tokens
<p>Using JWT with claims. Based on refresh access technology.</p>

#### Access claims

```
type JwtClaimsAccess struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	jwt.StandardClaims // ExpiresAt number `json:"exp"`
}
```

#### Refresh claims

```
type JwtClaimsRefresh struct {
	ID string `json:"id"`
	jwt.StandardClaims // ExpiresAt number `json:"exp"`
}
```

### Responses
<p>Error Response: { "Message": "Some info."}</p>
<p>Success Response: { "Status": "ok"}</p>

### Routes
<p>api/v1/ - public api routes</p>
<p>api/v2/ - protected api routes by jwt</p>
<p>api/v3/ - open api routes</p>

#### api/v1/;
api/v1/auth/login POST<br>
api/v1/auth/registration POST<br>
api/v1/auth/refresh POST<br>

<!-- api/v1/teams/        GET<br>
api/v1/teams/:id     GET<br>

api/v1/users/      GET<br>
api/v1/users/getById/:id   GET<br>
api/v1/users/getByEmail/:email   GET<br> -->

#### api/v2/;

<!-- api/v2/teams/        POST<br>
api/v2/teams/:id     PUT<br>
api/v2/teams/:id     DELETE<br>

api/v2/users/:id     PUT<br>
api/v2/users/:id     DELETE<br>

api/v2/teamTasks/    POST<br>
api/v2/teamTasks/    GET<br>
api/v2/teamTasks/:id GET<br>
api/v2/teamTasks/:id PUT<br>
api/v2/teamTasks/:id DELETE<br> -->

#### api/v3/;

### Logic

#### Authorization

##### api/v1/auth/login
Type | JSON | headers
--- | --- | ---
Request | {"username": "Username", "password": "WDLKj1#@!#@odj123#@(0)-hash-"} | ---
Response | {"accessToken": "WDLKwekj1l32kjl.WEOQj213h.qwoei12KL#@!lk", "refreshToken": "qwe.123oid21.wqe123#@!lk"} | --- 
Error response | Error response | ---

##### api/v1/auth/registration
Type | JSON | headers
--- | --- | ---
Request | { "email":"email@email.com", "username": "Glebegor", "password": "HASHSHSHDHHS123p09D@1930j" } | ---
Response | Success response | ---
Error response | Error response | ---

##### api/v1/auth/refresh
Type | JSON | headers
--- | --- | ---
Request | {"refreshToken": "()213jdw).e21dddd.213dWQe21#@!qwe"} | ---
Response | {"accessToken": "WDLKwekj1l32kjl.WEOQj213h.qwoei12KL#@!lk", "refreshToken": "qwe.123oid21.wqe123#@!lk"} | ---
Error response | Error response | ---

## Enviroment

### Env

```
DB_PASSWORD="123321"
SECRET="key"
```

### Config

```
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

```
type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Email    string             `json:"email" bson:"email"`
	Role     string             `json:"role" bson:"role"`
}
```

#### Team

#### Profile

#### Team tasks

Roles:

<ul>
<li>Admin</li>
<li>User</li>
<li>Subscriber</li>
</ul>