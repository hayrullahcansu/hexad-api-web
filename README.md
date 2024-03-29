# Library Management


## Table of Contents

- [Library Management](#library-management)
  - [Table of Contents](#table-of-contents)
  - [Description](#description)
  - [Features](#features)
  - [Structure](#structure)
  - [Testing](#testing)
    - [Step 1: Go to api folder of project](#step-1-go-to-api-folder-of-project)
    - [Step 2: Run on Docker](#step-2-run-on-docker)
    - [Step 3: Test Output](#step-3-test-output)
  - [Run on Local](#run-on-local)
    - [Step 1: Go to root folder of project](#step-1-go-to-root-folder-of-project)
    - [Step 2: Run on Docker](#step-2-run-on-docker-1)
    - [Step 3: Go to browser](#step-3-go-to-browser)
  - [API Document](#api-document)

---


## Description

Library Management is developed for Hexad company.

## Features

* TDD
* Docker Image & Compose

## Structure


To be avoided from circular reference error, we have to obtain these rules. 

 * `Handler Layer` includes request handlers, and routing files  should be here. You can just reference Repo Layer, Data Layer and Utility Layer.

 * `Repo Layer` includes all business logics, structures, service files. You can just reference Data layer and Utility layer.

 * `Data Layer` includes Models, DTOs, views, constants. You `cannot reference any object from other layers`.

 * `Utility Layer` Includes global methods, some helper methods without business. You can just reference Data layer.
  

## Testing

### Step 1: Go to api folder of project
```
cd api
```
### Step 2: Run on Docker

```
go test -v ./... -count=1
```

### Step 3: Test Output
```
(base) ┌─(~/Workspace/golang/libman/api)──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────(hayrullahcansu@Hayrullahs-MacBook-Pro:s001)─┐
└─(12:12:40 on master ✹)──> go test -v ./... -count=1                                                                                                                                     ──(Mon,May23)─┘
?       lib-api [no test files]
?       lib-api/data    [no test files]
=== RUN   TestGETBooks
migration first
migration done
=== RUN   TestGETBooks/User_can_view_books_in_library
=== RUN   TestGETBooks/User_can_borrow_a_book_from_the_library

2022/05/23 12:12:46 /Users/hayrullahcansu/Workspace/golang/libman/api/repo/library_repository.go:60 record not found
[0.098ms] [rows:0] SELECT * FROM `borrows` WHERE user = "test1_username" AND name = "TestBook1" ORDER BY `borrows`.`user` LIMIT 1
=== RUN   TestGETBooks/User_can_see_borrowed_list
=== RUN   TestGETBooks/User_can_return_a_book_from_the_library
--- PASS: TestGETBooks (0.01s)
    --- PASS: TestGETBooks/User_can_view_books_in_library (0.00s)
    --- PASS: TestGETBooks/User_can_borrow_a_book_from_the_library (0.00s)
    --- PASS: TestGETBooks/User_can_see_borrowed_list (0.00s)
    --- PASS: TestGETBooks/User_can_return_a_book_from_the_library (0.00s)
PASS
ok      lib-api/handler 0.440s
=== RUN   TestDbContext
migration first
migration done
--- PASS: TestDbContext (0.00s)
=== RUN   TestLibraryRepository
=== RUN   TestLibraryRepository/returns_the_list_of_books_in_the_library_from_db
--- PASS: TestLibraryRepository (0.00s)
    --- PASS: TestLibraryRepository/returns_the_list_of_books_in_the_library_from_db (0.00s)
PASS
ok      lib-api/repo    0.370s
?       lib-api/utility [no test files]
```


## Run on Local 

### Step 1: Go to root folder of project


### Step 2: Run on Docker

```
docker compose up
```
```
└─(09:20:45 on readme ✹)──> docker compose up                                                                                                                                                                      ──(Mon,May23)─┘
[+] Running 2/0
 ⠿ Container lib-web  Created                                                                                                                                                                                                 0.0s
 ⠿ Container lib-api  Created                                                                                                                                                                                                 0.0s
Attaching to lib-api, lib-web
lib-api  | migration first
lib-api  | migration done
lib-api  | database seed first
lib-api  | database seed done
lib-api  | 2022/05/23 06:21:16 Defaulting to port 5002
lib-api  | 2022/05/23 06:21:16 Listening on port 5002
lib-web  | Starting up http-server, serving dist
lib-web  | 
lib-web  | http-server version: 14.1.0
lib-web  | 
lib-web  | http-server settings: 
lib-web  | CORS: disabled
lib-web  | Cache: 3600 seconds
lib-web  | Connection Timeout: 120 seconds
lib-web  | Directory Listings: visible
lib-web  | AutoIndex: visible
lib-web  | Serve GZIP Files: false
lib-web  | Serve Brotli Files: false
lib-web  | Default File Extension: none
lib-web  | 
lib-web  | Available on:
lib-web  |   http://127.0.0.1:8080
lib-web  |   http://172.18.0.2:8080
lib-web  | Hit CTRL-C to stop the server
```

### Step 3: Go to browser
`http://localhost:8080/#/`

## API Document

Supported ContentType is `x-www-form-urlencoded` for POST request.


All endpoints here
- Book Resource `/api/v1/books`
  - GET `/` returns array of all books in the library as string.
    ```
    [{"Name":"Sapiens: A Brief History of Humankind","Quantity":5},{"Name":"Outliers: The Story of Success","Quantity":2},{"Name":"Thinking, Fast and Slow","Quantity":1}]
    ```
  - POST `/borrow` to borrow a book from the library.
  
    Request Parameters

    | Param | Type   | Example Value                         |
    |-------|--------|---------------------------------------|
    | User  | String | test1                                 |
    | Book  | String | Sapiens: A Brief History of Humankind |

    - Success Response body, StatusCode is 200 
    ```json
    {"User":"test1","Name":"Thinking, Fast and Slow"}
    ```
    - Fail Response body, StatusCode is 404 
    ```txt
    you cannot borrow Thinking, Fast and Slow book
    ```

- POST `/return` to return a book which was borrowed from the library.
  
    Request Parameters

    | Param | Type   | Example Value                         |
    |-------|--------|---------------------------------------|
    | User  | String | test1                                 |
    | Book  | String | Sapiens: A Brief History of Humankind |

    - Success Response body, StatusCode is 200 
    ```json
    {"User":"test1","Name":"Thinking, Fast and Slow"}
    ```
    - Fail Response body, StatusCode is 404 
    ```txt
    you cannot return Thinking, Fast and Slow book, because you don't have
    ```

- Borrowed Resource `/api/v1/borrowed`
  - GET `/{username}` returns array of all books which are borrowed by username as string
    ```
    [{"User":"test1","Name":"Sapiens: A Brief History of Humankind"},{"User":"test1","Name":"Thinking, Fast and Slow"}]
    ```