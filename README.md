# 📝 Task CLI

A simple command-line **task manager** built with Go.  
You can add, update, delete, and list tasks stored in a JSON file.

---

## 🚀 Features

- Add new tasks
- Delete tasks
- Update task description
- Mark tasks as **in-progress** or **done**
- List tasks (with optional filter by status)
- JSON file storage (no database required)

---

## 📦 Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/task-cli.git
   cd task-cli


2. Build the binary:

   ```bash
   go build -o task
   ```

3. Run the CLI:

   ```bash
   ./task <command> [args]
   ```

---

## 🛠 Usage

### Add a new task

```bash
./task add "Buy groceries"
```

### Delete a task by ID

```bash
./task delete 0
```

### Update a task description

```bash
./task update 0 "Buy groceries and fruits"
```

### Mark a task as in-progress

```bash
./task mark-in-progress 0
```

### Mark a task as done

```bash
./task mark-done 0
```

### List all tasks

```bash
./task list
```

### List tasks filtered by status

```bash
./task list todo
./task list in-progress
./task list done
```

---

## 📂 Project Structure

```
.
├── main.go          # CLI entry point
├── task.go          # Task struct & logic
├── storage.go       # JSON read/write helpers
├── tasks.json       # Task storage (auto-created)
└── README.md        # Project documentation
```

---

## 🗒 Example

```bash
$ ./task add "Learn Go"
$ ./task add "Build CLI app"
$ ./task list

[0] Learn Go          | todo        | created: 2025-08-22
[1] Build CLI app     | todo        | created: 2025-08-22

$ ./task mark-in-progress 0
$ ./task list in-progress

[0] Learn Go          | in-progress | created: 2025-08-22
```

---

## 🤝 Contributing

1. Fork the project
2. Create a new branch (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -m "Add new feature"`)
4. Push to your branch (`git push origin feature/your-feature`)
5. Open a Pull Request 🚀

---
