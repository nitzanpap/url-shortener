# URL Shortener

## Introduction

This project is a simple URL shortener service, similar to Bit.ly or TinyURL. It allows users to input a long URL and receive a shortened version of it. This makes sharing links easier and helps in tracking and managing URLs.

## Features

- **Shorten URLs:** Convert long URLs into manageable short links.
- **Easy to use:** Simple interface for creating and managing URLs.

## Technologies

- **Backend:** Go, Gin
- **Database:** PostgreSQL
- **Frontend:** Next.js, TypeScript, Tailwind CSS

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

    - For development, run with live reload:

        ```bash
        make run/live
        ```

    - Or build and run the server:

        1. Manually:

            ```bash
            make build && make run
            ```

        2. Via Docker:

            ```bash
            docker compose up
            ```

### Client Setup

1. Change into the client directory:

    ```bash
    cd client
    ```

2. Create a `.env.local` file in the client directory:

    ```bash
    touch .env.local
    ```

    Make sure to populate the `.env.local` file according to the `.env.example` file.

3. Install the dependencies:

    ```bash
    npm install
    # or
    yarn
    # or
    pnpm install
    # or
    bun install
    ```

4. Run the development server:

    - For development:

        ```bash
        npm run dev
        # or
        yarn dev
        # or
        pnpm dev
        # or
        bun dev
        ```

    - For production:

        ```bash
        npm run build && npm start
        # or
        yarn build && yarn start
        # or
        pnpm build && pnpm start
        # or
        bun build && bun start
        ```
