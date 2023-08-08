# Web-Based GoLang Project with Docker

This is a simple example project that demonstrates how to create a web-based GoLang application and containerize it using Docker.

## Prerequisites

Before you start, make sure you have the following tools installed:

- [GoLang](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)

## Getting Started

1. **Clone the Repository**: Start by cloning this repository to your local machine:

   ```bash
   git clone https://github.com/your-username/go-web-app.git
   ```

   Replace `your-username` with your actual GitHub username.

2. **Navigate to the Project Directory**: Change to the project directory:

   ```bash
   cd go-web-app
   ```

3. **Edit `main.go` (Optional)**: You can modify the `main.go` file to customize the message displayed by the application.

4. **Build the Docker Image**: Build the Docker image using the following command:

   ```bash
   docker build -t go-web-app .
   ```

   This command tells Docker to build an image named `go-web-app` based on the `Dockerfile` in the current directory.

5. **Run the Docker Container**: Run a container from the created image:

   ```bash
   docker run -p 8080:8080 go-web-app
   ```

6. **Access the Web Application**: Open a web browser and navigate to `http://localhost:8080`. You should see the message "Hello, Kubernetes!" displayed.

7. **Stopping the Container**: To stop the running container, press `Ctrl + C` in the terminal where it's running.

## Additional Notes

- The `Dockerfile` specifies the base image as `golang:1.16-alpine` and sets up the working directory, copies the project files, builds the GoLang application, exposes port 8080, and defines the command to run the application.

- You can customize the port mapping in the `docker run` command. For example, you can use `-p 8000:8080` to map port 8000 on your machine to port 8080 in the container.

- Feel free to explore and expand upon this example by adding more functionality to your GoLang application and experimenting with Docker features.
