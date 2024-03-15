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

Embox is a simple unauthenticated CRUD application for managing names of movies. The application exposes the API endpoints:
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
