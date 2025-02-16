# FULLSTACK_ASSIGNMENT

## Objective

Demonstrate the ability of the candidate to build a full-stack application by integrating front-end, back-end, and Docker database components.
Required for deployment, with emphasis on best practices in project structuring and implementation.

## Overview

Frontend Development

- Use React library to Implement Frontend Component.
- Use WebPack to bundle web.
- Use Nginx and Dockerfile to serve web.

Backend Development

- Use Go with echo framework to create HTTP server
- Use GORM for connect database
- Use golang-migrate/migrate for migation database schema

## Preparing

- Install Migrate: [https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)

- Install Makefile:
    
    ```bash
    brew install make
    ```
    
## Step to run the project

1. Build and run

```bash
make reload
```

2. Load test data into database
3. Migrate database schema

```bash
 make migrate-up dir=db/migrations/sql db="mysql://myuser:mypassword@tcp(127.0.0.1:3306)/mydatabase"
```

4. Open URL: http://locahost it will lead you to  [http://localhost/splash](http://localhost/splash?userId=00006207e1a211ef95a30242ac180002) then you can add userId to be a query param like this http://localhost/splash?userId=00006207e1a211ef95a30242ac180002 (this solution assume the user device have a user id from some solution before)

## Project structures

```bash
api/
    cache/
        cache.go
        mocks/
    cmd/
        main.go
    config/
        config.go
    controller/
        auth_controller.go
        user_controller.go
    Dockerfile
    go.mod
    go.sum
    model/
        entity/
    repository/
        mocks/
        user_repository.go
    service/
        auth_service_test.go
        auth_service.go
        user_service.go
db/
    init/
        0_schema.sql
        ...
    migrations/
docker-compose.yaml
Makefile
web/
    Dockerfile
    nginx.conf
    package.json
    src/
    tsconfig.json
    webpack.config.js
```

**‭**

- **api/**: Contains the backend Go application.
    - **cache/**: Contains caching logic and mock implementations for testing.
        - cache.go: Defines the caching interface and its implementation.
        - `mocks/`: Contains mock implementations for testing.
    - **cmd/**: Contains the entry point of the Go application.
        - main.go
    - **config/**: Contains configuration-related files.
        - config.go
    - **controller/**: Contains the HTTP handlers/controllers.
        - auth_controller.go
        - user_controller.go
    - **model/**:
        - `entity/`: Contains the entity definitions for the database.
    - **repository/**: Contains the data access layer.
        - user_repository.go
        - `mocks/`: Contains mock implementations for testing.
    - **service/**: Contains the business logic.
        - auth_service.go
        - auth_service_test.go
        - user_service.go
    - Dockerfile: Dockerfile for building and running the Go application.
    - go.mod
    - go.sum
- **db/**: Contains database-related files.
    - **init/**: Contains initial SQL scripts for setting up the database schema.
        - 0_schema.sql: Initial schema setup script.
    - **migrations/**: Contains database migration files.
- **web/**: Contains the frontend React application.
    - Dockerfile: Dockerfile for building and running the React application.
    - nginx.conf: Nginx configuration for serving the React application.
    - package.json
    - src/: Contains the source code of the React application.
    - tsconfig.json
    - webpack.config.js
- **docker-compose.yaml**: Docker Compose file for setting up the multi-container environment.
- **Makefile**: Contains make commands for building, running, and managing the project.

## API endpoints

- **Login with PIN**
    - Endpoint: `/auth/login/pin`
    - Method: `POST`
    - Description: Authenticates a user using their ID and PIN.
    
    ### Body
    
    ```json
    {
      "id": "string",
      "pin": "string"
    }
    ```
    
    - **id**: The unique identifier for the user (string).
    - **pin**: The PIN code for authentication (string).
    
    ## Response Format
    
    ### Success Response
    
    **Status Code**: 200 OK
    
    ```json
    {
        "data": {
            "userId": "000043b3e1a211ef95a30242ac180002",
            "name": "User_000043b3e1a211ef95a30242ac180002",
            "banners": [
                {
                    "bannerId": "000043d4e1a211ef95a30242ac180002",
                    "userId": "000043b3e1a211ef95a30242ac180002",
                    "title": "Want some money?",
                    "description": "You can start applying",
                    "image": "https://dummyimage.com/54x54/999/fff"
                }
            ],
            "userGreetings": {
                "userId": "000043b3e1a211ef95a30242ac180002",
                "greeting": "Hello User_000043b3e1a211ef95a30242ac180002"
            },
            "accounts": [
                {
                    "accountId": "000044ebe1a211ef95a30242ac180002",
                    "userId": "000043b3e1a211ef95a30242ac180002",
                    "type": "saving-account",
                    "currency": "THB",
                    "accountNumber": "568-2-82678",
                    "issuer": "TestLab",
                    "dummyCol3": "dummy_value_3",
                    "accountBalances": {
                        "accountId": "000044ebe1a211ef95a30242ac180002",
                        "userId": "000043b3e1a211ef95a30242ac180002",
                        "amount": 41996.12,
                        "dummyCol4": "dummy_value_4"
                    },
                    "accountDetails": {
                        "accountId": "000044ebe1a211ef95a30242ac180002",
                        "userId": "000043b3e1a211ef95a30242ac180002",
                        "color": "#24c875",
                        "isMainAccount": true,
                        "progress": 61,
                        "dummyCol5": "dummy_value_5"
                    },
                    "accountFlags": [
                        {
                            "flagId": 3810103,
                            "accountId": "000044ebe1a211ef95a30242ac180002",
                            "userId": "000043b3e1a211ef95a30242ac180002",
                            "flagType": "system",
                            "flagValue": "Disbursement",
                            "createdAt": "2025-02-02T20:12:17Z",
                            "updatedAt": "2025-02-02T20:12:17Z"
                        },
                        {
                            "flagId": 3810104,
                            "accountId": "000044ebe1a211ef95a30242ac180002",
                            "userId": "000043b3e1a211ef95a30242ac180002",
                            "flagType": "system",
                            "flagValue": "Flag3",
                            "createdAt": "2025-02-02T20:12:17Z",
                            "updatedAt": "2025-02-02T20:12:17Z"
                        }
                    ]
                },
                {
                    "accountId": "00004908e1a211ef95a30242ac180002",
                    "userId": "000043b3e1a211ef95a30242ac180002",
                    "type": "saving-account",
                    "currency": "THB",
                    "accountNumber": "568-2-21484",
                    "issuer": "TestLab",
                    "dummyCol3": "dummy_value_3",
                    "accountBalances": {
                        "accountId": "00004908e1a211ef95a30242ac180002",
                        "userId": "000043b3e1a211ef95a30242ac180002",
                        "amount": 1978.96,
                        "dummyCol4": "dummy_value_4"
                    },
                    "accountDetails": {
                        "accountId": "00004908e1a211ef95a30242ac180002",
                        "userId": "000043b3e1a211ef95a30242ac180002",
                        "color": "#24c875",
                        "isMainAccount": false,
                        "progress": 45,
                        "dummyCol5": "dummy_value_5"
                    },
                    "accountFlags": [
                        {
                            "flagId": 3810105,
                            "accountId": "00004908e1a211ef95a30242ac180002",
                            "userId": "000043b3e1a211ef95a30242ac180002",
                            "flagType": "system",
                            "flagValue": "Overdue",
                            "createdAt": "2025-02-02T20:12:17Z",
                            "updatedAt": "2025-02-02T20:12:17Z"
                        },
                        {
                            "flagId": 3810106,
                            "accountId": "00004908e1a211ef95a30242ac180002",
                            "userId": "000043b3e1a211ef95a30242ac180002",
                            "flagType": "system",
                            "flagValue": "Flag5",
                            "createdAt": "2025-02-02T20:12:17Z",
                            "updatedAt": "2025-02-02T20:12:17Z"
                        }
                    ]
                },
                {
                    "accountId": "00004cc2e1a211ef95a30242ac180002",
                    "userId": "000043b3e1a211ef95a30242ac180002",
                    "type": "saving-account",
                    "currency": "THB",
                    "accountNumber": "568-2-41008",
                    "issuer": "TestLab",
                    "dummyCol3": "dummy_value_3",
                    "accountBalances": {
                        "accountId": "00004cc2e1a211ef95a30242ac180002",
                        "userId": "000043b3e1a211ef95a30242ac180002",
                        "amount": 55211.55,
                        "dummyCol4": "dummy_value_4"
                    },
                    "accountDetails": {
                        "accountId": "00004cc2e1a211ef95a30242ac180002",
                        "userId": "000043b3e1a211ef95a30242ac180002",
                        "color": "#24c875",
                        "isMainAccount": false,
                        "progress": 53,
                        "dummyCol5": "dummy_value_5"
                    },
                    "accountFlags": [
                        {
                            "flagId": 3810107,
                            "accountId": "00004cc2e1a211ef95a30242ac180002",
                            "userId": "000043b3e1a211ef95a30242ac180002",
                            "flagType": "system",
                            "flagValue": "Flag5",
                            "createdAt": "2025-02-02T20:12:17Z",
                            "updatedAt": "2025-02-02T20:12:17Z"
                        },
                        {
                            "flagId": 3810108,
                            "accountId": "00004cc2e1a211ef95a30242ac180002",
                            "userId": "000043b3e1a211ef95a30242ac180002",
                            "flagType": "system",
                            "flagValue": "Overdue",
                            "createdAt": "2025-02-02T20:12:17Z",
                            "updatedAt": "2025-02-02T20:12:17Z"
                        }
                    ]
                }
            ],
            "debitCards": [
                {
                    "cardId": "000043d4e1a211ef95a30242ac180002",
                    "userId": "000043b3e1a211ef95a30242ac180002",
                    "name": "My Debit Card",
                    "debitCardDesign": {
                        "cardId": "000043d4e1a211ef95a30242ac180002",
                        "userId": "000043b3e1a211ef95a30242ac180002",
                        "color": "#00a1e2",
                        "borderColor": "#ffffff"
                    },
                    "debitCardDetails": {
                        "cardId": "000043d4e1a211ef95a30242ac180002",
                        "userId": "000043b3e1a211ef95a30242ac180002",
                        "issuer": "TestLab",
                        "number": "1895 6835 8492 1957"
                    },
                    "debitCardStatus": {
                        "cardId": "000043d4e1a211ef95a30242ac180002",
                        "userId": "000043b3e1a211ef95a30242ac180002",
                        "status": "Active"
                    }
                }
            ],
            "transactions": [
                {
                    "transactionId": "000043c6e1a211ef95a30242ac180002",
                    "userId": "000043b3e1a211ef95a30242ac180002",
                    "name": "Transaction_135018",
                    "image": "https://dummyimage.com/54x54/999/fff",
                    "isBank": false
                }
            ]
        },
        "message": "success"
    }
    ```
    
- **Get User by ID**
    - Endpoint: `/user/:id`
    - Method: `GET`
    - Description: Retrieves user information based on the provided user ID.
    
    ## Request Parameters
    
    - `id` (path parameter): The unique identifier of the user to retrieve. This should be a string.
    
    ## Request Example
    
    ```
    GET /user/12345
    ```
    
    ## Response Format
    
    ### Success Response
    
    - **Status Code**: 200 OK
    - **Content**:
    
    ```json
    {
        "data": {
            "userId": "000043b3e1a211ef95a30242ac180002",
            "name": "User_000043b3e1a211ef95a30242ac180002",
            "banners": null,
            "userGreetings": {
                "userId": "",
                "greeting": ""
            },
            "accounts": null,
            "debitCards": null,
            "transactions": null
        },
        "message": "success"
    }
    ```
   
### วิธีการ Migrate Data

อ้างอิงจาก <https://www.freecodecamp.org/news/database-migration-golang-migrate/>

1. ติดตั้งเครื่องมือที่ใช้ในการ migrate ผ่าน terminal ตามลิงก์ แยกตามระบบปฏิบัติการ (Windows และ macOS)
2. สร้างไฟล์ migration ในโฟลเดอร์ `db/migration/<data-base-type>` โดยใช้คำสั่ง:

   ```bash
   make migrate-create dir=[path] name=[name]
   ```

   คำสั่งนี้จะสร้างไฟล์ migration สำหรับ `up` และ `down` อย่างละหนึ่งไฟล์

3. เขียนสคริปต์สำหรับไฟล์ `up/down` ตามโครงสร้างที่ต้องการ
4. รันคำสั่ง migration:

   **migration_up:**

   ```bash
   make migrate-up path=[path] database=[database]
   ```

   **migration_down:**

   `step` คือจำนวนไฟล์ที่ต้องการ rollback (หากต้องการ rollback กลับไปก่อนหน้า 1 ไฟล์ ให้ใส่ `1`)

   ```bash
    make migrate-down path=[path] database=[database] step=[step]
   ```

5. หากเกิดข้อผิดพลาด เช่น `Dirty database version 20231006035122. Fix and force version.`
   อ้างอิงจากลิงก์ <https://stackoverflow.com/questions/59616263/dirty-database-version-error-when-using-golang-migrate>

   **แก้ไขโดยใช้ force version:**

   ```bash
    make migrate-force path=[path] database=[database] version=[version]
   ```

   จากนั้นให้รันคำสั่ง `up` เพื่อ migrate ข้อมูลอีกครั้ง:

   ```bash
    make migrate-up path=[path] database=[database]
   ```

