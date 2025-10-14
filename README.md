# Digantara Backend Engineer (Golang) Assignment

**Note:** _This README.md file will explain the project structure and the tech stack I have used to accomplish this assignment.
To know assignment-specific details, please check `apps/scheduler-service/README.md` file._

## 📋 Project Overview

This project provides a foundation for building RESTful APIs with Go. It includes:

- A clean architecture design
- API endpoint examples with request/response validation
- Database integration using Bun ORM
- Error handling and logging
- Environment-based configuration



## 🏗️ Project Structure

```
apps/
  scheduler-service/           # Main API service
    app/                 # Application core
      scheduler/          # Logic for scheduling the Jobs
      setup/             # App configuration and setup
      shared/            # Shared resources
      job/               # handling Jobs logic 
    migrations/          # Database migrations
    routes/              # API route definitions
    internal-lib/      # Internal libraries
        snowflake/       # Snowflake ID implementation
        utils/           # Utility functions
    Dockerfile           # Container definition
    example.env          # Environment variable example
    go.mod               # Go module definition
    main.go              # Application entry point
```

## 🚀 Features

- RESTful API endpoints using Huma framework
- Chi router for HTTP request handling
- Database integration using Bun ORM
- Snowflake ID generation for distributed systems
- Environment-based configuration
- Dependency injection with Wire

## 📚 Documentation

API documentation is automatically generated through the Huma framework and available at the `/docs` endpoint when the server is running.


## 📦 Dependencies

- [Huma](https://github.com/danielgtaylor/huma/): API framework
- [Chi Router](https://github.com/go-chi/chi): HTTP routing
- [Bun](https://github.com/uptrace/bun): SQL ORM
- [Sentry](https://github.com/getsentry/sentry-go): Error tracking
- [Wire](https://github.com/google/wire): Dependency injection

