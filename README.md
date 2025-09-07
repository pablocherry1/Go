# Go for SRE – Learning Journey

This repository documents my progress through a Go for SRE learning plan. It contains projects and code from each week, organized for spaced repetition and practice.

## 🧠 Spaced Repetition Schedule

To ensure I retain what I've learned, I will review the topics and practice the code at specific intervals.

  * **Initial Review:** 1 day after learning.
  * **Second Review:** 3 days after the initial review.
  * **Third Review:** 7 days after the second review.
  * **Fourth Review:** 30 days after the third review.

### Review Calendar

  * **Week 1 (Fundamentals):** Review on Day 1, Day 3, Day 7, Day 30.
  * **Week 2 (Error Handling & Modules):** Review on Day 1, Day 3, Day 7, Day 30.
  * **Week 3 (Concurrency):** Review on Day 1, Day 3, Day 7, Day 30.
  * **Week 4 (Files & Data):** Review on Day 1, Day 3, Day 7, Day 30.

## 📝 Practice & Test Your Knowledge

You can copy and paste the following XML prompt into chat to request a quiz on the covered topics. The quiz will ask you to fill in the blanks in code and answer both open-ended and multiple-choice questions one at a time.

```xml
<quiz_request>
<first_question>week_selector</first_question>

<weeks>
    - 📘 Week 1 – Fundamentals of Go
    - ⚠️ Week 2 – Error Handling & Modules
    - 🧵 Week 3 – Concurrency in Go
    - 📂 Week 4 – Working with Files, JSON, YAML
    - 🌐 Week 5 – HTTP Clients & APIs
    - 🖥️ Week 6 – Building HTTP Servers
    - 🛠️ Week 7 – CLI Tools for SRE Work
    - 🚀 Week 8 – Capstone Project
  </weeks>

<topics>
## 📘 Week 1 – Fundamentals of Go

- [x]  Understand variables, constants, and types
- [x]  Learn functions, loops (`for`), and conditionals (`if`, `switch`)
- [x]  Understand Go file/project structure
- [x]  Project: CLI program that prints a name
- [x]  Project: Basic calculator (add, subtract, multiply, divide)

---

## ⚠️ Week 2 – Error Handling & Modules

- [x]  Learn `error` type and `if err != nil` pattern
- [x]  Understand `fmt.Errorf` and `errors.New`
- [x]  Use `go mod` to initialize a module
- [x]  Project: Word counter that reads a file and handles missing file errors

---

## 🧵 Week 3 – Concurrency in Go

- [x]  Learn `goroutines` and `channels`
- [x]  Use `select` with multiple channels
- [x]  Understand `sync.WaitGroup` and `sync.Mutex`
- [x]  Project: Ping multiple URLs in parallel and show response times
- [x]  Project: Simulate log processor with concurrent handlers

---

## 📂 Week 4 – Working with Files, JSON, YAML

- [x]  Learn to read/write files using `os` and `io`
- [x]  Parse JSON with `encoding/json`
- [x]  Parse YAML using `gopkg.in/yaml.v2`
- [x]  Project: Read `config.json` and print formatted output
- [x]  Project: Extract targets from a `prometheus.yml` file

---

## 🌐 Week 5 – HTTP Clients & APIs

- [ ]  Use `http.Get`, `http.Post`, and custom headers
- [ ]  Parse API responses in JSON
- [ ]  Understand response codes and error handling
- [ ]  Project: Status checker for a list of URLs
- [ ]  Project: API health checker with output to file

## 🖥️ Week 6 – Building HTTP Servers

- [ ]  Build HTTP server with `net/http`
- [ ]  Create multiple routes and handlers
- [ ]  Return JSON responses from endpoints
- [ ]  Project: Server with `/healthz` returning `{"status":"ok"}`
- [ ]  Project: Static `/metrics` endpoint in Prometheus format

---

## 🛠️ Week 7 – CLI Tools for SRE Work

- [ ]  Learn to use `flag` package for arguments
- [ ]  Parse input from `os.Args` or `os.Stdin`
- [ ]  Install and use `cobra` for CLI tooling
- [ ]  Project: Tail and filter log file by severity
- [ ]  Project: CLI that builds `kubectl logs` command based on input

---

## 🚀 Week 8 – Capstone Project

- [ ]  Serve HTTP on `/healthz` and `/metrics`
- [ ]  Perform health checks on URLs from YAML config
- [ ]  Return Prometheus-style metrics on availability
- [ ]  Write logs to a file
- [ ]  Accept CLI flags like `-config path/to/file.yaml`

</topics>
<quiz_format>
- fill_in_the_blanks: true
- open_questions: true
- multiple_choice: true
</quiz_format>
</quiz_request>
```