# go-driver-register

## Description

Backend application to manage drivers and vehicles, focusing on fuel efficiency.

## Technologies Used

- Go
- PostgreSQL
- Docker
- Echo Framework

## How to Run

### Prerequisites

- Docker
- Docker Compose
- Python 3
- Pip 3
- Go

### Steps

1. Clone the repository:
    ```sh
    git clone https://github.com/paluan-batista/go-driver-register.git
    ```

2. Navigate to the project directory:
    ```sh
    cd go-driver-register
    ```

3. Build and start the containers:
    ```sh
    make start-env
    ```

4. Access the application at [http://localhost:8000](http://localhost:8000).

## Endpoints

### Drivers

- `POST /drivers` - Create a new driver
- `GET /drivers` - List all drivers
- `GET /drivers/:id` - Get a driver by ID
- `PUT /drivers/:id` - Update a driver
- `DELETE /drivers/:id` - Delete a driver

### Vehicles

- `POST /vehicles` - Create a new vehicle
- `GET /vehicles` - List all vehicles
- `GET /vehicles/:id` - Get a vehicle by ID
- `PUT /vehicles/:id` - Update a vehicle
- `DELETE /vehicles/:id` - Delete a vehicle
- `POST /vehicles/:vehicle_id/assign/:driver_id` - Assign a driver to a vehicle
