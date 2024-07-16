# Simple HTTP Server in Go

This is a basic HTTP server implemented in Go. It serves static files from the local directory and listens on a specified port.

## Features

- Serves static files (e.g., HTML, CSS, JavaScript) from the local directory.
- Handles GET requests and returns the requested file.
- Supports custom port specification via command-line flag.

## Usage

1. **Clone the repository:**

   ```sh
   git clone https://github.com/edaywalid/http-server-go.git
   cd http-server-go
   ```

2. **Build the server:**

   ```sh
   go build -o server
   ```

3. **Run the server:**

   ```sh
   ./server
   ```

   By default, the server listens on port `8080`. You can specify a custom port using the `-p` flag:

   ```sh
   ./server -p 8081
   ```

   Replace `8081` with the desired port number.

4. **Access the server:**

   Open a web browser and navigate to `http://localhost:8080/file` (or the custom port you specified). Replace `file` with the name of the file you want to access.

## HOW IT WORKS

- the server first listen to incoming request on the specified port
- the server establish a connection with the client after the client send a request
- the server read the request and extract the file name from the request
- the server then read the file and send it to the client
- the client then display the file to the user
