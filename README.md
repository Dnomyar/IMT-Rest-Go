# Rest API implementation in Go

The objective of this project is to implement a Rest API (Crud) in Go.

This project is not finished yet.

## Api contract
You may look at the API contract at : https://listedepokemon.restlet.io

### CREATE
Require :
 * Authorization

Parameters :
 * Pokemon to create

Success :
 * 201 CREATED

Failure :
 * 400 BAD REQUEST
 * 401 UNAUTHORIZED

### READ
Require :
 * Empty

Parameters :
 * Empty OR
 * Identifier

Success :
 * 200 OK -- return a collection OR the specific item if identifier given

Failure :
 * 404 NOT FOUND

### UPDATE
Require :
 * Authorization

Parameters :
 * Pokemon to update

Success :
 * 200 OK

Failure :
 * 400 BAD REQUEST
 * 401 UNAUTHORIZED
 * 404 NOT FOUND

### DELETE
Require :
 * Authorization

Parameters :
 * Identifier of the pokemon to delete

Success :
 * 200 OK

Failure :
 * 400 BAD REQUEST
 * 401 UNAUTHORIZED
 * 404 NOT FOUND