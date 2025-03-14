

### **URL Shortener (Golang & Redis)**  

A simple and efficient URL shortener built with **Golang (Gin Framework)** and **Redis** for fast URL storage and retrieval.  

## 🚀 Features  
✅ Shorten long URLs with randomly generated or custom aliases  
✅ Set expiration times for shortened URLs  
✅ Redirect users from short URLs to the original URLs  
✅ High-speed performance using **Redis**  
✅ Lightweight API using **Gin Framework**  

---

## 🛠 Installation & Setup  

### **1️⃣ Prerequisites**  
Ensure you have the following installed:  
- **Go** (Download: [https://go.dev/dl/](https://go.dev/dl/))  
- **Redis** (Download: [https://redis.io/docs/getting-started/](https://redis.io/docs/getting-started/))  

### **2️⃣ Clone the Repository**  
```bash
git clone https://github.com/theaathish/Url_shorter.git
cd Url_shorter
```

### **3️⃣ Install Dependencies**  
```bash
go mod tidy
```

### **4️⃣ Run Redis (If Not Running)**  
```bash
redis-server
```

### **5️⃣ Start the URL Shortener**  
```bash
go run main.go
```
Server will start at: `http://localhost:8080` 🚀  

---

## 📌 API Endpoints  

### **🔹 Shorten a URL**  
**POST** `/shorten`  
- **Request Body (JSON)**  
  ```json
  {
    "original_url": "https://example.com",
    "custom_alias": "myalias",
    "expiry": 3600
  }
  ```
- **Response (JSON)**  
  ```json
  {
    "short_url": "http://localhost:8080/myalias"
  }
  ```

### **🔹 Redirect a Shortened URL**  
**GET** `/:shortURL`  
- Example: `http://localhost:8080/myalias`  
- Redirects to `https://example.com`  

---

## ⚡️ Example Usage (cURL)  

```bash
curl -X POST http://localhost:8080/shorten \
     -H "Content-Type: application/json" \
     -d '{"original_url": "https://google.com", "expiry": 3600}'
```

---

## 🔧 To-Do List  
- [ ] Add Click Tracking 📊  
- [ ] Generate QR Codes for Short URLs 🔗  
- [ ] Deploy to a Cloud Server ☁️  

---

## 👨‍💻 Author  
**Aathish**  
🔗 GitHub: [theaathish](https://github.com/theaathish)  

