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
5. 
Running the Application
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
curl -X POST http://localhost:8080/api/v1/calculate -d '{"items": 12001}' -H "Content-Type: application/json"
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

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

License
MIT
