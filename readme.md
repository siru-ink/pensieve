# Pensieve

![Go Version](https://img.shields.io/badge/go-1.26-blue)
![License](https://img.shields.io/badge/license-BSD--3--Clause-green)
![Self-hosted](https://img.shields.io/badge/self--hosted-yes-blue)
[![Go Report Card](https://goreportcard.com/badge/git.siru.ink/siru/pensieve)](https://goreportcard.com/report/git.siru.ink/siru/pensieve)

A minimal, dependency-free Go server for collecting anonymous comments via a single HTTP endpoint.

## Installation

Compile the `main.go` file from source via

```
go build
```

or download an appropriate binary from the [releases](https://git.siru.ink/siru/pensieve/releases) page. Due to the way Go functions, this should be a completely independent binary. Simply put it wherever you want, and run it under a user that has permissions to create the database file in the current working directory. Alternatively, you can look at the [systemd service config file](/systemd.service) for a suggestion on how to initiate the server via systemd.

## Use Cases

- Static websites that need a lightweight comment backend (i.e. [Bearblog](https://bearblog.dev))
- Personal blogs with manual moderation workflows
- Anonymous feedback collection
- Minimalist self-hosted alternative to Disqus-like systems

## Features

- **Simple HTTP API**: Single endpoint for anonymous comment submission
- **SQLite Storage**: Lightweight, file-based database with no external dependencies
- **Concurrency Safe**: Handles multiple simultaneous requests safely
- **Automatic Setup**: Creates database and table structure on first run

## Technology Stack

- **Go**: Quick, memory-safe, concurrent language
- **SQLite**: Via the standard modernc.org/sqlite dependency (pure Go implementation)
- **Built-in HTTP Server**: net/http package

## API Endpoint

```
POST /
```

Accepts form-encoded comment data and stores it in the database.

**Form Parameters**:

| Parameter |  Type  | Description|
| --------- | ------ | ---------- |
|   name    | string | Commenter's name (required) |
|  comment  | string | The comment text (required) |
|  siteurl  | string | URL to redirect back to after submission (required)|

## FAQ

<details>
<summary>What about CSRF?</summary>

Since this server does not handle any authentication in order to be able to receive anonymous comments, no CSRF protection is necessary.

</details>

<details>
<summary>What about spam?</summary>

The concept for this server was always to allow anyone to post a comment without necessarily having to identify themselves first. This invites spam, but allows for uninhibited communication as well. The suggestions on prevent spam are: 1, do not directly display submitted comments to end users. Instead, they first need to be manually filtered and checked. 2, as seen in the [nginx config file](/nginx.conf) it would be good to run this behind a reverse proxy that implements rate limiting of some variety.

</details>

<details>
<summary>Is Pensieve production-ready?</summary>

It depends on your use case. Pensieve is intentionally minimal and does not include authentication, rate limiting, or spam protection. It is best used behind a reverse proxy and with manual moderation workflows.

</details>

<details>
<summary>Can I configure the port or database location?</summary>

Not currently. These values are hardcoded but can be modified in the source code. Adding environment variable support is a potential future improvement.

</details>

<details>
<summary>What happens if required fields are missing?</summary>

The server does not currently enforce validation. Empty values may still be stored in the database. Validation should be handled by the client or a reverse proxy.

</details>

<details>
<summary>Why is there no endpoint to retrieve comments?</summary>

Pensieve is designed as a write-only ingestion service. Reading, filtering, and displaying comments is expected to happen in a separate system or workflow.

</details>
