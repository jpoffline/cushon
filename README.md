# Investor deposit API - Natwest Cushon

- Submission: Jonathan Pearson

This is my submission for the Natwest Cushon Software Engineer Recruitment Scenario.

I have created a backend API which implements the workflow and restrictions laid out in the document provided to me. The implementation is as a golang webserver with postgres database data persistence.

## Business requirement

> A `customer` can make a `deposit` of a specific `amount` into a `fund`. Both the customer and the fund must exist in the system for the deposit to be accepted.

This functionality is enforced in two layers:

1. **Application level validation in the go code** First check that the `fundId` exists in the db, then that the `customerId` exists, and then persist the deposit in the db. This approach allows for testing and explicit error messaging as to "why" the deposit may have failed.

2. **Database-level integrity via foreign keys** While not strictly necessary when controlling writes through the API, this adds a layer of protection for direct database interactions. The tradeoff is that error messages from database constraint violations are less user-friendly.

When a deposit is made, we store `when` the deposit was made along with the deposit `amount`. We link to `funds` and `customers` via `uuid`s only, so that storage and interaction with PII is minimised.

## Tech stack

I have made various decisions in this project:

- `postgres` for data persistence - SQL rather than NoSQL seems to model the data structures pretty well. Would allow for indexing, scaling, e.g., in the future.
- `golang` API to create the backend - a wonderful, statically typed compiled language. Simple to work with, excellent development ecosystem.
- `gin` as the HTTP framework
- `ent` as a database ORM - this is a fast, "meta"-backed ORM giving type safety, good query generation. Can be bad if needing extremely complex efficient queries to be built.
- `mockery` for mocking interfaces
- `github workflow` for CI - running the test suite "on push" (this is a very basic CI-flow for now)

There are three models: `fund`, `customer`, and `deposit` (these are represented in the ORM).

- `fund` with fields `id`:uuid, `name`:str
- `customer` with fields `id`:uuid, `name`:str
- `deposit` with fields `id`:uuid, `fund_id`:uuid&fk, `customer_id`:uuid&fk, `amount`:float, `created_at`:datetime

The `fund`, `customer`, and `deposit` model instances are stored in a SQL database (rather than NoSQL) - the models lend themselves naturally to being tabular. Will also allow for simple update and cross-model queries.

The overall architecture separates out implementation from usage where practical in such a simple example application, via interfaces. This allows for dependency injection, which significantly improves testability and enhances maintanance/extensibility.

Unit and integration tests cover the deposit service and several HTTP endpoints, particularly those involving business logic.

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
