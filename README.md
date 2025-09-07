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
<topics>
- Golang Fundamentals (variables, constants, types, functions, loops, and conditionals)
- Golang Project Structure
- Error Handling (`if err != nil`, `fmt.Errorf`, and `errors.New`)
- Go Modules (`go.mod`, `go.sum`)
- Concurrency (goroutines, channels, `sync.WaitGroup`, `sync.Mutex`, `select`)
- File I/O (`os`, `io`)
- Structured Data Parsing (JSON, YAML)
</topics>
<quiz_format>
- fill_in_the_blanks: true
- open_questions: true
- multiple_choice: true
</quiz_format>
</quiz_request>
```