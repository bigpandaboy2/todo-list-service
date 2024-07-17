# Todo List Microservice

## Project Description
This microservice provides a RESTful API for managing a Todo List. It provides endpoints for creating, updating, deleting, marking as done, and listing tasks.

## Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/todo-list-service.git
   cd todo-list-service
   ```

2. Build and run the project using Docker Compose:
   ```sh
   make build
   make run
   ```

## Deployment on RENDER
URL: https://todo-list-service-39p6.onrender.com

## API Usage
### Create a New Task
- **Request: POST /api/todo-list/tasks**
- **Request Body**:
  ```json
  {
        "title": "Buy a book",
        "activeAt": "2023-08-04"
  }

## Update an Existing Task

- **Request: PUT /api/todo-list/tasks/{id}**
- **Request Body:**:

    ```sh
    {
        "title": "Buy a book - High Load Applications",
        "activeAt": "2023-08-05"
    }

    ```sh
## Delete a Task
- **Request: DELETE /api/todo-list/tasks/{id}**

## Mark a Task as Done
- **Request: PUT /api/todo-list/tasks/{id}/done**

## List Tasks by Status
- **Request: GET /api/todo-list/tasks?status=active**
- **Request: GET /api/todo-list/tasks?status=done**

## Health Check
- **Request: GET /health**
