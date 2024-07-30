# go-package-master
A modular Go application that calculates the optimal combination of packs to fulfil customer orders.

## Installation

1. Clone the repository:

```sh
cd go-package-master
```
2. Install Dependencies
```sh
go mod tidy
```

3. Update the config.json file if necessary to adjust pack sizes.
   sample config.json
```json
  {
      "packs": [
          { "size": 250 },
          { "size": 500 },
          { "size": 1000 },
          { "size": 2000 },
          { "size": 5000 }
      ]
  }
```
4. Running the Application
Start the application by running:

```sh
cd cmd
go build .
```

## API
Calculate Packs (v1)
URL: /api/v1/calculate
Method: POST

Request Body:
```json

{
  "items": 12001
}
```
Response:
```json
[
  {"size": 5000},
  {"size": 5000},
  {"size": 2000},
  {"size": 250}
]
```
Testing
Run unit and feature tests using:

```sh
cd tests
go test .
```

## Example
To test the API, you can use curl:

```sh
curl -X POST http://localhost:8080/api/v1/calculate \
  -H "Content-Type: application/json" \
  -H "Accept: application/json" \
  -d '{"items": 12001}'

```

Expected response:

```json
[
  {"size": 5000},
  {"size": 5000},
  {"size": 2000},
  {"size": 250}
]
```

## Considerations

- **Treblle SDK Integration:**  
  The project is integrated with the Treblle SDK for API operations, for API ops. Detailed instructions for setting up the Treblle SDK can be found in the `Dockerfile`.

- **Versioned Endpoints:**  
  Endpoints have been versioned to ensure backward compatibility and support for future enhancements etc.

- **GitHub Actions for CI/CD:**  
  The project uses GitHub Actions to automate testing, building, and deploying the Docker image to Docker Hub. This ensures a streamlined and reliable deployment process.

For more details on configuration and setup, please refer to the relevant sections in the `Dockerfile` and GitHub Actions workflow.

License
MIT
