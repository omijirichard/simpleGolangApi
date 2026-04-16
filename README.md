

```markdown
# Go REST API Server

A lightweight, high-performance REST API written in Go using the **Gin** framework. This server includes endpoints for API status, health checks, and user profile information.

## 🚀 Prerequisites

To run this application, you need the following installed on your machine:

*   **Go** (Golang): Version 1.20 or higher is recommended.
*   **Git**: (Optional, if cloning from a repository)
*   **Run Environment**: Linux, macOS, or Windows (WSL).

---

## ⚙️ How to Run

This application is designed to be a standalone executable binary. Follow these steps to set up and run the server:

### 1. Initialize Go Modules
Ensure you are in the directory containing your `main.go` file. If you haven't used Go Modules yet, initialize the project and fetch dependencies:

```bash
go mod init api-server
go get github.com/gin-gonic/gin
go mod tidy
```

### 2. Build the Executable
Compile the source code into a binary executable. You can name the output binary whatever you prefer (e.g., `my-api`, `server`, or `api`):

```bash
go build -o api-server
```

*   **Linux/Mac:** Run `go build` to generate `api-server`.
*   **Windows:** Run `go build` to generate `api-server.exe`.

### 3. Run the Executable
Once the binary is built, you can run it directly. **Note:** Ensure you close any other applications using port **2222** first.

```bash
./api-server
```

*   **Windows:** `api-server.exe`
*   **Linux/Mac:** `./api-server`

The server will start and listen on **Port 2222**.

---

## 📖 API Endpoints

Once the executable is running, you can interact with the API using `curl`, `Postman`, or a web browser.

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/` | **Root Endpoint**<br>Confirms the API is running. |
| `GET` | `/health` | **Health Check**<br>Returns server status. |
| `GET` | `/me` | **User Profile**<br>Returns user info (Name, Email, GitHub). |

### Example Requests

**Using cURL:**

```bash
# Check if API is alive
curl http://localhost:2222/

# Check Health Status
curl http://localhost:2222/health

# Get User Info
curl http://localhost:2222/me
```

**Using a Browser:**
Simply navigate to: `http://localhost:2222` (or your specific machine's IP if remote).

---

## 📝 Source Code Overview

The server is defined in `main.go`.

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Root route
    router.GET("/", func(currentContext *gin.Context){
        currentContext.JSON(http.StatusOK, gin.H{ "message": "API is running"})
    })

    // Health check route
    router.GET("/health", func(currentContext *gin.Context){
        currentContext.JSON(http.StatusOK, gin.H{ "message": "healthy"})
    })

    // User info route
    router.GET("/me", func(currentContext *gin.Context){
        currentContext.JSON(http.StatusOK, gin.H{ 
            "name": "Omiji Richard",
            "email": "omijirichard@gmail.com", 
            "github": "https://github.com/omijirichard" 
        })
    })

    // Start the server on port 2222
    router.Run(":2222")
}
```

---

## 📄 License

Feel free to use this code for personal or commercial projects.

## 📧 Contact

If you have questions or need help with the API:

*   **Email:** omijirichard@gmail.com
*   **GitHub:** https://github.com/omijirichard
```
