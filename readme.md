# Validate [![Test](https://github.com/markelog/validate/workflows/Test/badge.svg?branch=master)](https://github.com/markelog/validate/actions)

> Simple validation HTTP-service

## Intro

It's a simple HTTP validation service, with one available route `/email/validate`. Which validates the email via `POST` request

## Start

```sh
$ docker run -t -p 8080:8080 -e PORT=8080 markelog/validate
```

## Example

### Request

`POST /email/validate`

```sh
$ curl -XPOST -d '{"email":"markelog@gmail.com"}' http://localhost:8080/email/validate
```

### Response

```sh
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Vary: Origin
Date: Mon, 24 Feb 2020 12:18:58 GMT
Content-Length: 150

{"valid":true,"validators":{"dmarc":{"valid":true},"domain":{"valid":true},"regexp":{"valid":true},"reputation":{"valid":true},"smtp":{"valid":true}}}
```

### Checks

- `regexp` – checks the syntax of the provided email
- [`dmarc`](https://en.wikipedia.org/wiki/DMARC) – checks DMARC related presence in the domain DNS
- `domain` – checks if email domain exist
- `smtp` – establishes connection to the SMTP service, sends `RCPT TO` request thus checking if such email address exist
- `reputation` – checks the [reputation](https://en.wikipedia.org/wiki/Reputation_system) of the email via https://emailrep.io/.
  - _Note:_ Amount of request from one IP-address is limited, it's better to provide the key (see [`/.env.example`](https://github.com/markelog/validate/tree/master/.env.example)). If limit is exceed, response from `emailrep.io` will not be present in the response

## Development

### Commands

- `make install` – installs stuff
- `make dev` – starts the server and watches the changes
- `make unit-tests` – executes the unit-tests
- `make integrations-tests` – executes the integration-tests
- `make test` – executes unit and integration tests
- `make lint` – lint GO code via multiple linters
