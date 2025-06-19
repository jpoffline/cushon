# Investor deposit

Natwest Cushon

- Submission: Jonathan Pearson

This is my submission for the Natwest Cushon Software Engineer Recruitment Scenario.

I have created an API which implements the workflow and restrictions laid out in the document provided to me. The implementation is as a golang webserver with postgres database data persistence.

## Business requirement

> A `customer` can make a `deposit` with an `amount` into a `fund`. The fund and customer must exist in the system in order for the deposit to be allowed.

This has been implemented as a requirement within the solution in two ways:

- via explicit `go`-code business logic: first check that the `fundId` exists in the db, then that the `customerId` exists, and then persist the deposit in the db. This approach allows for a) testing and b) explicit error messaging as to "why" the deposit may have failed.

- via foreign key requirements on the database schema (`edges` in `ent`). This approach "shouldnt be needed" but would be useful if another consumer were to attempt to write to the db. This could be used alone, but would have the drawback of having fairly cryptic error messages should a deposit write fail.

When a deposit is made, we store `when` the deposit was made along with the deposit `amount`. We link to `funds` and `customers` via `uuid`s only, so that storage and interaction with PII is minimised.

## Tech choices

I have made various decisions in this project:

- `postgres` for data persistence - SQL rather than NoSQL seems to model the data structures pretty well.
- `golang` API to create the backend - a wonderful, statically typed compiled language. Simple to work with, excellent development ecosystem.
- `gin` as the HTTP framework
- `ent` as a database ORM - this is a fast, "meta"-backed ORM giving type safety, good query generation. Can be bad if needing extremely complex efficient queries to be built.
- `mockery` for mocking interfaces
- `github workflow` for CI - running the test suite "on push" (this is a very basic CI-flow for now)

There are three models: `fund`, `customer`, and `deposit` (these are represented in the ORM).

The `fund`, `customer`, and `deposit` model instances are stored in a SQL database (rather than NoSQL) - the models lend themselves naturally to being tabular. Will also allow for simple update and cross-model queries.

The overall architecture separates out implementation from usage where practical in such a simple example application. This allows for dependency injection, which significantly improves testability and enhances maintanance/extensibility.

Interfaces are used - so that each service can define the interface of dependencies, and the application can give a service that implements the interface.

Automated tests for some of the HTTP endpoints, and for the deposit service are provided - this has the business logic and so deserves good tests.

## Usage

```
docker-compose up
```

The project has been configured to allow hot-reloading to aid developer experience.

This command will bring up a `postgres` server and the `golang` api.

Visit the [swagger docs](http://localhost:8080/swagger/index.html) to view and execute the endpoints.

### Endpoints implemented

#### `POST /fund`

- Create a fund
- Pass in the json body

```json
{
  "name": "FundName"
}
```

#### `GET /fund/list`

- List all the funds

#### `POST /customer`

- Create a customer
- Pass in the json body

```json
{
  "name": "Customer Name"
}
```

#### `GET /customer/list`

- List all the customers

#### `GET /customer/{customer_id}/deposits`

- Get all deposits for the customer

#### `POST /deposit`

- Create a deposit for the customer into the fund
- Pass in the json body

```json
{
  "amount": 42.84,
  "customer_id": "009a5ab6-3c47-4912-b5c3-43661b0ef193",
  "fund_id": "c308d3b5-045e-4a9c-9245-c311e61d1112"
}
```

Note that the `customer_id` and `fund_id` must already exist in the system, otherwise the insertion isnt allowed. In these cases, the server returns with `500` and an error message explaining what the issue is.

## Work for the future

- logging and tracing (e.g., integrate with a structured logger like `zap` and/or tooling like `datadog`)
- middleware (e.g., auth, CORS)
- improve architecture: decouple the `dto` models from the `ent` ones.
- improve the CI piece for safer deployments (e.g., env-variable management)
- enhancing the model meta data (e.g., data about the funds may be useful to filter on. This would require more business context)
- setup githooks to run "unit tests" on commit (running integration tests can take too long, but should definitly be run in the full CI/CD suite)
- setup branching/PR policies
- move ent schema to own folder
- implement extra functionality: e.g., bulk deposits (different amounts into different funds in one transaction)
- error handling: different HTTP error codes for different reasons for failure
- swapping out the frameworks and technologies - after heavy useage it may become apparent that, e.g., `ent` or `postgres` arent suitable solutions. The codebase as written can have these swapped out for other solutions (e.g., GORM, or MonogoDb as appropriate).
