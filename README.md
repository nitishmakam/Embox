# Embox

Embox is an extemely simple Web application written in Go. This repository is an attempt to familiarize myself with the following technologies:
* Go
* Web Development in Go using Gin
* Mongo DB
* Docker
* Docker Compose

## Technologies & Libraries Used
* Go (1.22.1)
* Gin - Go Web Framework
* Mongo DB - As a Datastore
* MongoDB Go Driver - To connect Go backend to MongoDB
* Docker
* Docker Compose - For building and deploying the required containers to run the application

## API Spec

- <code>GET</code> /movies - List all the movies
- <code>GET</code> /movies/:id - Return a movie given it's ID
- <code>POST</code> /movies - Create a movie
- <code>PUT</code> /movies/:id - Update a movie given it's ID
- <code>DELETE</code> /movies/:id - Delete a movie given it's ID

Request/Response Payloads will have the following structure - 
```json
{
  "id": "6abc6gd7hete7dab7e6",
  "name": "Guardians Of The Galaxy"
}
```

## System Requirements
- Docker on a Linux/Mac Environment

## How to Run
1. Clone the repository to your machine.
2. Navigate to the cloned repo via the terminal.
3. Execute
 ```bash
   docker compose up
   ```
4. The web application would be providing services on port 8443

## Example API Requests
```http
GET http://localhost:8443/movies
```
```http
GET http://localhost:8443/movies/46ab4db57se2bn4lo4
```
```http
POST http://localhost:8443/movies
Content-Type: application/json

{
 "id": "46ab4db57se2bn4lo4",
 "name": "The Godfather"
}
```
```http
PUT http://localhost:8443/movies/46ab4db57se2bn4lo4
Content-Type: application/json

{
 "id": "46ab4db57se2bn4lo4",
 "name": "The Dictator"
}
```
```http
DELETE http://localhost:8443/movies/46ab4db57se2bn4lo4
```
