# Receipt Processor Challenge (in Go this time)

## Summary

**Objective:** Build a web service that implements the specified API for processing receipts.

This API was built to fulfill a code challenge for [Fetch Rewards](https://fetch.com/). For detailed instructions and rules, refer to the [fetch-rewards/receipt-processor-challenge](https://github.com/fetch-rewards/receipt-processor-challenge) repository.

**Note**: This is my second time building this API. Since I was unfamiliar with Go before starting this sproject, I [built it with TypeScript](https://github.com/derekvmcintire/receipt-processor) the first time around.

---

### **Technology Used**

- [Go](https://go.dev/): Written with Go
- [Gin](https://gin-gonic.com/): HTTP web framework for Go, used to define routes, handle HTTP requests, and build RESTful APIs.
- [Docker](https://www.docker.com/): The app is available in a Docker container via `make docker-run`.
- [Testify](https://pkg.go.dev/github.com/stretchr/testify): A set of Go testing utilities that simplify writing unit tests with assertions, mocking, and suite-based testing.
- **In-memory Storage**: The application uses an in-memory map for data storage.

---

## API Endpoints

### 1. **Process Receipt**

- **Path**: `/receipts/process`
- **Method**: `POST`
- **Request Payload**:
  The request should contain a JSON object representing a receipt.
  Example:

  ```json
  {
    "retailer": "Target",
    "purchaseDate": "2022-01-01",
    "purchaseTime": "13:01",
    "items": [
      {
        "shortDescription": "Mountain Dew 12PK",
        "price": "6.49"
      },
      {
        "shortDescription": "Emils Cheese Pizza",
        "price": "12.25"
      },
      {
        "shortDescription": "Knorr Creamy Chicken",
        "price": "1.26"
      },
      {
        "shortDescription": "Doritos Nacho Cheese",
        "price": "3.35"
      },
      {
        "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
        "price": "12.00"
      }
    ],
    "total": "35.35"
  }
  ```

- **Response**:
  The response will contain the ID of the processed receipt. This ID will be used to retrieve points associated with the receipt later.
  Example:

  ```json
  {
    "id": "7fb1377b-b223-49d9-a31a-5a02701dd310"
  }
  ```

- **Description**:
  This endpoint processes a receipt and generates an ID for it. The receipt data (e.g., store name, item prices) is processed in-memory, and the receipt ID is returned. The number of points awarded is determined based on the receipt's content.

---

### 2. **Get Points for Receipt**

- **Path**: `/receipts/{id}/points`
- **Method**: `GET`
- **Path Parameter**:

  - `id`: The unique identifier of the receipt, which was returned when the receipt was processed using the `/receipts/process` endpoint.

- **Response**:
  The response will contain the number of points awarded for the receipt with the provided ID.
  Example:

  ```json
  {
    "points": 32
  }
  ```

- **Description**:
  This endpoint retrieves the points awarded for a particular receipt. The points are calculated based on the rules specified in the code.

---

## Instructions for Running the Application

### Prerequisites

- **Go (Golang)**: This application is written in Go, so you need to have Go installed on your system to run it.
- **Docker** (optional): If you prefer to run the application in a containerized environment, Docker can be used.

### Running the Application

1. **Clone the Repository**:
   First, clone the repository from your source control provider (e.g., GitHub).

   ```bash
   git clone https://github.com/your-username/go-receipt-processor.git
   cd go-receipt-processor
   ```

2. **Install Dependencies**:
   Run the following Go command to download the necessary dependencies:

   ```bash
   go mod tidy
   ```

3. **Run the Application**:
   You can run the application locally using the following command:

   ```bash
   make run
   ```

   By default, the application will start an HTTP server on port `8080`. You should see output like:

   ```
   Listening and serving HTTP on :8080
   ```

4. **Accessing the API**:
   - The **Process Receipt** endpoint will be available at `POST http://localhost:8080/receipts/process`.
   - The **Get Points** endpoint will be available at `GET http://localhost:8080/receipts/{id}/points`, where `{id}` is the receipt ID you receive after processing a receipt.

---

### Running with Docker

If you prefer running the application inside a Docker container, follow these steps:

1. **Build the Docker Image**:
   First, create the Docker image using the provided `Dockerfile`:

```bash
make docker-build
```

2. **Run the Docker Container**:
   Start the application in a Docker container:

```bash
make docker-run
```

3. **Access the API**:
   Once the Docker container is running, you can access the API the same way as if you were running it locally, using `http://localhost:8080`.

---

### Makefile Commands

A `Makefile` is included for convenience, providing shortcuts for common tasks:

| **Target**     | **Description**                                                  | **Command**                                 |
| -------------- | ---------------------------------------------------------------- | ------------------------------------------- |
| `fmt`          | Format Go source code using `go fmt`                             | `go fmt ./...`                              |
| `build`        | Build the Go project                                             | `go build ./...`                            |
| `test`         | Run tests using `go test`                                        | `go test ./...`                             |
| `run`          | Run the Go application (`cmd/api/main.go`)                       | `go run cmd/api/main.go`                    |
| `docker-build` | Build the Docker image for the project                           | `docker build -t receipt-processor .`       |
| `docker-run`   | Run the Docker container for the application (exposes port 8080) | `docker run -p 8080:8080 receipt-processor` |

### Example Usage

- **Format Go Code**:

```bash
make fmt
```

- **Build the Go Project**:

```bash
make build
```

- **Run Tests**:

```bash
make test
```

- **Run the Application**:

```bash
make run
```

- **Build Docker Image**:

```bash
make docker-build
```

- **Run Docker Container**:

```bash
make docker-run
```

---

## Example Requests

### Example 1: Process Receipt

```bash
curl -X POST http://localhost:8080/receipts/process \
    -H "Content-Type: application/json" \
    -d '{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}'
```

**Response**:

```json
{
  "id": "7fb1377b-b223-49d9-a31a-5a02701dd310"
}
```

### Example 2: Get Points for Receipt

```bash
curl http://localhost:8080/receipts/7fb1377b-b223-49d9-a31a-5a02701dd310/points
```

**Response**:

```json
{
  "points": 109
}
```

---

## Code Structure

### Main File (`cmd/api/main.go`)

The `main.go` file serves as the entry point for the application. It sets up the HTTP routes and starts the server. It uses the Gin framework to handle incoming requests and dependencies are managed through a custom container (`internal/container`).

### Dockerfile

If using Docker, there is a `Dockerfile` that specifies the base image and how the application should be built and run inside the container.

---

## Tests

Tests for the application can be run using the `go test` command. The code includes test cases to validate the functionality of the receipt processing and point calculation logic.

To run the tests, simply execute:

```bash
make test
```

---

## Notes

- The application does not persist data across restarts. Once the application stops, all data (receipts and points) are lost.

---

## Appilcation Architecture:

This API was designed following Hexegonal Architecture principles.

### **Hexagonal Architecture Overview**

Hexagonal Architecture, also known as the **Ports and Adapters Architecture**, is a design pattern that emphasizes separation of concerns, making applications more maintainable, testable, and adaptable to change. The architecture organizes the application into three main layers:

1. **Core (Business Logic):**

   - Contains the application's domain models and business rules.
   - Isolated from external frameworks, libraries, or dependencies.
   - Example: `domain/receipt.go` and `application/receipt_service.go`.

2. **Ports (Interfaces):**

   - Define abstractions for how the application interacts with external systems or internal business logic.
   - Ports act as boundaries, allowing the core to remain agnostic of specific implementations.
   - Example: `ports/repository/receipt_repository.go` and `ports/http/response/get_receipt_points_response.go`.

3. **Adapters (Implementations):**
   - Implement the port interfaces to connect the core with external systems (e.g., databases, APIs, user interfaces).
   - Adapters translate between the external systems and the application's core.
   - Example: `adapters/http/get_receipt_points_handler.go` and `adapters/memory/receipt_store.go`.

### **Key Benefits**

- **Flexibility:** Easily swap out or modify adapters (e.g., replace an in-memory repository with a database implementation) without changing core logic.
- **Testability:** The core logic can be tested in isolation using mock adapters.
- **Maintainability:** Clear separation of concerns reduces complexity and coupling.

---

## Folder Structure

```
/receipt-processor
│
├── cmd/
│ ├── api/
│ │   └── main.go
│ ├── container/
│ │   └── container.go
│ │   │
├── internal/
│ ├── adapters/
│ │   ├── http/
│ │   │   └── get_receipt_points_handler.go
│ │   │   └── receipt_process_handler.go
│ │   ├── memory/
│ │   │   └── receipt_store.go
│ ├── application/
│ │   └── points_calculator_rules.go
│ │   └── points_calculator.go
│ │   └── receipt_service.go
│ ├── domain/
│ │   └── receipt.go
│ ├── ports/
│ │   ├── core/
│ │   │   └── points_calculator.go
│ │   │   └── points_rules.go
│ │   │   └── receipt_service.go
│ │   ├── http/
│ │   │   └── response/
│ │   │       └── get_receipt_points_response.go
│ │   │       └── process_receipt_response.go
│ │   ├── repository/
│ │   │   └── receipt_repository.go
│ │   │
├── pkg/
│ └── utils/
│ │   └── receipt_date_time.go
│ │   │
├── test/
│ ├── application/
│ │   ├── rules/
│ │   │   └── item_count_rule_test.go
│ │   │   └── ...remaining rule tests
│ │   ├── points_calculator_test.go
│ │   ├── receipt_service_test.go
│ ├── adapters/
│ │   ├── http/
│ │   │       └── get_receipt_points_handler_test.go
│ │   │       └── receipt_process_handler_test.go
│ ├── local_mocks/
│ │   └── mock_receipt_service.go
│ │   └── mock_points_calculator.go
│ ├── memory/
│ │   └── receipt_store_test.go

```
