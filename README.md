# test-project-1


Framework used in this project: <br>
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=blue)](https://github.com/labstack/echo)

# Table of Content
- [Description](#description)
- [How to use](#how-to-use)
- [Endpoints](#endpoints)
- [Credits](#credits)

# Description
This is submition of coding test of klikA2C
```
In this project i use MVC architecture because it's the most common achitecture used by programmer
```

# How to use
- Install Go and MySQL or (install docker and docker-compose)

- run this project
```
$ go run main.go or $ docker-compose up --build
```
# Endpoints

| Method | Endpoint | Description| Authentication 
|:-----|:--------|:----------| :----------:|
| POST  | /login | Login existing user | No 
| POST | /register | Register a new user| No 
|---|---|---|---|
| GET | /additem | add item | Yes 
| GET | /edititem/:id | edit item by id | Yes 
| GET | /getitem | Get list of all item | Yes 
| GET | /deleteitem/:id | delete item by id | Yes
|---|---|---|---|
| GET | /newtransaction | create new transaction | Yes 
|---|---|---|---|

```
- example POST
http://localhost:8080/newtransaction
-JSON Input 
{
	"item_input":[
		{
			"id": 1,
			"quantity": 10
		},
		{
			"id":2,
			"quantity": 25
		}
	]
}
```

```
- example GET
http://localhost:8080/getitem
```

<br>

## Credits

- [Darien Kentanu](https://github.com/darienkentanu) (Author and maintainer)
