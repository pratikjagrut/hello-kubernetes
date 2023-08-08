# Hello Kubernetes
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

To deploy the application on a Kubernetes cluster using a private registry, proceed as follows:

1. Create an Image Pull Secret in Kubernetes to access your private registry:

   ```bash
   kubectl create secret docker-registry my-docker-secret \
     --docker-username=<your-username> \
     --docker-password=<your-password> \
     --docker-server=<your-registry-server>
   ```

2. Apply the deployment YAML to your Kubernetes cluster:

   ```bash
   kubectl apply -f hello-k8s-deployment.yaml
   ```

3. Use port-forwarding to access your application:

   ```bash
   kubectl port-forward <pod-name> 8080:8080
   ```

4. Open a web browser and navigate to [http://localhost:8080](http://localhost:8080) to see the deployed "Hello, Kubernetes!" message or use curl.
   ```sh
   curl http://localhost:8080
   Hello, Kubernetes!%
   ```