# WebSocket Server

This is an example of a WebSocket server written in Go. It utilizes the Gorilla WebSocket library to handle WebSocket connections.

## Features

The server has the following features:

- Accepts WebSocket connections on the "/ws" route.
- Allows clients to connect and send messages.
- Maintains a list of connected clients.
- Broadcasts received messages from a client to all other connected clients.

## Usage

1. Make sure you have Go installed on your system.
2. Download or clone this repository.
3. Open a terminal and navigate to the root directory of the project.
4. Run the following command to start the server:

   ```shell
   go run main.go
   ```

5. The server will be running on port 8080.
6. Open a web browser and access http://localhost:8080.
7. The web client will display a page with a text input field.
8. Open multiple browser tabs or windows to simulate multiple clients.
9. Type a message in the text box and press Enter.
10. The message will be sent to all other connected clients.

## Project Structure

- The `main.go` file contains the main logic of the WebSocket server.
- The `github.com/gorilla/websocket` package is used to enhance the functionality of the `net/http` package.
- The `HomeHandler` function handles the initial route and serves the `public/index.html` file.
- The `WsHandler` function is called when a client connects to the "/ws" route.
- The `AddClient` function adds the client to the list of connected clients.
- The `BroadcastMessageToAll` function sends a message to all connected clients.
- The `BroadcastMessageToClient` function sends a message to a specific client.
- The `BroadcastMessageToOthers` function sends a message to all clients except the specific client.
- The `HandleClientMessages` function handles messages received from a client and broadcasts them to other clients.
- The `CleanupClients` function checks client connections and removes disconnected clients from the list.
- The `MonitorClientConnections` function periodically checks client connections.

## Dependencies

This project depends on the Gorilla WebSocket library. You can install it using the following command:

```shell
go get github.com/gorilla/websocket
```

Make sure you have set up the Go environment correctly for the dependencies to be installed properly.

## Conclusion

This is a simple example of a WebSocket server in Go. It allows multiple clients to connect and exchange real-time messages. Feel free to explore and modify the code to fit your specific needs.