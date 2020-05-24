# RESTful API Server with GOlang
This is a baby steps in GOlang. 

## How to build and run?

<pre>

    go build server.go
    ./server // for mac

</pre>

### Allowed methods: **POST**, **GET**, **PATCH**, **DELETE**

## /posts (GET)

* List all posts

## /posts (POST)

| Parameter | Type  |
|---|---|
| ID  | integer  |
| Title | string |
| Description | string |

## /posts/{id} (GET)

* Get specific post by ID

| Parameter      | Type |
| ----------- | ----------- |
| ID      | integer       |

## /posts/{id} (PATCH)

* Update post by ID

| Parameter | Type  |
|---|---|
| Title | string |
| Description | string |

## /posts/{id} (DELETE)

* Delete post by ID

| Parameter      | Type |
| ----------- | ----------- |
| ID      | integer       |
