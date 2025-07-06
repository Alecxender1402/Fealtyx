**Student CRUD API with Ollama AI Integration** A simple, concurrent REST API in Go for managing students, featuring in-memory CRUD operations and AI-powered profile summaries using Ollama LLMs.

---

## 🚀 Features

- Create, Read, Update, Delete (CRUD) operations for student records
- In-memory data storage (no database required)
- Concurrent safe with mutex locks
- AI-generated student summaries via Ollama LLM API
- Input validation and robust error handling
- Easy to test with Postman or curl

---

## 🛠️ Tech Stack

- **Backend:** GO(Golang), Gorilla Mux (HTTP router)
- **AI Integration:** Ollama (local LLM server)
- **API Testing:** curl, Postman

---

## 📦 Getting Started

### 1. Clone the Repository
```bash
git clone https://github.com/Alecxender1402/Fealtyx.git
cd Fealtyx
```

### 2. Backend Setup
```bash
go mod tidy
```

### 3. Pull an Ollama Model and Start Ollama
```bash
ollama serve
ollama pull llama3.2:1b
```

### 4. Run the API Server
```bash
go run .
```

The server will start at:
``` http://localhost:8080 ```

---

## 🔑 API Endpoints

| Method | Endpoint                                 | Description                           |
|--------|------------------------------------------|---------------------------------------|
| POST   | `/students`                              | Create a new student                  |
| GET    | `/students`                              | List all students                     |
| GET    | `/students/{id}`                         | Get a student by ID                   |
| PUT    | `/students/{id}`                         | Update a student by ID                |
| DELETE | `/students/{id}`                         | Delete a student by ID                |
| GET    | `/students/{id}/summary`                 | Generate AI summary for student       |  

---

