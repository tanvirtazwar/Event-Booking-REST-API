# EVENT BOOKING REST API
## Project  

Building a REST API with Go.

## Features

- Event Booking API.
- User Authentication (signup & login)
- Authorization (only authenticated users can create/update events).
- User-based event ownership protection (only creators can update/delete their events).
- JSON Web Token (JWT) based authentication.

## Project Stack

- **Backend:** *Go language* 
- **Database:** *sqlite3*
- **Authentication:** *JWT*

## How to Run the Project:

### Prerequisites:

- Install Go language according to the official instructions: https://go.dev/doc/install.
- Go SQLite Package: https://github.com/mattn/go-sqlite3.

### Setup:

[Clone this repository.](https://github.com/tanvirtazwar/Event-Booking-REST-API.git)
   ```
   Bash
   git clone https://github.com/tanvirtazwar/Event-Booking-REST-API.git
   ```

### Run the Project:

   ```
   Bash
   go run main.go
   ```
