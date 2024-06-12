# URL Shortener

## Introduction

This project is a simple URL shortener service, similar to Bit.ly or TinyURL. It allows users to input a long URL and receive a shortened version of it. This makes sharing links easier and helps in tracking and managing URLs.

## Features

- **Shorten URLs:** Convert long URLs into manageable short links.
- **Easy to use:** Simple interface for creating and managing URLs.

## Technologies

- **Backend:** Go, Gin
- **Database:** PostgreSQL (with GORM)
- **Frontend:** Next.js, SCSS.modules

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What things you need to install the software and how to install them:

- Go: [Download and install Go](https://golang.org/doc/install)
- PostgreSQL: [Download and install PostgreSQL](https://www.postgresql.org/download/)

### Installing

A step by step series of examples that tell you how to get a development env running:

1. Clone the repository:

    ```bash
    git clone github.com/nitzanpap/url-shortener
    ```

2. Change into the project directory:

    ```bash
    cd url-shortener
    ```

#### Server Setup

1. Change into the server directory:

    ```bash
    cd server
    ```

2. Create a `.env` file in the server directory:

    ```bash
    touch .env
    ```

    Make sure to populate the `.env` file according to the `.env.example` file.

3. Run the server via make:

    ```bash
    make run/live
    ```

    Or via Docker:

    ```bash
    docker compose up
    ```

### Client Setup

- TBD
