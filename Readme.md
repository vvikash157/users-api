Database Connection Guide 🚀
This repository uses PostgreSQL as the database and Redis as a caching layer. Below are the steps to set up and connect to the database.

🔹 Setup Using Docker
Run the following commands to build and start the necessary containers:

`docker build -t Login .`
`docker compose up -d --build`
These commands will:
✅ Build the Docker image
✅ Start a PostgreSQL and Redis server in containers

🔹 PostgreSQL Connection String
Use the following connection string to connect to PostgreSQL:

`"postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable"`
💡 Connecting to PostgreSQL
DBBeaver: You can use DBeaver to connect to the database.
VS Code: Several extensions are available to connect to PostgreSQL in VS Code.
🔹 Connecting to Redis
Use RedisInsight to connect and manage your Redis database.

🔹 Important Note
🔸 All credentials are stored in the .env file. Please check the file for configuration details.



🔹 User Authentication Flow
1️⃣ User Signup:

When a user signs up, a user ID and an access token are generated.
The access token is stored in Redis for session management.
2️⃣ API Authorization:

The user must include the access token in the Authorization header of all API requests.
3️⃣ Access Token Expiry Handling:

If the access token expires, the user will be redirected to the Login API.
The user must log in again to get a new access token, which is then stored in Redis.

🔹 Security Features
✅ Authentication Middleware:

Applied to all APIs to enforce authentication.
✅ Rate Limiting:

Implemented using the user's IP address to prevent API abuse.

🔹 Database Migrations
A database migration path is set up to handle future model changes easily.

To create a new migration file, run the following command:
`migrate create -ext sql -dir ./.db/migrations add_your_file_name`
Place the sql commands to the created file inside the .db/migrations/ folder to apply changes.

🔹 Logging System
This project uses the Logrus package for flexible logging.

🔹 Supports different log levels:

Info level 🟢
Error level 🔴
Debug level 🟡


Curls for Api:
1. UserSignup:

 curl --location 'http://localhost:8080/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "vikash",
    "email": "abc@gmail.com",
    "password": "abc@12",
    "age": 30
}'  

2.User Login

curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqd3RfZXhwaXJ5IjoxNzQwODYxNTU1LCJ1c2VyaWQiOiIwM2QzNDUyZC1mYWIzLTQ3NmYtYjAxMS0xY2Y1NzUxOTRiMTIifQ.omul7iGCn-PZnvkwmTrx3yzdgOMV_pQrMcXT_ri-zM4' \
--data-raw '{
    "email":"abc@gmail.com",
    "password":"abc@12"
}'

3.Create Task:
curl --location 'http://localhost:8080/tasks' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqd3RfZXhwaXJ5IjoxNzQwODYxNTU1LCJ1c2VyaWQiOiIwM2QzNDUyZC1mYWIzLTQ3NmYtYjAxMS0xY2Y1NzUxOTRiMTIifQ.omul7iGCn-PZnvkwmTrx3yzdgOMV_pQrMcXT_ri-zM4' \
--data '{
    "title": "singham",
    "description": "actor is ajay ",
    "status": "approved"
}'

4. Get Task By page and no of records 
curl --location 'http://localhost:8080/tasks?page=1&pageSize=2&status=Done' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqd3RfZXhwaXJ5IjoxNzQwODYxNTU1LCJ1c2VyaWQiOiIwM2QzNDUyZC1mYWIzLTQ3NmYtYjAxMS0xY2Y1NzUxOTRiMTIifQ.omul7iGCn-PZnvkwmTrx3yzdgOMV_pQrMcXT_ri-zM4' \
--data ''

5. Get task by id 
curl --location 'http://localhost:8080/tasks/get/4' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqd3RfZXhwaXJ5IjoxNzQwODYxNTU1LCJ1c2VyaWQiOiIwM2QzNDUyZC1mYWIzLTQ3NmYtYjAxMS0xY2Y1NzUxOTRiMTIifQ.omul7iGCn-PZnvkwmTrx3yzdgOMV_pQrMcXT_ri-zM4'

6. Update a task 
curl --location --globoff --request PUT 'http://localhost:8080/tasks/{1}' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqd3RfZXhwaXJ5IjoxNzQwODYxNTU1LCJ1c2VyaWQiOiIwM2QzNDUyZC1mYWIzLTQ3NmYtYjAxMS0xY2Y1NzUxOTRiMTIifQ.omul7iGCn-PZnvkwmTrx3yzdgOMV_pQrMcXT_ri-zM4' \
--data '{
    "description": "thriller movie",
    "status": "approved"
}'

7. delete a task by id 


