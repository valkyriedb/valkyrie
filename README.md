# Valkyrie

![CI](https://github.com/valkyriedb/valkyrie/actions/workflows/ci.yml/badge.svg)
![CD](https://github.com/valkyriedb/valkyrie/actions/workflows/cd.yml/badge.svg)

Dependency agnostic in-memory database with a focus on performance and strict typing.

## Features

- Saving data in RAM for quick access;
- Secure parallel access to the database;
- Minimizing dependence on third-party libraries;
- Communication with the client via a low-level TCP protocol for efficient data exchange;
- Convenient library for interacting with the data warehouse.

## Running Valkyrie with Docker

### 1. Pull the Docker image:

```bash
docker pull valkyriedb/valkyrie
```

### 2. Create `.env` file

Rename `.example.env` to `.env` and change environment variables you want.

### 3. Run the container

Run the container with environment variables and port forwarding:

```bash
docker run --env-file .env -p {PORT}:{PORT} valkyriedb/valkyrie
```

## API

To communicate with the db, connect via TCP connection and follow the structure:

![API](assets/api.svg)
