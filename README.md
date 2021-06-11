## Overview
:musical_note: This is a backend server part of MusicPlayer written in Go (with gin-gonic framework) and uses MSSQL to store data. Interaction with backend realized by REST API.

## Features
* :key: JWT-tokens authentication
* :man: User's accounts (create, delete, modify)
* :dollar: Placeholder for premium features
* :headphones: Creating and modifying tracks
* :notes: Managing tracks by playlists
* :tv: Support video
* :paperclip: Attachment media files to track

## Usage

### Configure backend
Before start you should go through some steps:
1. Create database with sql-script in schema folder
2. Set up common config in `configs/config.toml`
3. Write sensitive credentials to .env file

Example of .env file:
```
DB_PASSWORD="password1234"

MAIL_BOX="user@example.com"
MAIL_PASSWORD="password1233"

# or change it to release
GIN_MODE="debug"
```

And now to run you can just type `make run`

### Client-side
REST API client easy to write, so you can make your own client. As example (ugly example) you can use this [project](https://github.com/wellWINeo/MusicPlayer). The REST-API client part collected in `api` folder and not as bad as other parts =)


## TODO

- [ ] Currently media-files storing in a filesystem, move it to DB
- [ ] Make client not so awful
