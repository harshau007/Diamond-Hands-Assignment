# List Manipulation API

## Overview

This project implements a thread-safe integer list manager in Go, exposed via a Gin-powered HTTP API. Inputs are appended if they match the sign of the list’s first element (or list is empty/zero), or used to reduce existing entries when of opposite sign.

## Project Structure

```
listmanagerapi/
├── go.mod
├── go.sum
├── cmd/
│   └── main.go
├── internal/
│   ├── listmanager/
│   │   ├── listmanager.go
│   │   └── listmanager_test.go
│   ├── models/
│   │   └── models.go
│   └── router/
│       ├── router.go
│       └── router_test.go
└── README.md
```

## Running the API

1. **Install dependencies**
   ```bash
   go mod tidy
   ```
2. **Start the server**
   ```bash
   go run cmd/main.go
   ```
   The server listens on `:8080`.

## API Endpoints

- `POST /add` : Add a number. Request JSON `{ "number": <int> }`. Returns the updated list.
- `GET  /list` : Get current list.
- `POST /reset`: Reset list to empty.

## Example Output

For the sequence of inputs: `[5, 10, -6]`:

```
> POST /add {"number":5}
[5]
> POST /add {"number":10}
[5,10]
> POST /add {"number":-6}
[9]
```

## Video Demo

[![Watch the video](https://github.com/harshau007/Diamond-Hands-Assignment/raw/refs/heads/main/assets/listmanagerapi.mp4)](https://github.com/harshau007/Diamond-Hands-Assignment/raw/refs/heads/main/assets/listmanagerapi.mp4)

## Edge Cases

- **Empty List**: First input always added.
- **Zero Handling**: Zero matches any sign, and is appended.
- **Complete Elimination**: Negative input magnitude ≥ sum of absolutes empties the list.
- **Partial Reduction**: Negative input reduces front elements partially.
- **Over-reduction**: Excess negative input drops all elements, resulting in empty list.

## Testing

Run all unit and API tests:

```bash
go test ./...
```

Benchmarks are included in the `router` and `listmanager` packages.
