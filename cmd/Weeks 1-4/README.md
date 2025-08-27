# Go for SRE – Learning Journey (Weeks 1-4)

This folder contains all the projects and code from the first four weeks of my Go learning journey. Each project is organized by week and topic, aligning with the learning schedule below.

## 📘 Week 1 – Fundamentals of Go

* **Topics:** Variables, constants, types, functions, loops, and conditionals.
* **Projects:**
    * **`cmd/hello-cli`**: A basic CLI program to print a name, demonstrating `fmt.Println` and `os.Args`.
    * **`cmd/calc-cli`**: A simple calculator, demonstrating the `flag` package for command-line arguments and `switch` statements for logic.

## ⚠️ Week 2 – Error Handling & Modules

* **Topics:** Go modules, error types, `if err != nil`, `fmt.Errorf`, and `errors.New`.
* **Projects:**
    * **`cmd/word-counter`**: A CLI tool that reads a file and counts words, focusing on robust error handling.

## 🧵 Week 3 – Concurrency in Go

* **Topics:** Goroutines, channels, `sync.WaitGroup`, `sync.Mutex`, and `select`.
* **Projects:**
    * **`cmd/ping-urls`**: A program that pings multiple URLs concurrently, demonstrating goroutine synchronization and communication.

## 📂 Week 4 – Working with Files, JSON, YAML

* **Topics:** Reading and writing files with `os`, parsing JSON with `encoding/json`, and parsing YAML with `gopkg.in/yaml.v2`.
* **Projects:**
    * Code demonstrating how to read and write files, parse a `config.json` file, and extract data from a `prometheus.yml` file.

## Quiz Request Prompt

You can copy and paste the following XML prompt into our chat to request a quiz on these topics at any time:

```xml
<quiz_request>
<topics>
- Understand variables, constants, and types
- Learn functions, loops (for), and conditionals (if, switch)
- Understand Go file/project structure
- Error handling with `if err != nil`, `fmt.Errorf`, and `errors.New`
- Go modules (`go.mod`, `go.sum`)
- Concurrency (goroutines, channels, `sync.WaitGroup`, `sync.Mutex`, `select`)
- File I/O (`os`, `io`)
- Structured data parsing (JSON, YAML)
</topics>
</quiz_request>