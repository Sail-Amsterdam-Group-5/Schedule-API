# Schedule-API

### 1 - Prerequisite
- install azurite on your computer.
    - via npm: npm install -g azurite
- duplicate the `.env-example` and rename it to `.env`
    - optional: change the contents in the `.env` file

### 2 - Start the application
- start azurite in cmd with the command `azurite`
- start the application in cmd with the command `go run main.go`

### 3 - Add dummy data
    - do a `POST` request to the `localhost:8080/schedule/` endpoint