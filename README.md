# Golang Test Project: User Registration System

This project aims to test your proficiency in Go, your understanding of HTTP protocols, and application structuring effectiveness. We'll also assess your ability to handle concurrency, implement basic authentication, unit testing, write API documentation, and dockerize applications.

## 👨‍💻 Project Description

Your assignment is to build an HTTP-based User Registration System in Go. This system will primarily offer three services:

1. Register a new user.
2. Fetch registered user details.
3. Delete a registered user.

After registering a new user, the system should simulate sending a registration confirmation email in the background. 

## 🚀 Core Features

Your application should cater to these core features:

1. An endpoint (`POST`) to register a new user with data: name, email, and password.
2. An endpoint (`GET`) to fetch a user's details using an authorization token.
3. An endpoint (`DELETE`) to delete a user using an authorization token.
4. An in-memory map to store user data.
5. Appropriate validation of incoming user data.
6. A simulation of sending a registration confirmation email in a non-blocking manner.

## 🔧 Technical Requirements

Your solution should match the following technical specifications:

- Employ Go's native `net/http` and `slog` packages for server creation and logging, respectively.
- You can leverage existing tools such as go-chi, gorilla/mux, or others to streamline and enhance the routing functionality.
- Organize your code according to best practices (DDD is suggested but not compulsory).
- Your code should be clearly expressive and easy to maintain.
- Efficiently handle errors and return informative error messages to the client.
- Implement a form of token-based authentication.
- Design containers using Docker (this is a nice-to-have but not a requirement).
- Include unit tests for your code.
- Implement basic rate limiting on your endpoints.

## 📝 Post Development

After you complete your code, do the following:

- Push your final code, along with test files and Dockerfile, if applicable, to your GitHub repository.
- Make regular commits with clear, descriptive messages.

## 📋 Evaluation

Our evaluation of your work will be based on:

- Functional completion
- Proper application of coding best practices
- Code clarity and organization
- Effective data validation and error handling
- Quality and clarity of GitHub commit messages
- Unit testing
- Your application's ability to handle concurrency efficiently
- Dockerization of your application (nice to have)
- Clear and concise API documentation (nice to have)

## ⏰ Time Frame

The challenge should be completed within one week.

We can't wait to see your efficient and well-structured User Registration System in Go. Happy coding!
