# User Management Go App

This is a basic Go application for user management with CRUD operations (Create, Read, Update, Delete). It follows a four-layer architecture with the following layers:

- **Handler**: Handles incoming HTTP requests.
- **Controller**: Coordinates requests between the handler and the service.
- **Service**: Contains the business logic.
- **Repo**: Handles interactions with the MySQL database.

The app uses Protobuf messages for defining the user structure and MySQL as the database, running in a Docker Compose environment.

## Features
- Create a new user
- Get user details
- Update user information
- Delete a user

## Setup

   ```bash
   git clone https://github.com/adosalkanovicc/crud.git
   cd crud
   go run main.go
   ```

