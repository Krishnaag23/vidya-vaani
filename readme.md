# Vidya Vaani

Vidya Vaani is a scalable, real-time audio captioning system designed to transcribe audio in real-time, provide accurate captions, and offer multilingual support, ensuring data privacy and compliance with industry standards.

## Table of Contents

- [Requirements](#requirements)
- [Getting Started](#getting-started)
- [Cloning the Repository](#cloning-the-repository)
- [Setting Up Environment Variables](#setting-up-environment-variables)
- [Building and Running the Application](#building-and-running-the-application)
- [API Endpoints](#api-endpoints)
- [Kafka Integration](#kafka-integration)
- [Running Tests](#running-tests)

## Requirements

- Go 1.21 or above
- Docker and Docker Compose
- [Kafka](https://kafka.apache.org/) (running via Docker)

## Getting Started

This project uses Docker to manage dependencies and services, ensuring consistent environments for all developers.

### Cloning the Repository

1. Open your terminal.
2. Clone the repository:
   ```bash
   git clone https://github.com/krishnaag23/vidya-vaani.git
   cd vidya-vaani
   ```

### Setting Up Environment Variables

1. Copy the example environment file to create your local `.env` file:
   ```bash
   cp .env.example .env
   ```
2. Edit the `.env` file to configure your environment variables based on your local setup.

### Building and Running the Application

1. Build and start the application using Docker Compose:
   ```bash
   docker-compose up --build
   ```
  - This command will start the API, PostgreSQL database, Zookeeper, and Kafka services.
2. The application will be accessible at `http://localhost:8080`.

### API Endpoints

- **Health Check**
  - **Endpoint**: `GET /health`
  - **Description**: Checks if the API is running.
  - **Response**:
    ```json
    { "status": "up" }
    ```
- **Transcribe Audio**
  - **Endpoint**: `POST /transcribe`
  - **Description**: Accepts audio data and returns the transcription.
  - **Request Body**:
    ```json
    {
      "audio_data": "<base64_binary_data>"
    }
    ```
  - **Response**:
    ```json
    {
      "transcription": "The transcribed text"
    }
    ```

### Kafka Integration

The application integrates with Kafka to produce transcription messages. This allows the application to send real-time data to other services.

- Kafka Setup: The Kafka service is included in the docker-compose.yml file, and it will run automatically when you execute docker-compose up.
1. By default, the application communicates with Kafka through `localhost:9092`. This can be changed in the Kafka producer initialization.
2. The transcription output is produced to the `transcription_topic`.


### Running Tests

You can run tests in your application by executing:

```bash
go test ./...
```
