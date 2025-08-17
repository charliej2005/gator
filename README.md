# Gator
A blog aggre**gator** CLI written in Go using PostgreSQL, SQLc, Goose and psql.
This project is part of the [Boot.dev guided project series](https://www.boot.dev/courses/build-blog-aggregator-golang).

## Features

- User login and registration system
- Live retrieval of data from RSS feeds
- Feed following and unfollowing
- Storage of posts to PostgreSQL database

## Usage

1. Install the Gator CLI directly using Go:

    ```sh
    go install github.com/charliej2005/gator@latest
    ```

2. Create a new database in PostgreSQL.

3. Create a config file in your home directory named `.gatorconfig.json` with the following contents:

    ```json
    {
    "db_url": "<postgres_connection_string>",
    "current_user_name": ""
    }
    ```

    Replace `<postgres_connection_string>` with the PostgreSQL connection string to the database you created.

4. Make sure you have [Goose](https://github.com/pressly/goose) installed. Then, from `/sql/schema`, run:

    ```sh
    goose postgres "<postgres_connection_string>" up
    ```

    Replace `<postgres_connection_string>` with the same connection string used in your `.gatorconfig.json` file.  

5. Use the CLI by running:

    ```sh
    gator <command>
    ```
    Where `<command>` is any of the commands listed below.

## Commands

- `register <username>` — Register a new user
- `login <username>` — Log in as an existing user
- `users` — List all users

- `addfeed <name> <url>` — Add a new RSS feed to follow
- `feeds` — Lists all added feeds

- `follow <url>` — Follows the RSS feed with the given URL
- `unfollow <url>` — Unfollows the RSS feed with the given URL
- `following` — Lists all followed feeds

- `agg <duration>` — Aggregate a feed's posts every duration (e.g., `agg 1m` for an update every minute)
- `browse <limit>` — Shows the `<limit>` most recent posts across all aggregated feeds. `<limit>` defaults to 2.

- `reset` — Clears the database (be careful with this one!)

> **Note:** typing an invalid command will result in an error.

## Requirements

- Go v1.22+
- PostgreSQL
- Goose
