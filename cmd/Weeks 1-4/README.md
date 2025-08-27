# 🚀 Go for SRE – Learning Journey (Weeks 1-4)
This folder contains all the projects and code from the first four weeks of my Go learning journey. Each project is organized by week and topic, aligning with the learning schedule below.

---
## 📘 Week 1 – Fundamentals of Go
**Topics:** Variables, constants, types, functions, loops, and conditionals.
**Projects:**
- **Week 1 - hello-cli:** A basic CLI program to print a name, demonstrating `fmt.Println` and `os.Args`.
- **Week 1 - calc-cli:** A simple calculator, demonstrating the `flag` package for command-line arguments and `switch` statements for logic.

---
## ⚠️ Week 2 – Error Handling & Modules
**Topics:** Go modules, error types, `if err != nil`, `fmt.Errorf`, and `errors.New`.
**Projects:**
- **Week 2 - wordCounter-cli:** A CLI tool that reads a file and counts words, focusing on robust error handling.

---
## 🧵 Week 3 – Concurrency in Go
**Topics:** Goroutines, channels, `sync.WaitGroup`, `sync.Mutex`, and `select`.
**Projects:**
- **Week 3 - ping-urls:** A program that pings multiple URLs concurrently, demonstrating goroutine synchronization and communication.
- **Week 3 - log-processor:** A project to simulate a log processor, demonstrating `sync.Mutex` and `select`.

---
## 📂 Week 4 – Working with Files, JSON, YAML
**Topics:** Reading and writing files, parsing JSON, and parsing YAML.
**Projects:**
- **Week 4 - write_file:** Your file I/O project.
- **Week 4 - JSON_Parser:** A program that reads a JSON file and parses the data.
- **Week 4 - YAML_Parser:** A program that reads a YAML file and extracts targets.

---
### 📝 Quiz Request Prompt
You can copy and paste the following XML prompt into our chat to request a quiz on these topics at any time:

```xml
<quiz_request>
<topics>
- Golang Fundamentals (variables, constants, types, functions, loops, and conditionals)
- Golang Project Structure
- Error Handling (`if err != nil`, `fmt.Errorf`, and `errors.New`)
- Go Modules (`go.mod`, `go.sum`)
- Concurrency (goroutines, channels, `sync.WaitGroup`, `sync.Mutex`, `select`)
- File I/O (`os`, `io`)
- Structured Data Parsing (JSON, YAML)
</topics>
</quiz_request>