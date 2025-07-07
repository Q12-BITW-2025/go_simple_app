# Go Hello World Web Application

This project is a simple web application written in Go that responds with "Hello, World!" when accessed.

## Project Structure

```
go-hello-world
├── src
│   └── main.go
├── Dockerfile
├── go.mod
└── README.md
```

## Prerequisites

- Go installed on your machine (for building the application)
- Docker installed on your machine (for building and running the Docker image)

## Building the Docker Image

To build the Docker image for this application, navigate to the project directory and run the following command:

```
docker build -t go-hello-world .
```

## Running the Docker Container

After building the image, you can run the Docker container with the following command:

```
docker run -p 8080:8080 go-hello-world
```

This command maps port 8080 of the container to port 8080 on your host machine.

## Accessing the Application

Once the container is running, you can access the application by navigating to `http://localhost:8080` in your web browser. You should see the message "Hello, World!" displayed on the page.