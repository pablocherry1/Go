# HTTP Status Codes: Categories Explained

Instead of memorizing every status code, the most practical approach is to **understand their categories**. This allows you to quickly diagnose whether an issue comes from the client or the server.

---

## 📘 Categories of HTTP Status Codes

### 🔹 1xx – Informational
- The request has been received and the process is continuing.
- Example: `100 Continue`

---

### 🔹 2xx – Success
- The request was successfully received, understood, and accepted.  
- Example: `200 OK`

---

### 🔹 3xx – Redirection
- Further action is required to complete the request.  
- Example: `301 Moved Permanently`

---

### 🔹 4xx – Client Error
- The client (your program) made a mistake.  
- Examples:  
  - `400 Bad Request`  
  - `404 Not Found`

---

### 🔹 5xx – Server Error
- The server failed to fulfill a valid request.  
- Example: `500 Internal Server Error`

---

## ✅ Why Focus on Categories?
By knowing these categories, you can:
- Quickly identify if an error is **your responsibility** (4xx) or the **server’s responsibility** (5xx).  
- Understand the flow of communication between client and server without memorizing every single code.  

