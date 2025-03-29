# MarkFlow 🚀

MarkFlow is a simple and efficient Markdown file management system built using Go and Gin. It provides grammar checking, file uploads, and the ability to retrieve Markdown files as HTML.

## ✨ Features
- ✅ Grammar checking for Markdown files
- 📂 Upload Markdown files
- 🔍 Retrieve Markdown files as HTML
- 📜 List all uploaded files

## 📥 Installation
### Prerequisites
- 🛠 Go 1.24.1 or higher
- 🌐 cURL (for API requests)

#### Install Go using Pacman (Arch Linux)
```sh
sudo pacman -S go
```

#### Install cURL
```sh
sudo pacman -S curl
```

### Clone the Repository
```sh
git clone https://github.com/Abhishek2010dev/MarkFlow.git
cd MarkFlow
```

### Install Dependencies
```sh
go mod tidy
```

### Run the Application
```sh
go run main.go
```

## 📡 API Endpoints
### 📝 Check Grammar
**POST** `/check-grammer`
- **Request:**
  - 📎 Form-data: `file` (Markdown file)
- **Response:** JSON object with grammar suggestions

**cURL Example:**
```sh
curl -X POST http://localhost:8080/check-grammer \
  -F "file=@test.md" \
  -H "Content-Type: multipart/form-data" | jq
```

### 📤 Upload Markdown File
**POST** `/notes`
- **Request:**
  - 📎 Form-data: `file` (Markdown file)
- **Response:** ✅ Confirmation of file upload

**cURL Example:**
```sh
curl -X POST http://localhost:8080/notes \
  -F "file=@test.md" \
  -H "Content-Type: multipart/form-data"
```

### 📄 Retrieve Markdown as HTML
**GET** `/notes/{filename}`
- **Response:** 🖥 HTML version of the Markdown file

**cURL Example:**
```sh
curl http://localhost:8080/notes/test.md
```

### 📜 List All Files
**GET** `/notes`
- **Response:** 📄 JSON object listing uploaded files

**cURL Example:**
```sh
curl http://localhost:8080/notes
```

## 📦 Dependencies
MarkFlow uses the following dependencies:
- `github.com/gin-gonic/gin` - ⚡ HTTP web framework
- `github.com/yuin/goldmark` - 📝 Markdown parser
- `github.com/goccy/go-json` - 🚀 High-performance JSON library

## 🔗 Reference
For more inspiration on building a Markdown note-taking app, check out [this guide](https://roadmap.sh/projects/markdown-note-taking-app).

## 📜 License
This project is licensed under the MIT License.

## 👨‍💻 Author
[Abhishek2010dev](https://github.com/Abhishek2010dev)


