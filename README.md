# URL Shortener

## Introduction

This is a simple URL shortener project, made for learning purposes.

It is similar to Bit.ly or TinyURL, and allows users to input a long URL and receive a shortened version of it. This makes sharing links easier and helps in tracking and managing URLs.

View the live application [here](https://usni.vercel.app/) 👈.

## Features

- **Shorten URLs:** Convert long URLs into manageable short links.
- **Easy to use:** Simple interface for creating and managing URLs.
- **PWA:** Easily install the application on Windows, macOS, Android, and iOS. [Click here to see how](https://www.google.com/search?q=how+to+download+pwa)

## Technologies

- **Backend:** Go, Gin
- **Database:** PostgreSQL
- **Frontend:** Next.js, TypeScript, SCSS.modules

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What things you need to install the software and how to install them:

For the server:

- Go: [Download and install Go](https://golang.org/doc/install)
  - Air: [Install Air](https://github.com/air-verse/air?tab=readme-ov-file#installation) if you want to use live reload
- PostgreSQL: [Download and install PostgreSQL](https://www.postgresql.org/download/)

For the client:

- Node.js: [Download and install Node.js](https://nodejs.org/en/download/)
- npm: [Download and install npm](https://www.npmjs.com/get-npm)

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

- **Note:** If you are using VS Code, you can use the `Go` extension to install the necessary tools and dependencies.

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

#### Client Setup

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
    ```

4. Run the development server:

    - For development:

        ```bash
        npm run dev
        ```

    - For production:

        ```bash
        npm run build && npm start
        ```

## Debugging

If you are using VS Code, you can use the `launch.json` configurations to debug the server and client code.

## Deployment

I am deploying the server to render.com and the client to Vercel. You can deploy the server and client to any platform of your choice.
