# ter-todo

A terminal-based TODO manager for Linux, built with Go, [Bubble Tea](https://github.com/charmbracelet/bubbletea), and SQLite.

![screenshot](screenshot.png) <!-- Add a screenshot if you have one -->

## Features

- Add, edit, and delete tasks from your terminal
- Tasks stored in a local SQLite database
- Keyboard navigation and shortcuts
- Simple, clean TUI interface

## Installation

### Prerequisites

- Go 1.20+ installed
- Linux (tested on Ubuntu, Fedora, Arch)

### Build from Source

```sh
git clone https://github.com/fireab/ter-todo.git
cd ter-todo
go build -o ter-todo
./ter-todo
```

### Download Binary

> _Coming soon: Pre-built binaries for Linux will be available on the [Releases](https://github.com/fireab/ter-todo/releases) page._

## Usage

- **Tab**: Switch between task input, description input, and table
- **Enter**: Add a new task (when in input fields)
- **Left/Right**: Change task status (when table is focused)
- **Delete**: Remove selected task
- **Ctrl+C**: Quit

## Database

- Tasks are stored in `task.db` in the current directory.
- The database is created automatically on first run.

## Contributing

Pull requests and issues are welcome! Please open an issue for bugs or feature requests.

## License

[MIT](LICENSE)

---
