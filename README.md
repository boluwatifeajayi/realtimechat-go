# Chat App API Documentation

## Introduction

This document provides detailed information about the Chat App API, including endpoints, request/response formats, and examples. The API is built using Go and Gin web framework and can be run locally for development and testing purposes.

## Base URL

```
http://localhost:8080/api
```

## Endpoints

### 1. Register User

- **Endpoint:** `/register`
- **Method:** POST
- **Description:** Register a new user with the provided details.

#### Request Body:

```json
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "password": "password123"
}
```

#### Request Headers:

```
Content-Type: application/json
```

#### Response:

- Success (200 OK):
  ```json
  {
    "message": "User registered successfully"
  }
  ```

- Error (400 Bad Request):
  ```json
  {
    "error": "Name, email, and password are required"
  }
  ```

- Error (409 Conflict):
  ```json
  {
    "error": "Email already exists"
  }
  ```

- Error (500 Internal Server Error):
  ```json
  {
    "error": "Internal server error"
  }
  ```

### 2. Login User

- **Endpoint:** `/login`
- **Method:** POST
- **Description:** Login a user with the provided email and password.

#### Request Body:

```json
{
  "email": "john.doe@example.com",
  "password": "password123"
}
```

#### Request Headers:

```
Content-Type: application/json
```

#### Response:

- Success (200 OK):
  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
  ```

- Error (400 Bad Request):
  ```json
  {
    "error": "Invalid email or password"
  }
  ```

- Error (401 Unauthorized):
  ```json
  {
    "error": "Invalid email or password"
  }
  ```

- Error (500 Internal Server Error):
  ```json
  {
    "error": "Internal server error"
  }
  ```

### 3. Get All Users

- **Endpoint:** `/users`
- **Method:** GET
- **Description:** Get a list of all users.

#### Response:

- Success (200 OK):
  ```json
  [
    {
      "id": "60d0fe4f5311236168a109ca",
      "name": "John Doe",
      "email": "john.doe@example.com",
      "profile_pic": "https://www.shutterstock.com/image-vector/user-profile-icon-vector-avatar-600nw-2247726673.jpg",
      "join_date": "2023-01-01T00:00:00Z",
      "last_active": "2023-01-01T00:00:00Z",
      "is_active": true
    },
    {
      "id": "60d0fe4f5311236168a109cb",
      "name": "Jane Smith",
      "email": "jane.smith@example.com",
      "profile_pic": "https://www.shutterstock.com/image-vector/user-profile-icon-vector-avatar-600nw-2247726673.jpg",
      "join_date": "2023-01-02T00:00:00Z",
      "last_active": "2023-01-02T00:00:00Z",
      "is_active": true
    }
  ]
  ```

- Error (500 Internal Server Error):
  ```json
  {
    "error": "Internal server error"
  }
  ```

### 4. Search Users

- **Endpoint:** `/users/search`
- **Method:** GET
- **Description:** Search users by name.

#### Query Parameters:

- `query` (string, required): Search query.

#### Response:

- Success (200 OK):
  ```json
  [
    {
      "id": "60d0fe4f5311236168a109ca",
      "name": "John Doe",
      "email": "john.doe@example.com",
      "profile_pic": "https://www.shutterstock.com/image-vector/user-profile-icon-vector-avatar-600nw-2247726673.jpg",
      "join_date": "2023-01-01T00:00:00Z",
      "last_active": "2023-01-01T00:00:00Z",
      "is_active": true
    }
  ]
  ```

- Error (500 Internal Server Error):
  ```json
  {
    "error": "Internal server error"
  }
  ```

### 5. Get Chat List

- **Endpoint:** `/users/{user_id}/chat-list`
- **Method:** GET
- **Description:** Get a list of users that a particular user has existing chats with.

#### Path Parameters:

- `user_id` (string, required): User ID.

#### Response:

- Success (200 OK):
  ```json
  [
    "60d0fe4f5311236168a109cb",
    "60d0fe4f5311236168a109cc"
  ]
  ```

- Error (500 Internal Server Error):
  ```json
  {
    "error": "Internal server error"
  }
  ```

### 6. Get User By ID

- **Endpoint:** `/users/{user_id}`
- **Method:** GET
- **Description:** Get user details by ID.

#### Path Parameters:

- `user_id` (string, required): User ID.

#### Response:

- Success (200 OK):
  ```json
  {
    "id": "60d0fe4f5311236168a109ca",
    "name": "John Doe",
    "email": "john.doe@example.com",
    "profile_pic": "https://www.shutterstock.com/image-vector/user-profile-icon-vector-avatar-600nw-2247726673.jpg",
    "join_date": "2023-01-01T00:00:00Z",
    "last_active": "2023-01-01T00:00:00Z",
    "is_active": true
  }
  ```

- Error (404 Not Found):
  ```json
  {
    "error": "User not found"
  }
  ```

- Error (500 Internal Server Error):
  ```json
  {
    "error": "Internal server error"
  }
  ```

### 7. Get User Profile

- **Endpoint:** `/users/{user_id}/profile`
- **Method:** GET
- **Description:** Get user profile details by ID.

#### Path Parameters:

- `user_id` (string, required): User ID.

#### Response:

- Success (200 OK):
  ```json
  {
    "username": "John Doe",
    "email": "john.doe@example.com",
    "join_date": "2023-01-01T00:00:00Z",
    "last_active": "2023-01-01T00:00:00Z"
  }
  ```

- Error (404 Not Found):
  ```json
  {
    "error": "User not found"
  }
  ```

- Error (500 Internal Server Error):
  ```json
  {
    "error": "Internal server error"
  }
  ```

### 8. Send Message

- **Endpoint:** `/send-message`
- **Method:** POST
- **Description:** Send a message between two users.

#### Request Body:

```json
{
  "sender_id": "60d0fe4f5311236168a109ca",
  "receiver_id": "60d0fe4f5311236168a109cb",
  "content": "Hello, how are you?"
}
```

#### Request Headers:

```
Content-Type: application/json
```

#### Response:

- Success (200 OK):
  ```json
  {
    "message": "Message sent successfully"
  }
  ```

- Error (400 Bad Request):
  ```json
  {
    "error": "Invalid request body"
  }
  ```

- Error (500 Internal Server Error):
  ```json
  {
    "error": "Internal server error"
  }
  ```

### 9. Get Messages

- **Endpoint:** `/messages/{sender_id}/{receiver_id}`
- **Method:** GET
- **Description:** Get messages between two users.

#### Path Parameters:

- `sender_id` (string, required): Sender's User ID.
- `receiver_id` (string, required): Receiver's User ID.

#### Response:

- Success (200 OK):
  ```json
  [
    {
      "id": "60d0fe4f5311236168a109cd",
      "sender_id": "60d0fe4f5311236168a109ca",
      "receiver_id": "60d0fe4f5311236168a109cb",
      "content": "Hello, how are you?",
      "timestamp": "2023-01-01T00:00:00Z"
    },
    {
      "id": "60d0fe4f5311236168a109ce",
      "sender_id": "60d0fe4f5311236168a109cb",
      "receiver_id": "60d0fe4f5311236168a109ca",
      "content": "I'm good, thanks!",
      "timestamp": "2023-01-01T00:01:00Z"
    }
  ]
  ```

- Error (500 Internal Server Error):
  ```json
  {
    "error": "Internal server error"
  }
  ```
