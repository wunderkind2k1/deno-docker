# deno-docker Tutorial

## 1. What is This Repo About?

This repository demonstrates how to use modern tools like **Deno**, **Docker**, and **Dagger** to create a simple yet powerful setup for building, testing, and deploying software. Let's break down what's inside this repository and what it's all about:

- **Technologies Used**:

  - **Deno**: The JavaScript/TypeScript runtime that provides a secure, efficient, and modern environment for building apps.
  - **Docker**: The platform for creating, deploying, and running containerized applications.
  - **Dagger**: A new CI/CD tool that lets you define pipelines as code, using your programming language of choice.

- **Prerequisites**:

  - A basic understanding of **Go** (for working with the Dagger pipeline code).
  - Familiarity with **Deno** and **Docker**.
  - **Go**, **Deno**, and **Docker** installed on your local machine.

- **Interplay of Technologies**:

  - This repo uses **Deno** to build an application that is then containerized with **Docker**. The build and containerization processes are orchestrated using **Dagger** through the `Daggerpipeline.go` script.

- **Outcome**: The final product is a Docker image containing a Deno application, ready to be deployed. The Dagger pipeline makes sure everything is streamlined and easy to replicate.

## 2. How to Use This Repo

### Step 1: Running the Dagger Pipeline (`Daggerpipeline.go`)

The Dagger pipeline (`Daggerpipeline.go`) is the brains of the operation. It defines how the application is built, tested, and packaged into a Docker container.

- To run the Dagger pipeline, navigate to the root directory of the repository and use the following command:

  ```shell
  go run Daggerpipeline.go
  ```

  This will trigger the pipeline, which compiles the Deno application and builds a Docker image containing the application.

  Of course, you can also compile or even cross-compile the `Daggerpipeline.go` file to an executable, allowing you to distribute the build pipeline like any other Go program.

- **How It Works**: The Dagger pipeline is written in **Go**, making it more flexible compared to traditional CI/CD files like `github.yml` or `gitlab.yml`. Since it's executable code, you can run it on your local machine as well as in a CI/CD environment, which means you get the same results no matter where it's run. Look mom: No build server needed!

### Step 2: Running the Generated Binaries and Docker Image

After running the pipeline, you'll end up with two key artifacts:

1. **The Compiled Deno Binary**: This binary is generated by Deno's cross-compilation capabilities. You can run it directly using:

   ```shell
   ./hello-mac //on osx
   ./hello-linux //on linux
   ./hello-win.exe //on windows
   ```

2. **The Docker Image**: The Docker image contains the Deno app, which you can run using the `run-container.sh` script:

   ```shell
   ./run-container.sh
   ```

   A more or less funny excuse should appear on your terminal.

### Step 3: Testing the Pipeline with `go test`

Testing is crucial for any good pipeline, and Dagger makes it easy to add tests.

- Run the following command to test the Dagger pipeline:
  ```sh
  go test -v ./...
  ```
  This command runs all tests in the repository.
- **What the Test Verifies**:
  - The `Daggerpipeline_test.go` file is designed to test various parts of the pipeline, ensuring that each step runs smoothly and produces the correct output.
  - Specifically, it verifies that:
    - **Deno Binary Compilation**: It ensures that the Deno application compiles successfully into three binaries (`hello-mac`, `hello-linux`, `hello-win.exe`) for different platforms.
    - **Docker Image Creation**: The test checks if the Docker image is built correctly from the generated binary.
    - **Output Verification**: It validates that the expected output, including the Docker image and cross-compiled binaries, matches what is defined in the pipeline.
    - **Cross-Platform Compatibility**: The test also ensures that the compiled binaries can be executed without errors on their respective platforms, simulating a real-world deployment scenario.
    - **Error Handling**: Any errors during the compilation, packaging, or containerization process are caught and reported, ensuring the robustness of the pipeline.
  - Writing tests for your pipeline helps catch issues early, ensuring a smooth CI/CD process and maintaining the reliability of your builds.

## 3. Conclusion: Why These Technologies?

### Why Deno?

- **Compilation & Cross-Compilation**: Unlike Node.js, Deno has first-class support for **compiling JavaScript/TypeScript** into executables, which makes it easier to distribute your application.
- **Security by Default**: Deno runs with strict permissions, requiring you to explicitly use flags like `--allow-net` to grant network access. This reduces the attack surface by default, improving the security posture of your app.
- **No \*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*****`node_modules`**: With **Deno**, you can say goodbye to `node_modules` and `package.json`. Deno's module system fetches dependencies directly via URLs, which means less clutter and less setup hassle.

### Why Dagger?

- **Flexibility of Go**: Writing your pipeline in **Go** provides a lot of power. Unlike traditional YAML-based CI/CD configurations, the pipeline is **compilable code**, which means you can do anything that you would do in a regular Go program.

Dagger also supports other languages besides Go, such as Python, Typescript, and others, giving you flexibility to choose the language that best fits your team's skill set or the needs of your project.

- **Run Locally or in CI/CD**: With Dagger, you can run the pipeline both **locally** and in your CI/CD environment. This "write once, run anywhere" approach ensures consistency and reduces surprises between development and production.
- **Readable and Understandable**: Writing the pipeline in Go makes it easier to understand for developers who are already familiar with programming. You don't have to learn a new domain specific YAML syntax. Instead, everything is in Go, which means **no more head-scratching YAML indents**!

### Writing Tests for the Pipeline

Testing pipelines is often overlooked, but it provides real value:

- **Catch Issues Early**: By writing Go tests for the Dagger pipeline, you can ensure that any future changes to the pipeline don't break existing functionality. This can also be used to enforce company-wide compliance rules, ensuring that your CI/CD processes adhere to organizational standards.
- **Less Debugging Hassle**: Rather than waiting for a CI job to fail and then struggling to reproduce the issue, you can **test locally**, debug more easily, and save time. <3

## Want to Try It Out?

Feel free to clone this repository, run the Dagger pipeline, and get hands-on with Deno, Docker, and Dagger. If you run into any issues or have questions, don't hesitate to open an issue or a pull request. Happy coding!
