# GoPostAPI

A simple Go REST API to fetch and create posts using an SQL Server database (localhost). Initially demonstrated with an external API (`jsonplaceholder.typicode.com`), then migrated to use a local `SQL Server 2019 Express` instance.

---

## 📁 Project Structure

GoProject/ 
├── models/ # Data models (Post struct) 
├── services/ # Business logic (GetFilteredPosts, SavePost) 
├── handlers/ # HTTP route handlers 
├── data/ # DB connection and queries 
├── main.go # App entrypoint (routing) 
└── go.mod # Go module file

---
## Sql Driver
[SQL Guide](docs/sql.md)

## AWS
[AWS Guide](docs/aws.md)

## GIT
[GIT Guide](docs/git.md)

## Jenkins
[Jenkins Guide](docs/jenkins.md)

## 🚀 Running the App

1. For MSSQL **Start SQL Server** (ensure port 1433 is open and TCP/IP enabled)
2. For MySQL 
    a. Install Custom
    b. tools needed: MySQL Server, Workbench, Shell (optional)
    c. TCP/IP: 3306
    d. Open Windows Firewall ports — Important if you access MySQL from another tool or app.
    e. Setup with legacy auth method

3. **Run the app**:
```bash
go mod init
go run main.go
✅ Step 4: Open your browser
Go to:
👉 http://localhost:8092/
You’ll see your index.html load and connect to your Go API automatically.
```

---
# Example requests
```go
curl -X POST http://localhost:8092/posts \
  -H "Content-Type: application/json" \
  -d "{\"userId\":1,\"title\":\"\",\"body\":\"This is a test\"}"
```
---------------------------------------
MIT License
© 2025 Sandeep B and Ramana. All rights reserved.
---------------------------------------
