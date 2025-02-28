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


#Commands to create table 

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    user_id TEXT UNIQUE NOT NULL,
    access_token TEXT,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    age INT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();


CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT DEFAULT 'Pending',
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_task_timestamp
BEFORE UPDATE ON tasks
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();




#docker command for check ip address 
`docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' containerID`