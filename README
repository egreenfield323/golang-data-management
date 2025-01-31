# User Management Application

This is a simple User Management Application built with Go (backend), MySQL (database), and HTML/JavaScript (frontend). It allows you to perform CRUD operations (Create, Read, Update, Delete) on users stored in a MySQL database.

## Prerequisites

Ensure you have the following installed:

- **Go (version 1.21 or higher)**: [Download Go](https://go.dev/dl/)
- **MySQL**: [Download MySQL](https://www.mysql.com/downloads/)
- **Python (for running the frontend. An npm http server will do too)**: [Download Python](https://www.python.org/downloads/)

## Setup Instructions

### 1. Set Up MySQL Database

#### Start MySQL Server:

Ensure your MySQL server is running locally.

This is easiest to do using MySQL Configurator and MySQL Workbench.

#### Create the Database:

Create the `data_management` database:

```sql
CREATE DATABASE data_management;
```

#### Create the `users` Table:

Switch to the `data_management` database:

```sql
USE data_management;
```

Create the `users` table:

```sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE
);
```


### 2. Set Up the GoLang Backend

#### Clone the Repository:

```bash
git clone https://github.com/your-username/user-management-app.git
cd user-management-app
```

#### Install Dependencies:

```bash
go mod tidy
```

#### Update Database Configuration:

Open `db/db.go` and update the MySQL connection string with your credentials:

```go
dsn := "root:password@tcp(localhost:3306)/data_management?parseTime=true"
```

Replace `root` and `password` with your MySQL username and password. If your database is named differently, change that too.

#### Run the Backend:

```bash
go run main.go
```

The server will start on [http://localhost:8080](http://localhost:8080).

### 3. Set Up the Frontend

#### Serve the Frontend:

Open a new terminal window and navigate to the project directory.

Start a python http server.

```bash
python -m http.server
```

#### Access the Application:

Open your browser and go to [http://localhost:8000](http://localhost:8000).

Use the form to create, read, update, and delete users.

## Application Features

- **Create User**: Add a new user to the database.
- **Read Users**: View a list of all users in the database.
- **Update User**: Edit an existing user's details.
- **Delete User**: Remove a user from the database.