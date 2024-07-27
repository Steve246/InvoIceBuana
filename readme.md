# Invoice Management System

## Overview
The Invoice Management System is a RESTful API built with Go and Gin for managing invoices. It supports operations such as creating, updating, retrieving, and deleting invoices, as well as managing customers and items.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Invoice Endpoints](#invoice-endpoints)
  - [Create Invoice](#create-invoice)
  - [Update Invoice](#update-invoice)
  - [Get Invoice by ID](#get-invoice-by-id)
  - [Get All Invoices](#get-all-invoices)
- [Customer Endpoints](#customer-endpoints)
  - [Create Customer](#create-customer)
  - [Get All Customers](#get-all-customers)
- [Item Endpoints](#item-endpoints)
  - [Create Item](#create-item)
  - [Get All Items](#get-all-items)

## Features
- Create, Get All, Get by ID, and Update invoices.
- Create, and Get all customers and items data.
- Calculate totals, tax, and grand total automatically.
- Transaction management to ensure data integrity.

## Prerequisites
- Go 1.15 or higher
- MySQL 
- Gin Gonic framework
- Gorm framework

## Installation

1. Go to .env file, then make sure to create a new Database named "steven_invoice" or use an existing empty database of your choice. Please suit the DB credentials, based on information used in the current device
    ```
    DB_HOST="localhost"
    DB_PORT="3306"
    DB_NAME="steven_invoice"
    DB_USER="root"
    DB_PASS="12345678"
    ```
2. Set the URL name, just use the URL below if this is not changed 
    ```
    API_URL="localhost:8080"
    ```
3. Let's migrate up to create tables inside the database by changing this parameter inside the .env file
    ```
    ENV="migration"
    MIGRATION_DIRECTION="up"
    ```
4. After migration finish, marked by "proses migrasi berhasil dilakukan". Please change the .env parameter as below to see the gin debug text
    ```
    ENV="dev"
    ```
5. In case you want to delete all tables inside the database. Please change the parameter inside the .env file as below
    ```
    ENV="migration"
    MIGRATION_DIRECTION="down"
    ```

6. To run the program, make sure 2 things are initialized
    ```
    ENV="dev" --> in .env file
    go run . --> in terminal
    ```

## Invoice Endpoints

### Create Invoice

- **URL:** `/invoice/add`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "subject" : "Invoice Pembelian 2023",
    "customer_id": "7f12c3e7-948b-4ad6-b5fc-41cfce0c30a7",
    "issue_date": "2024-07-23",
    "due_date": "2024-08-23",
    "items": [
        {
            "item_id": "61a8067d-2a38-48e4-b82f-797ad91e1a17",
            "quantity": 5
            
        },
         {
            "item_id": "21e5f961-272d-4333-a9a0-d880e9dd5924",
            "quantity": 2
            
        }
       
    ]
  }

- **Response Body:**
  ```json
  {
    "message": "Invoice Created Sucessfully",
    "data": {
        "id": 32,
        "invoice_id": "INV033",
        "subject": "Invoice Pembelian 2023",
        "customer_id": "7f12c3e7-948b-4ad6-b5fc-41cfce0c30a7",
        "issue_date": "2024-07-23",
        "due_date": "2024-08-23",
        "status": "unpaid",
        "customer": {
            "id": 1,
            "name": "Ambar",
            "address": "Jalan Sawa Sawit"
        },
        "items": [
            {
                "id": 66,
                "item_name": "Bayam",
                "quantity": 5,
                "unit_price": 120.23,
                "total_price": 601.15
            },
            {
                "id": 67,
                "item_name": "Anggur",
                "quantity": 2,
                "unit_price": 390.23,
                "total_price": 780.46
            }
        ],
        "totals": {
            "total_items": 2,
            "subtotal": 1381.61,
            "tax": 138.16,
            "grand_total": 1519.77
        },
        "created_at": "2024-07-27T15:46:50+07:00",
        "updated_at": "2024-07-27T15:46:50+07:00"
    }
  }



### Update Invoice

- **URL:** `/invoice/update/invoice/{INVOICE_ID}`
  - EXAMPLE --> `/invoice/update/invoice/INV018`
- **Method:** `PUT`
- **Request Body:**
  ```json
  {
    "subject": "Invoice Pembelian Terbaru",
    "issue_date": "2025-07-23",
    "due_date": "2025-08-23",
    "customer_id": "473e3980-53aa-4322-bb97-9ae214806fa4",
    "items": [
        {
        "item_id": "21e5f961-272d-4333-a9a0-d880e9dd5924",
        "quantity": 1
        },
        {
        "item_id": "dc042d89-c932-48ef-af6d-760b72e57828",
        "quantity": 23
        }
    ]
  }

- **Response Body:**
  ```json
  {
    "message": "Invoice Updated Successfully"
  }


### Get Invoice By Id

- **URL:** `/invoice/display/customInvoice/{INVOICE_ID}`
  - EXAMPLE --> `/invoice/display/customInvoice/INV018`
- **Method:** `GET`
- **Request Body:**
  ```json
  No Request Body

- **Response Body:**
  ```json
  {
    "message": "Invoice Data Succesfully Retrieved",
    "data": {
        "id": 17,
        "invoice_id": "INV018",
        "subject": "Updated Invoice Subject",
        "customer_id": "473e3980-53aa-4322-bb97-9ae214806fa4",
        "issue_date": "2025-07-23",
        "due_date": "2025-08-23",
        "status": "unpaid",
        "customer": {
            "id": 2,
            "name": "Ulama",
            "address": "Jalan Ulalal Sawit"
        },
        "items": [
            {
                "id": 65,
                "item_name": "Naga",
                "quantity": 23,
                "unit_price": 390.23,
                "total_price": 8975.29
            }
        ],
        "totals": {
            "total_items": 1,
            "subtotal": 1020.92,
            "tax": 102.09,
            "grand_total": 1123.01
        },
        "created_at": "2024-07-26T14:21:06+07:00",
        "updated_at": "2024-07-27T14:48:35+07:00"
    }
  } 


### Get All Invoice

- **URL:** `/invoice/display/invoice?limit={number}&offset={number}`
  - EXAMPLE --> `/invoice/display/invoice?limit=5&offset=0`
- **Method:** `GET`
- **Request Body:**
  ```json
  Query Param
    - limit (Example: 5)
    - offset (Example: 0)

- **Response Body:**
  ```json
  {
    "message": "Invoice Data Succesfully Retrieved",
    "data": [
        {
            "invoice_id": "INV018",
            "issue_date": "2025-07-23T00:00:00+07:00",
            "subject_invoice": "Updated Invoice Subject",
            "total_item": 1,
            "customer_name": "Ulama",
            "due_date": "2025-08-23T00:00:00+07:00",
            "status": "unpaid"
        },
        {
            "invoice_id": "INV019",
            "issue_date": "2024-07-23T00:00:00+07:00",
            "subject_invoice": "Spring Marketing",
            "total_item": 2,
            "customer_name": "Ambar",
            "due_date": "2024-08-23T00:00:00+07:00",
            "status": "unpaid"
        },
        {
            "invoice_id": "INV020",
            "issue_date": "2024-07-23T00:00:00+07:00",
            "subject_invoice": "Spring Marketing",
            "total_item": 2,
            "customer_name": "Ambar",
            "due_date": "2024-08-23T00:00:00+07:00",
            "status": "unpaid"
        },
        {
            "invoice_id": "INV021",
            "issue_date": "2024-07-23T00:00:00+07:00",
            "subject_invoice": "Spring Marketing",
            "total_item": 2,
            "customer_name": "Ambar",
            "due_date": "2024-08-23T00:00:00+07:00",
            "status": "unpaid"
        },
        {
            "invoice_id": "INV022",
            "issue_date": "2024-07-23T00:00:00+07:00",
            "subject_invoice": "Spring Marketing",
            "total_item": 2,
            "customer_name": "Ambar",
            "due_date": "2024-08-23T00:00:00+07:00",
            "status": "unpaid"
        }
    ]
  }


## Customer Endpoints

### Create Customer

- **URL:** `/invoice/add/customer`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "custName": "Ambar",
    "custAddress": "Jalan Rajawali 2 Timur"
  }

- **Response Body:**
  ```json
  {
    "message": "Customer Created Succesfully",
    "data": ""
  }

### Get All Customers

- **URL:** `/invoice/display/customer?limit={number}&offset={number}`
  - EXAMPLE --> `/invoice/display/customer?limit=3&offset=0`
- **Method:** `GET`
- **Request Body:**
  ```json
  Query Param
    - limit (Example: 3)
    - offset (Example: 0)

- **Response Body:**
  ```json
  {
    "message": "Customer Data Succesfully Retrieved",
    "data": [
        {
            "CustomerID": "7f12c3e7-948b-4ad6-b5fc-41cfce0c30a7",
            "CustomerName": "Ambar",
            "CustomerAddress": "Jalan Sawa Sawit"
        },
        {
            "CustomerID": "473e3980-53aa-4322-bb97-9ae214806fa4",
            "CustomerName": "Ulama",
            "CustomerAddress": "Jalan Ulalal Sawit"
        }
    ]
  }


## Item Endpoints

### Create Item

- **URL:** `/invoice/add/items`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "itemName": "Naga",
    "itemType": "Buah",
    "itemPrice": 390.23
  }

- **Response Body:**
  ```json
  {
    "message": "Item Created Succesfully",
    "data": ""
  }

### Get All Item

- **URL:** `/invoice/display/items?limit={number}&offset={number}`
  - EXAMPLE --> `/invoice/display/items?limit=3&offset=0`
- **Method:** `GET`
- **Request Body:**
  ```json
  Query Param
    - limit (Example: 3)
    - offset (Example: 0)

- **Response Body:**
  ```json
  {
    "message": "Item Data Succesfully Retrieved",
    "data": [
        {
            "itemId": "61a8067d-2a38-48e4-b82f-797ad91e1a17",
            "itemName": "Bayam",
            "itemType": "sayuran",
            "itemPrice": 120.23,
            "itemPriceFormatted": "120.23"
        },
        {
            "itemId": "21e5f961-272d-4333-a9a0-d880e9dd5924",
            "itemName": "Anggur",
            "itemType": "sayuran",
            "itemPrice": 390.23,
            "itemPriceFormatted": "390.23"
        },
        {
            "itemId": "dc042d89-c932-48ef-af6d-760b72e57828",
            "itemName": "Naga",
            "itemType": "Buah",
            "itemPrice": 390.23,
            "itemPriceFormatted": "390.23"
        }
    ]
  }