# Leave Request API

## \ud83d\udccc Table of Contents
- [\ud83d\udcda Introduction](#introduction)
- [\ud83d\udee0\ufe0f Installation](#installation)
- [\u2699\ufe0f Environment Setup](#environment-setup)
- [\ud83d\udcc2 Database & Service Configuration](#database--service-configuration)
- [\ud83d\udcdd API Endpoints](#api-endpoints)
  - [Create Leave Request](#create-leave-request)
  - [Get Leave Balance](#get-leave-balance)
- [\ud83d\ude80 Running the Service](#running-the-service)

---

## \ud83d\udcda Introduction
This is a **simple Leave Request API** that allows employees to request leave, check leave balances, and manage approvals.

---

## \ud83d\udee0\ufe0f Installation
1. Install **Go** (latest version) \u2192 [Download Go](https://go.dev/doc/install)
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

## \u2699\ufe0f Environment Setup
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

## \ud83d\udcc2 Database & Service Configuration
1. Import the database schema:
   ```sh
   psql -U postgres -d leave_db -f schema.sql
   ```
2. Configure the service in `config.yaml`.

---

## \ud83d\udcdd API Endpoints

### \ud83d\udccc **Create Leave Request**
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

### \ud83d\udccc **Get Leave Balance**
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

## \ud83d\ude80 Running the Service
Run the application using:
```sh
go run main.go
```
Then open the API at:
```
http://localhost:8080
```

