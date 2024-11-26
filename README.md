
# Task Manager

## Overview
The Task Manager is a simple RESTapi service for managing tasks. It supports creating, updating, deleting, and retrieving tasks. The service is implemented in Go.

## Features
- Add new tasks with required fields: `Title`, `Description`, and `Status`.
- Update existing task.
- Delete tasks.
- Get a task by ID or get list of tasks.

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/oshritmalach/task-manager.git
   ```
2. Navigate to the project directory:
   ```bash
   cd task-manager
   ```
3. Install dependencies (if any):
   ```bash
   go mod tidy
   ```

4. Run the application:
   ```bash
   go run main.go
   ```
## Testing
   Run the tests with:
   ```bash
    go test ./repository ./handler -v
   ```

## Using the Makefile

This project includes a `Makefile` to simplify common tasks. Below are the available commands and their usage:

1. **Build the Project**
   ```bash
   make build
   ```
2. **Run the Project**
   ```bash
   make run
   ```
3. **Test the Project**
   ```bash
   make test
   ```
3. **Lint the Project**
   ```bash
   make lint
   ```
   
   

## Usage

### API Endpoints
| Method | Endpoint               | Description         | Example Request Body                           |
|--------|-------------------------|---------------------|-----------------------------------------------|
| GET    | `/tasks`               | Get all tasks       | -                                             |
| POST   | `/task`                | Create a new task   | `{ "title": "Task 1", "description": "Test description", "status": "open" }` |
| GET    | `/task/{id}`           | Get a task by ID    | -                                             |
| POST   | `/task/{id}`           | Update a task by ID | `{ "title": "Updated Title" }`                |
| DELETE | `/task/{id}`           | Delete a task by ID | -                                             |

### Example Request and Response

#### Create a Task
**Request:**
```bash
curl -X POST http://localhost:8080/task \
-H "Content-Type: application/json" \
-d '{"title": "My Task", "description": "This is a test task", "status": "open"}'
```

**Response:**
```
Status Code: 201 Created
```
#### Update a Task
**Request:**
```bash
curl -X POST http://localhost:8080/task/1 \
-H "Content-Type: application/json" \
-d '{"status": "in_progress"}'
```

**Response:**
```
Status Code: 200 Ok
```
```json
{
   "title": "2 updated",
   "description": "This is a test task",
   "status": "in_progress",
   "createdAt": "2024-11-26T12:00:00Z"
}
```

#### Delete a Task
**Request:**
```bash
curl -X DELETE http://localhost:8080/task/1
```

**Response:**
```
Status Code: 204 No Content
```
#### Get a Task
```bash
curl -X GET http://localhost:8080/task/1
```
**Response:**
```
Status Code: 200 Ok
```
```json
{
   "title": "Item",
   "description": "3333",
   "status": "pending",
   "created_at": "2024-11-26T13:16:40.974907+02:00"
}
```
#### Get all Tasks
```bash
curl -X GET http://localhost:8080/tasks
```
**Response:**
```
Status Code: 200 Ok
```
```json
{
   "1": {
      "title": "Item",
      "description": "3333",
      "status": "pending",
      "created_at": "2024-11-26T13:16:40.974907+02:00"
   }
}
```

### Error Handling

All failed requests return a response in a unified JSON structure containing the following fields:

- **`code`**: The HTTP status code indicating the type of error (e.g., 400, 404, etc.).
- **`message`**: A descriptive message explaining the reason for the error.

#### Example Error Response:
```json
{
    "code": 404,
    "message": "task with ID 1 does not exist"
}
```



