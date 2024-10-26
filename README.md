# Go-JWT

## Overview

Go-JWT is a project that demonstrates user authentication using JSON Web Tokens (JWT) in a Go application with Gin framework.

## Project Structure

The project consists of the following main components:

-   `main.go`: Contains the main entry point of the application where routes are defined and server initialization takes place.
-   `controllers/userController.go`: Includes functions for user signup and login processes, interacting with the database and handling JWT token generation.
-   `middleware/authorize.go`: Middleware for authorizing user requests based on JWT tokens.
-   `middleware/noAuthorize.go`: Middleware for handling routes that do not require authorization.

## Setup

1. Ensure Go is installed on your machine.
2. Clone the repository.
3. Install dependencies by running `go mod tidy`.
4. Set up the required environment variables.
5. Run the application using `go run main.go`.

## Usage

-   **Signup**: Endpoint `/signup` allows users to create a new account by providing an email and password.

```bash
curl --location 'http://127.0.0.1:3030/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "a@gmail.com",
    "password": "password"
}'
```

-   **Signin**: Endpoint `/signin` enables users to log in with their credentials and receive a JWT token for authentication.

```bash
curl --location 'http://127.0.0.1:3030/signin' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "a@gmail.com",
    "password": "password"
}'
```

-   **Home**: Endpoint `/home` users can see if the authentication is susccess.

```bash
curl --location 'http://127.0.0.1:3030/home'
```

-   **Home**: Endpoint `/signout` delete the cookie and singout the user.

```bash
curl --location 'http://127.0.0.1:3030/signout'
```

-   **Authorization**: Middleware ensures that certain routes are only accessible with a valid JWT token.

## Dependencies

-   `github.com/gin-gonic/gin`: Web framework used for building the API.
-   `github.com/golang-jwt/jwt/v5`: Library for JWT token generation and validation.
-   `golang.org/x/crypto/bcrypt`: Package for hashing passwords securely.

## Contributors

-   [Thar Htoo](https://github.com/etharrra)

Feel free to contribute by forking the repository and submitting a pull request!
