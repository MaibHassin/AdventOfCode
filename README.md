## 🎄 Advent of Code Solutions

Welcome to my Advent of Code solutions repository! This project houses my solutions to the annual Advent of Code programming challenges, implemented across multiple languages.

---

## 📁 Repository Structure

This repository is organized by language and year, making it easy to navigate through solutions.

```
AdventOfCode/
├── Go/
│   ├── 2015/
│   │   ├── Day01/
│   │   │   ├── input.txt
│   │   │   ├── solution.go
│   │   │   └── solution_alt_approach.go  # Example of an alternative solution
│   │   └── 2016/
│   │       └── Day01/
│   │           └── ...
├── JavaScript/
│   ├── 2015/
│   │   └── Day01/
│   │       └── ...
└── Python/
    ├── 2015/
    │   └── Day01/
    │       └── ...
    └── 2023/
        └── Day01/
            ├── input.txt
            └── solution.py
```

- `Go/`, `JavaScript/`, `Python/`: Top-level folders for each programming language used.
- `YYYY/`: Inside each language folder, years are organized by their four-digit representation (e.g., `2015`, `2023`).
- `DayDD/`: Within each year, individual day solutions are stored in folders like Day01, Day02, etc.
- `input.txt`: This file contains the puzzle input for the specific day. (Note: These are not committed to Git for privacy/simplicity.)
- `solution.<ext>`: Your primary solution file for the day.
- `solution\_<name>.<ext>`: (Optional) Additional files for alternative approaches, optimizations, or different versions of a solution.

---

## 🚀 How to Run Solutions

To run a specific solution, navigate to its respective day's folder and execute the code using the appropriate language runtime.

#### Example (Go):

```bash
cd Go/2023/Day01
go run solution.go
```

---

## 🌱 My Approach & Workflow

- **Single Repository:** All solutions across all languages and years reside in this unified repository for easy management and a holistic view.
- **Feature Branching:** When starting a new day's puzzle or implementing an optimized solution for an existing one, I create a dedicated feature branch (e.g., `feat/2023-py-day05`, `opt/2022-go-day10`).
- **Modular Solutions:** Each day's solution typically resides in its own dedicated folder.
- **Version Control for Solutions:** For problems where I develop multiple approaches (e.g., a brute-force vs. an optimized dynamic programming solution), I often keep these as separate files within the day's directory (e.g., `solution_brute_force.py`, `solution_dp.py`) to preserve the learning process.

---

## 🌟 Goals

- To solve all Advent of Code puzzles.
- To practice and improve my skills in `C`, `C++`, `Go`, `JavaScript`, `TypeScript`, `Dart`, `Lua` and `Python`.
- To explore different algorithms and data structures.
- To create clean, efficient, and well-documented solutions.

---

Feel free to explore the solutions! If you have any questions or suggestions, don't hesitate to reach out.

Happy coding!
