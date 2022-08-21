# Flag Dash API
This is simple documentation to mention all Flag Dash API endpoints.

## General API Information
- The base endpoint is http(s)://domain:3000
- All endpoints return either JSON Object or JSON Array

## General Endpoint Information
- For `GET` endpoints request, send parameters as query string.
- For `POST`,`PATCH`, `DELETE` endpoints, only send parameters as the body
- Parameters can be send in any order.

# User Group Endpoints
All endpoints related to User and Authentication.

## User Login
```
POST /api/v1/auth/login
```
User login endpoints, it will return JWT Token, JWT Refresh Token, and Cookie
**Parameters:**

Name | Type | Mandatory | Description
------------ | ------------ | ------------ | ------------
email | STRING | YES | `BODY` parameter.
password | STRING | YES | `BODY` parameter.


**Response:**
```JSON
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAZXhhbXBsZS5jb20iLCJJRCI6IjEiLCJleHAiOjE2NjEwNTgxNTd9.KfC5kWZpsUedw4iUNf-IAj2wwcydJIwPGwOalFq9yHI",
    "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAZXhhbXBsZS5jb20iLCJJRCI6IjEiLCJleHAiOjE2NjEwNjE2OTd9.sqQm5SLBaaZs50MCzW8LKgReCxMP0aV8htusY6YdXMk"
}
```

## Refresh Token
```
GET /api/v1/auth/refresh-token
```
Refresh user token endpointsm it will return new token.
**Parameters:**

Name | Type | Mandatory | Description
------------ | ------------ | ------------ | ------------
| | |
**Response:**
```JSON
{
    "status": "201",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAZXhhbXBsZS5jb20iLCJJRCI6IjEiLCJleHAiOjE2NjEwNTgxNTd9.KfC5kWZpsUedw4iUNf-IAj2wwcydJIwPGwOalFq9yHI"
}
```

## Get Users
```
GET /api/v1/auth/
```
Get all users by using pagination.
**Parameters:**

Name | Type | Mandatory | Description
------------ | ------------ | ------------ | ------------
page_num | INT | YES |  Default 1;
limit | INT | YES | Default 12;
filter | STRING | NO | Default " ";
Authorization | HEADER | YES | Authentication, use `Token`.

**Response:**
```JSON
{
    "limit": 12,
    "page_num": 1,
    "max_page": 1,
    "users": [
        {
            "email": "test@example.com",
            "userProfile": {
                "firstName": "John",
                "lastName": "Doe"
            },
            "userRole": {
                "Name": "Low Staff",
                "Level": 1
            }
        }
    ]
}
```

## Get User By ID
```
GET /api/v1/auth/:id
```
Get user by using their ID.
**Parameters:**

Name | Type | Mandatory | Description
------------ | ------------ | ------------ | ------------
id | INT | YES |  `URL` parameter.
Authorization | HEADER | YES | Authentication, use `Token`.
**Response:**
```JSON
{
    "status": "200",
    "msg": "Successfully get user by id",
    "User": {
        "email": "test@example.com",
        "userProfile": {
            "firstName": "John",
            "lastName": "Doe"
        },
        "userRole": {
            "Name": "Low Staff",
            "Level": 1
        }
    }
}
```
## Get User By Email
```
GET /api/v1/auth/find-by-email
```
Get user by using their E-Mail.
**Parameters:**

Name | Type | Mandatory | Description
------------ | ------------ | ------------ | ------------
email | STRING | YES | `BODY` parameter.
Authorization | HEADER | YES | Authentication, use `Token`.

**Response:**
```JSON
{
    "status": "200",
    "msg": "Successfully get user by e-mail",
    "User": {
        "email": "test@example.com",
        "userProfile": {
            "firstName": "John",
            "lastName": "Doe"
        },
        "userRole": {
            "Name": "Low Staff",
            "Level": 1
        }
    }
}
```

## Get User Profile
```
GET /api/v1/auth/profile
```
Get user profile by using provided token.
**Parameters:**

Name | Type | Mandatory | Description
------------ | ------------ | ------------ | ------------
Authorization | HEADER | YES | Authentication, use `Token`.

**Response:**
```JSON
{
    "status": "200",
    "msg": "Successfully get user by id",
    "User": {
        "email": "test@example.com",
        "userProfile": {
            "firstName": "John",
            "lastName": "Doe"
        },
        "userRole": {
            "Name": "Low Staff",
            "Level": 1
        }
    }
}
```

## Get User Profile
```
GET /api/v1/auth/profile
```
Get user profile by using provided token.
**Parameters:**

Name | Type | Mandatory | Description
------------ | ------------ | ------------ | ------------
Authorization | HEADER | YES | Authentication, use `Token`.

**Response:**
```JSON
{
    "status": "200",
    "msg": "Successfully get user by id",
    "User": {
        "email": "test@example.com",
        "userProfile": {
            "firstName": "John",
            "lastName": "Doe"
        },
        "userRole": {
            "Name": "Low Staff",
            "Level": 1
        }
    }
}
```

## Edit User
```
PUT /api/v1/auth/:id
```
Edit user information by using provided id
**Parameters:**

Name | Type | Mandatory | Description
------------ | ------------ | ------------ | ------------
id | INT | YES | `URL` parameter.
Authorization | HEADER | YES | Authentication, use `Token`.

**Response:**
```JSON
{
    "status": "200",
    "msg": "Successfully edit user by id",
    "User": {
        "email": "test@example.com",
        "userProfile": {
            "firstName": "John",
            "lastName": "Doe"
        },
        "userRole": {
            "Name": "Low Staff",
            "Level": 1
        }
    }
}
```

## Delete User
```
DELETE /api/v1/auth/:id
```
Delete user information by using provided id
**Parameters:**

Name | Type | Mandatory | Description
------------ | ------------ | ------------ | ------------
id | INT | YES | `URL` parameter.
Authorization | HEADER | YES | Authentication, use `Token`.

**Response:**
```JSON
{
    "status": "200",
    "msg": "Successfully delete user by id",
}
```

# Project Group Endpoints
All endpoints related to Projects

## Get Projects
```
GET /api/v1/project/
```
Get projects by using pagination system.
**Parameters:**

Name | Type | Mandatory | Description
------------ | ------------ | ------------ | ------------
limit | INT | YES |
page_num | INT | YES | 
filter | STRING | NO |
Authorization | HEADER | YES | Authentication, use `Token`.

**Response:**
```JSON
{
    "projects": [
        {
            "id": 2,
            "name": "antex-dash",
            "access_key": null,
            "updated_at": {
                "Time": "2022-07-15T14:42:16.382299Z",
                "Valid": true
            }
        },
        {
            "id": 32,
            "name": "Test1",
            "access_key": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6IjMyIn0.05AD74B2s8iSYhnJI6vNUFUEQa1ICS7-HUS4zs9Ctj4",
            "updated_at": {
                "Time": "2022-08-19T07:11:30.841211Z",
                "Valid": true
            }
        },
        {
            "id": 33,
            "name": "test2",
            "access_key": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6IjMzIn0.fdZCJqC39r7llrNjdT9OKrpjbkyOuvjIzLyz-wmlnbM",
            "updated_at": {
                "Time": "2022-08-19T08:00:53.13137Z",
                "Valid": true
            }
        },
        {
            "id": 3,
            "name": "PasarOn",
            "access_key": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6IjMifQ.ApUPouOIXUyNbZ7HFbUXPOXUflM7C1NvB_rDvXubBqc",
            "updated_at": {
                "Time": "2022-08-19T08:06:53.060027Z",
                "Valid": true
            }
        }
    ],
    "limit": 5,
    "page_num": 1,
    "max_page": 1
}
```

## Get Project
```
GET /api/v1/project/:id
```
Get specific project by providing project id.
**Parameters:**

Name | Type | Mandatory | Description
------------ | ------------ | ------------ | ------------
id | INT | YES | `URL` parameter.
Authorization | HEADER | YES | Authentication, use `Token`.

**Response:**
```JSON
{
    "status": "200",
    "project": {
        "id": 2,
        "name": "antex-dash",
        "access_key": null,
        "updated_at": {
            "Time": "2022-07-15T14:42:16.382299Z",
            "Valid": true
        }
    }
}
```

## New Project
```
POST /api/v1/project/new-project
```
Create new project by providing several bodies needed.
**Parameters:**

Name | Type | Mandatory | Description
------------ | ------------ | ------------ | ------------
name | STRING | YES | `BODY` parameter.
Authorization | HEADER | YES | Authentication, use `Token`.

**Response:**
```JSON
{
    "status": "201",
    "project": {
        "id": 34,
        "name": "Test7",
        "access_key": null,
        "updated_at": {
            "Time": "2022-08-21T09:13:48.357844Z",
            "Valid": true
        }
    }
}
```

## Edit Project
```
PATCH /api/v1/project/:id
```
Edit project of provided id.
**Parameters:**

Name | Type | Mandatory | Description
------------ | ------------ | ------------ | ------------
id | INT | YES | `URL` parameter.
name | STRING | YES | `BODY` parameter.
Authorization | HEADER | YES | Authentication, use `Token`.

**Response:**
```JSON
{
    "status": "201",
    "project": {
        "id": 34,
        "name": "Hello World123",
        "access_key": null,
        "updated_at": {
            "Time": "2022-08-21T09:17:10.475803Z",
            "Valid": true
        }
    }
}
```

## Delete Project
```
DELETE /api/v1/project/:id
```
Delete project of provided id.
**Parameters:**

Name | Type | Mandatory | Description
------------ | ------------ | ------------ | ------------
id | INT | YES | `URL` parameter.
Authorization | HEADER | YES | Authentication, use `Token`.

**Response:**
```JSON
{
    "status": "200",
    "msg": "Delete Project Success"
}
```

# Flag Group Endpoints
All endpoints related to flags

## Get Flags
```
GET /api/v1/flag/
```
Get flags by pagination of provided project_id
**Parameters:**

Name | Type | Mandatory | Description
------------ | ------------ | ------------ | ------------
project_id | INT | YES | `QUERY` parameter / `BODY` parameter.
limit | INT | YES | Default 12;
page_num | INT | YES | Default 1;
filter | STRING | NO |
Authorization | HEADER | YES | Authentication, use `Token`.

**Response:**
```JSON
{
    "flags": [
        {
            "id": 30,
            "project_id": 33,
            "name": "new_feature1",
            "active": true,
            "updated_at": {
                "Time": "2022-08-21T09:43:08.897547Z",
                "Valid": true
            }
        }
    ],
    "limit": 12,
    "page_num": 1,
    "max_page": 1
}
```

## New Flag
```
POST /api/v1/flag/new-flag
```
Create new flag
**Parameters:**

Name | Type | Mandatory | Description
------------ | ------------ | ------------ | ------------
project_id | INT | YES | For reference of which project.
name | STRING | YES | Use name that will be use as flag.
active | BOOL | NO | Default false.
Authorization | HEADER | YES | Authentication, use `Token`.

**Response:**
```JSON
{
    "status": "201",
    "flag": {
        "id": 33,
        "project_id": 33,
        "name": "new-feature2",
        "active": false,
        "updated_at": {
            "Time": "2022-08-21T09:53:12.420943Z",
            "Valid": true
        }
    }
}
```

## Edit Flag
```
PATCH /api/v1/flag/:id
```
Edit flag of provided id
**Parameters:**

Name | Type | Mandatory | Description
------------ | ------------ | ------------ | ------------
id | INT | YES | For reference of which flag.
name | STRING | YES | Use name that will be use as flag.
active | BOOL | NO | Default false.
Authorization | HEADER | YES | Authentication, use `Token`.

**Response:**
```JSON
{
    "status": "200",
    "flag": {
        "id": 33,
        "project_id": 33,
        "name": "new-feature3",
        "active": true,
        "updated_at": {
            "Time": "2022-08-21T09:53:12.420943Z",
            "Valid": true
        }
    }
}
```

## Delete Flag
```
DELETE /api/v1/flag/:id
```
Delete flag of provided id
**Parameters:**

Name | Type | Mandatory | Description
------------ | ------------ | ------------ | ------------
id | INT | YES | For reference of which flag.
Authorization | HEADER | YES | Authentication, use `Token`.

**Response:**
```JSON
{
    "status": "200",
    "msg": "Successfully delete flag"
}
```