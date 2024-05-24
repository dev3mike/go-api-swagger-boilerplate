![golang-api-boilerplate](https://github.com/dev3mike/go-api-swagger-boilerplate/assets/69217981/2f59fe1c-f87f-4d56-bb94-aee272bf97d8)

# Go API Boilerplate 🚀
Welcome to the Go API Boilerplate! This boilerplate is designed to help you quickly set up a robust and scalable API using Go. It incorporates several best practices and powerful libraries to make your development process smoother and more efficient.

## Features 🌟

1. 🌍 **Environment Validation with xEnv**
   - Ensure your environment variables are loaded from .env and validated and accessible in a type-safe manner.

2. ✅ **Input Validation with xMapper**
   - Validate JSON inputs, structs, and even single parameters using the powerful xMapper library.

3. 📄 **Perfect API Documentation with Swagger**
   - Generate comprehensive API documentation with Swagger, making your API easy to understand and use.

4. 🚦 **HTTP Router with go-chi**
   - Utilize the lightweight and powerful go-chi router for handling HTTP requests.

5. 📝 **Structured Logging with slog**
   - Implement structured logging to keep your logs clear and easy to manage.

6. 🕊️ **Database Migrations with goose**
   - Manage your database schema changes with ease using the goose migration tool.

7. 🍞 **Database ORM with bun**
   - Interact with your database using the bun ORM for efficient data management. Learn more at [bun.uptrace.dev](https://bun.uptrace.dev).

8. 🗂️ **Enterprise-Ready Folder Structure**
   - Follow an enterprise-level folder structure to keep your code organized and maintainable.

9. 🔐 **Authentication and Authorization with Clerk**
   - Secure your API with authentication and authorization using Clerk with implemented webhook.

10. 🧪 **Best Practices for Dependencies and Testing**
    - Manage dependencies and make your code testable by following industry best practices.

11. ⚠️ **Custom Error Handling**
    - Implement custom error handling to improve error management across your API.

12. 📜 **Makefile for Easy Commands**
    - Use the Makefile to run common commands effortlessly.

13. 🛠️ **Easily Change Module Name**
    - Change the module name easily with a built-in script.

14. 🌱 **Database Seeding**
    - Seed your database with initial data for testing and development.


**This repository is under development** 🛠️

# Makefile Documentation 📑

| Command             | Description                                      |
|---------------------|--------------------------------------------------|
| `make run`          | Run the server                                   |
| `make swagger`      | Generate Swagger documentation                   |
| `make db-status`    | Check database migration status                  |
| `make up`           | Apply database migrations                        |
| `make down`         | Revert the last database migration               |
| `make delete-all`   | Delete all migrations                            |
| `make reset`        | Reset the database and reapply all migrations    |
| `make seed`         | Seed the database                                |
| `make reset-seed`         | Reset and delete all table, run the migration and seed the database                                |
| `make change-mod-name MODULE_NAME=new/package/name` | Change module name |

#### Database Migrations
To handle database migrations, we use Goose. First, make sure you have installed Goose by following the instructions here: [Goose GitHub](https://github.com/pressly/goose).

To create a new migration, use this command:
```sh
goose create TITLE_OF_MIGRATION sql
```

Then, move the generated SQL file to the `migrations` folder and update it as needed.


### Installing Dependencies

To install all the dependencies listed in the `go.mod` file, use the following command:

```sh
go mod download
```

#### License
This code is made available under the MIT License. This means you can use it freely in your personal and commercial projects. For more details, see the LICENSE file in the repository.
