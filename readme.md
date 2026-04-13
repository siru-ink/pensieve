# Pensieve

A lightweight Go web server for collecting and storing, possibly anonymous, blog comments in a SQLite database. Pensieve provides a simple HTTP endpoint that accepts POST requests and persists comments for later retrieval.

## How to Run?

Compile the `main.go` file from source via

```
go build
```

or download an appropriate binary from the [releases](https://git.siru.ink/siru/pensieve/releases) page. Due to the way Go lang functions, this should be a completely independent binary. Simply put it wherever you want, and run it under a user that has permissions to create the database file in the current working directory. Alternatively, you can look at the [systemd service config file](/systemd.service) for a suggestion on how to initiate the server via systemd.

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

The concept for this server was always to allow anyone to post a comment without necessarily having to identify themselves first. This invites spam, but allows for uninhibited communication as well. The suggestions on prevent spam are: 1, do not directly display submitted comments to end users. Instead, they first need to be manually filtered and checked. 2, as seen in the [nginx config file](/nginx.conf) it would be good to run this behind a reverse proxy that implements rate limiting of some varitey.

</details>
