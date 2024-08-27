# Enterpris Inventory and Supply Chain Management System

The Inventory Supply Chain System is a robust and scalable application designed to manage the entire lifecycle of products, orders, shipments, vendors, and inventory items within a supply chain. This system allows users to track products from suppliers to warehouses, monitor inventory levels, manage shipments, and analyze supply chain performance.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Installation](#installation)
- [Environment Variables](#environment-variables)
- [API Documentation](#api-documentation)
  - [Authentication](#authentication)
  - [Inventory](#inventory)
  - [Orders](#orders)
  - [Shipments](#shipments)
  - [Vendors](#vendors)
  - [Items](#items)
  - [Suppliers](#suppliers)
- [Usage](#usage)
- [License](#license)
- [Contributing](#contributing)
- [Support](#support)

## Overview

The Inventory Supply Chain System is built to streamline and automate the supply chain management process. It handles a variety of business-critical operations, such as inventory tracking, vendor management, shipment tracking, and order management. This system is suitable for small to large-scale businesses that need an integrated solution for supply chain automation.

## Features

- **Inventory Management**: Create, read, update, and delete inventory items.
- **Order Management**: Manage customer orders with full CRUD support.
- **Shipment Tracking**: Track and manage shipments, including status updates and location tracking.
- **Vendor Management**: Manage supplier and vendor details.
- **Authentication**: Secure API access with JWT-based authentication.
- **Supplier Management**: Manage supplier details and relationships with the supply chain.
- **Analytics**: Gain insights into supply chain performance, orders, and shipments.

## Tech Stack

- **Backend**: Go (Golang), Gorilla Mux, Gorm
- **Database**: PostgreSQL
- **API Security**: JWT Authentication
- **Deployment**: Docker, Kubernetes

## Installation

### Prerequisites

- Go v1.16+
- PostgreSQL v12+
- Docker (optional, for containerized deployment)

### Clone the Repository


git clone https://github.com/yourusername/inventory-supply-chain-system.git
cd inventory-supply-chain-system


### Install Dependencies


go mod tidy
Setup PostgreSQL Database
Ensure you have PostgreSQL running and create the database with the following credentials:


CREATE DATABASE inventory_db;
Running the Application
1.	Setup Environment Variables: Create a .env file in the project root:
env

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=inventory_db
JWT_SECRET=yourjwtsecret
2.	Run the Application:


go run cmd/main.go

The server will be running on http://localhost:8080.

Environment Variables

The following environment variables are required:
•	DB_HOST: The database host (e.g., localhost).
•	DB_PORT: The port on which PostgreSQL is running (default 5432).
•	DB_USER: The database user (e.g., postgres).
•	DB_PASSWORD: The password for the PostgreSQL user.
•	DB_NAME: The name of the PostgreSQL database.
•	JWT_SECRET: The secret used for signing JWT tokens.

```bash
### API Documentation
Authentication
•	POST /api/auth/register: Register a new user.
•	POST /api/auth/login: Authenticate a user and retrieve a JWT token.
Inventory
•	POST /api/inventory: Create a new inventory item.
•	GET /api/inventory/{id}: Retrieve details of an inventory item by ID.
•	PUT /api/inventory/{id}: Update an existing inventory item.
•	DELETE /api/inventory/{id}: Delete an inventory item by ID.
Orders
•	POST /api/orders: Create a new order.
•	GET /api/orders/{id}: Retrieve details of an order by ID.
•	PUT /api/orders/{id}: Update an existing order.
•	DELETE /api/orders/{id}: Delete an order by ID.
Shipments
•	POST /api/shipments: Create a new shipment.
•	GET /api/shipments/{id}: Retrieve details of a shipment by ID.
•	PUT /api/shipments/{id}: Update an existing shipment.
•	DELETE /api/shipments/{id}: Delete a shipment by ID.
Vendors
•	POST /api/vendors: Create a new vendor.
•	GET /api/vendors/{id}: Retrieve details of a vendor by ID.
•	PUT /api/vendors/{id}: Update an existing vendor.
•	DELETE /api/vendors/{id}: Delete a vendor by ID.
Items
•	POST /api/items: Create a new item.
•	GET /api/items/{id}: Retrieve details of an item by ID.
•	PUT /api/items/{id}: Update an existing item.
•	DELETE /api/items/{id}: Delete an item by ID.
Suppliers
•	POST /api/suppliers: Create a new supplier.
•	GET /api/suppliers/{id}: Retrieve details of a supplier by ID.
•	PUT /api/suppliers/{id}: Update an existing supplier.
•	DELETE /api/suppliers/{id}: Delete a supplier by ID.

### Usage
Register a New User

curl -X POST http://localhost:8080/api/auth/register -d '{
  "email": "user@example.com",
  "password": "yourpassword",
  "name": "John Doe"
}'
Authenticate and Get Token


curl -X POST http://localhost:8080/api/auth/login -d '{
  "email": "user@example.com",
  "password": "yourpassword"
}'

You will receive a JWT token to use in the Authorization header for protected routes.

### License
This project is licensed under the MIT License - see the LICENSE file for details.

### Contributing
Contributions are welcome! Please read the CONTRIBUTING.md file for more details on how to get involved.
1.	Fork the project
2.	Create a feature branch (git checkout -b feature/YourFeature)
3.	Commit your changes (git commit -m 'Add some feature')
4.	Push to the branch (git push origin feature/YourFeature)
5.	Open a Pull Request

### Support
For support, open an issue in the GitHub repository, or contact the project maintainers via email at shadrach.abdul@gmail.com






