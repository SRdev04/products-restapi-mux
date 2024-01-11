Product CRUD API using Golang with Gorilla Mux
Overview

This project is a simple CRUD (Create, Read, Update, Delete) API for managing products, implemented in Golang with the Gorilla Mux router. It provides endpoints to interact with a database to perform basic CRUD operations on a "products" table.
Features

    Create Product: Add a new product to the database.
    Read Product: Retrieve product details based on product ID.
    Update Product: Modify existing product information.
    Delete Product: Remove a product from the database.

Dependencies

    Golang: The programming language used for the API.
    Gorilla Mux: A powerful URL router and dispatcher for Golang.
    MySQL: A relational database management system.


Getting Started

    Clone the Repository:

    bash

git clone https://github.com/SRdev04/products-restapi-mux 

Install Dependencies:

bash

# Make sure to have Golang installed
go mod download

Database Setup:

    Set up your database and update the connection details in config/config.go.

Run the Application:

bash

    go run main.go

    API Endpoints:
        Access the API at http://localhost:8080/api/products.

API Endpoints

    GET /api/products: Retrieve a list of all products.
    GET /api/products/{id}: Retrieve details of a specific product by ID.
    POST /api/products: Create a new product.
    PUT /api/products/{id}: Update details of a specific product by ID.
    DELETE /api/products/{id}: Delete a specific product by ID.

Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve this CRUD API.
License

This project is licensed under the MIT License.
