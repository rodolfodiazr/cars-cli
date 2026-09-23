# 🚗 Cars CLI

A command-line client for interacting with the [Cars API](https://github.com/rodolfodiazr/cars-api).

This project provides a simple way to interact with the API directly from the terminal, without requiring tools such as Postman.

## 🛠️ How to Run

### 1. Clone the repository

```bash
git clone https://github.com/rodolfodiazr/cars-cli.git
cd cars-cli
```

### 2. Run the CLI

```bash
go run ./cmd/cars
```

## 📋 Available Commands

```text
cars list
cars get <id>
cars create
cars update <id>
cars delete <id>
```

The CLI communicates with the Cars API through HTTP requests.
