# Hello Kubernetes

This repository contains a simple web application built with GoLang and Docker, along with instructions for deploying it on a Kubernetes cluster. The application serves a "Hello, Kubernetes!" message and provides a hands-on introduction to working with Kubernetes and containerized applications.

## Getting Started

To get started with this project, follow these steps:

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/pratikjagrut/hello-kubernetes.git
   ```

2. Navigate to the cloned repository directory:

   ```bash
   cd hello-kubernetes
   ```

## Building and Running the Application

Follow these steps to build and run the application using Docker:

1. Build the Docker image for the application:

   ```bash
   docker build -t github.com/pratikjagrut/hello-kubernetes .
   ```

2. Run a Docker container from the built image:

   ```bash
   docker run -p 8080:8080 github.com/pratikjagrut/hello-kubernetes
   ```

3. Open a web browser and visit [http://localhost:8080](http://localhost:8080) to see the "Hello, Kubernetes!" message.

## Deploying on Kubernetes

To deploy the application on a Kubernetes cluster, you can follow the instructions provided in the blog post associated with this repository.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Feel free to enhance and customize this README to include more detailed information, links, and additional instructions as needed. This version provides a starting point for users who visit your repository.