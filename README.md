Here is your **`README.md`** file for the URL Shortener project:

---

```markdown
# URL Shortener (Go + PostgreSQL)

A simple URL shortener built using **Go**, **Gin**, **GORM**, and **PostgreSQL**.

---

## **Features**
- Shorten long URLs into unique short codes.
- Redirect short codes to the original URLs.
- Auto-database migration with GORM.
- Environment-based configuration using `.env`.
- REST API endpoints for shortening and retrieving URLs.

---

## **Tech Stack**
- **Backend:** Go
- **Framework:** Gin
- **Database:** PostgreSQL
- **ORM:** GORM
- **Environment Management:** godotenv
- **ID Generator:** shortid

---

## **Project Structure**
```

url-shortener/
│
├── main.go
├── go.mod
├── go.sum
├── .env
├── .gitignore
│
├── database/
│   └── db.go
│
├── models/
│   └── url.go
│
├── routes/
│   └── url\_routes.go
│
└── handlers/
└── url\_handler.go

````

---

## **Setup Instructions**

### **1. Clone the repository**
```bash
git clone https://github.com/your-username/url-shortener.git
cd url-shortener
````

### **2. Install dependencies**

```bash
go mod tidy
```

### **3. Create a `.env` file**

Add your database credentials:

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=urlshortener
DB_PORT=5432
```

### **4. Set up PostgreSQL**

Create a new database:

```sql
CREATE DATABASE urlshortener;
```

---

## **5. Run the Project**

```bash
go run main.go
```

Server will start at:

```
http://localhost:8080
```

---

## **6. API Endpoints**

### **POST /shorten**

Request:

```json
{
  "original_url": "https://example.com/very-long-url"
}
```

Response:

```json
{
  "short_url": "http://localhost:8080/abc123"
}
```

### **GET /\:shortId**

Redirects to the original URL.

---

## **7. Build the Project**

```bash
go build -o url-shortener
./url-shortener
```

---

## **8. Dependencies**

Install all required packages:

```bash
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/joho/godotenv
go get github.com/teris-io/shortid
```

---

## **License**

MIT License

```

---

### **Do you want me to generate a `main.go` file with fully working routes (POST /shorten and GET /:id) so you can test this project immediately?**
```
