# go-package-master
A modular Go application that calculates the optimal combination of packs to fulfil customer orders for Gymshark 🦈.

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
{"2000":1,"250":1,"5000":2}
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
  -H "Authorization: Bearer GSLC-123-0R" \ 
  -d '{"items": 12001}'

// Note: the authorization  token is something that ive kept as a constant for now with  value as shown above
```

Expected response:

```json
{"2000":1,"250":1,"5000":2}
```

## Considerations
- **Versioned Endpoints:**  
  Endpoints have been versioned to ensure backward compatibility and support for future enhancements etc.

- **Treblle SDK Integration:**  
 The project is integrated with the Treblle SDK for API operations and for API ops.  For more details on the configuration and setup of the Treblle SDK, please refer to the relevant sections in the `Dockerfile` and GitHub Actions workflow..

- **GitHub Actions for CI/CD:**  
  The project uses GitHub Actions to automate testing, building, and deploying the Docker image to Docker Hub. This ensures a streamlined and reliable deployment process.


License
MIT
