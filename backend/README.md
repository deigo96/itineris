# Leave Request API

## 📌 Table of Contents
- [📚 Introduction](#introduction)
- [🛠️ Installation](#installation)
- [⚙️ Environment Setup](#environment-setup)
- [📂 Database & Service Configuration](#database--service-configuration)
- [📝 API Endpoints](#api-endpoints)
  - [Create Leave Request](#create-leave-request)
  - [Get Leave Balance](#get-leave-balance)
- [🚀 Running the Service](#running-the-service)

---

## 📚 Introduction
This is a **simple Leave Request API** that allows employees to request leave, check leave balances, and manage approvals.

---

## 🛠️ Installation
1. Install **Go** (latest version) → [Download Go](https://go.dev/doc/install)
2. Clone this repository:
   ```sh
   git clone https://github.com/deigo96/itineris.git
   ```
3. Navigate to the project directory:
   ```sh
   cd itineris/backend
   ```
4. Install dependencies:
   ```sh
   go mod tidy
   ```

---

## ⚙️ Environment Setup
1. Create a `.env` file:
   ```sh
   touch .env
   ```
2. Add your environment variables:
   ```env
    SERVICE_NAME=""
    SERVICE_HOST=""
    SERVICE_PORT=""
    SECRET_KEY=""


    # database
    DB_HOST=""
    DB_PORT=""
    DB_USER=""
    DB_NAME=""
    DB_PASSWORD=""
   ```

---

## 📂 Database & Service Configuration
1. Import the database schema:
   ```sh
   psql -U postgres -d leave_db -f schema.sql
   ```
2. Configure the service in `config.yaml`.

---

## 📝 API Endpoints

### `url = {{BASE_URL}}/api/v1`

### 📌 Login
- **Endpoint:** `POST {{url}}/auth/login`
- **Description:** Login to get the token.
- **Headers:**
  ```json
    {
        "Content-Type": "application/json"
    }
  ```
- **Request Body:**
  ```json
    {
        "nip": "1234567", //numeric, required
        "password": "12345678" //required
    }
  ```
- **Response:**
  ```json
    {
        "code": "2000",
        "message": "success",
        "data": {
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDA3NTc1NTMsImlkIjoyLCJuaXAiOiIxMjM0NTY3Iiwicm9sZSI6IlN0YWZmIn0.17lfvpQ1p-nCG_LRfCBclI8sNJ5DAMsafB7RwrwVuyM",
            "token_type": "Bearer"
        }
    }
  ```

### 📌 Create Leave Request
- **Endpoint:** `POST /leave`
- **Description:** Submit a leave request.
- **Headers:**
  ```json
  {
    "Authorization": "Bearer <token>",
    "Content-Type": "application/json"
  }
  ```
- **Request Body:**
  ```json
  {
    "start_date": "2024-02-01",
    "end_date": "2024-02-05",
    "reason": "Vacation",
    "leave_type": 1
  }
  ```
- **Response:**
  ```json
  {
    "message": "Leave request submitted successfully"
  }
  ```

### 📌 Get Leave Balance
- **Endpoint:** `GET /leave/balance`
- **Description:** Retrieves the available leave balance for the logged-in user.
- **Headers:**
  ```json
  {
    "Authorization": "Bearer <token>"
  }
  ```
- **Response:**
  ```json
  {
    "leave_balance": 12
  }
  ```

---

## 🚀 Running the Service
Run the application using:
```sh
go run main.go
```
Then open the API at:
```
http://localhost:8080
```

