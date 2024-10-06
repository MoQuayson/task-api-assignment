# Long-Running Task JSON REST API

This project is a simple JSON REST API built using Golang that simulates long-running background tasks. It uses SQLite as the database and bearer token authorization for security. This API allows clients to:

- Initiate a long-running task via a `POST` request.
- Retrieve the status of the task.
- Fetch the output of the completed task.

The API is implemented using Golang's standard library, ensuring no third-party packages are required, except for the SQLite driver.

## Features

- **Bearer Token Authentication**: All endpoints are secured using a simple bearer token system.
- **Background Processing**: Tasks are processed asynchronously in the background.
- **Status & Output Retrieval**: Clients can check the status and retrieve the results of the task after completion.
- **SQLite Database**: Task status and results are stored in a lightweight, embedded SQLite database.

## Requirements

- Go 1.18 or later
- SQLite


## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/MoQuayson/task-api-assignment.git
   cd task-api-assignment

2. Install Go dependencies:
    ```bash 
   go mod tidy
3. Run the server:
    ```bash
   go run main.go
    
## Endpoints

1. Token Generation
   - **Url**: /api/auth/access-token
   - **Method**: `POST`
   - **Description**: Generates a token

2. Make Payment
   - **Url**: /api/payments
   - **Method**: `POST`
   - **Description**: Initiates a long-running payment task that runs in the background.
   - **Authorization**: `Bearer token`

3. Get Payment Status
   - **Url**: /api/payments/status/:transaction_id
   - **Method**: `GET`
   - **Description**: Retrieves the status of the payment
   - **Authorization**: `Bearer token`

4. Get Payment output
   - **Url**: /api/payments/completed/:transaction_id
   - **Method**: `POST`
   - **Description**: Initiates a long-running payment task that runs in the background.
   - **Authorization**: `Bearer token`


