# Golang Backend Exercise

Welcome to the programming exercise for Golang Backend Developers! :)

## Your Task

Your task is to implement a very simple account management microservice using Golang. It should at least be able to:

* Register new users using their email address, password, and name
* Log the user in using their email address and password
* Retrieve a list of all active users
* Change a user's name
* Delete a user

You should include some unit tests. Please use this repository as a starting point for your implementation and employ appropriate best practices.

As the functional requirements set above should be trivial to implement, we focus our evaluation on the quality of your code and architecture. Generally, your goal is to get the service as close as possible to a production-ready state. This could, for example, include adding:
* Proper logging capabilities
* Metrics export to Prometheus
* Health checks for Kubernetes

## Setup

1. Install [Golang](https://go.dev), [Make](https://www.gnu.org/software/make/) and [Docker](https://docs.docker.com/get-docker/).
2. After cloning the repository, run `make` to verify you can build the Go app.
3. Run `make database-up` to start the development database in a Docker container.
4. Start the application using `make run`.
5. Happy coding!

The app exposes the REST API on port `8080`.

If you would like to drop the database and its contents, run `make database-down` and then `make database-up` again to start with a clean database.

If you would like to perform your SQL migrations, run `make -C tasks/migrate up`, and to roll them back, `make -C tasks/migrate down`.
 
Please take a look into the `Makefile` to see all available operations.

## File Structure

The contents of this repo are structured in the following way:

* `app/`: The application source files
  * `account/`: The package containing the `account` domain model and business logic
  * `storage/`: The package implementing persistend storage in PostgreSQL
  * `web/`: The package exposing REST API for account management
* `tasks/migrate/`: The source code and .sql files required for database migrations
* `bin/`: Build artifacts (excluded from version control)

## Contact

In case of any questions or problems you may have, please reach out to me at [marcin.praski@immunkarte.de](mailto:marcin.praski@immunkarte.de).
