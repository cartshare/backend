# backend
> Backend service for CartShare

Account management
- [POST /login](#post-login)
- [POST /signup](#post-signup)
- [GET /session](#get-session)
- [GET /logout](#get-logout)

List management
- [GET /list](#get-list)
- [POST /createItem](#post-createitem)
- [POST /completeItem](#post-completeitem)
- [POST /setItemWishlisted](#post-setitemwishlisted)

Neighbor features
- [GET /neighborList](#get-neighborlist)

Notification features
- [GET /notifications](#get-notifications)
- [POST /deleteNotification](#post-deletenotification)

## Account Management

### POST `/login`

Request:

```js
{
	"email": "me@ethanent.me",
	"password": "letmein"
}
```

Response:

*Sets cookie for session.*

```js
{
	"error": null
}
```

### POST `/signup`

Request:

```js
{
	"name": "Ethan Davis",
	"email": "me@ethanent.me",
	"password": "letmein",
	"address": "82 Bayview Ave, Mendocino, CA"
}
```

Response:

*Sets cookie for session.*

```js
{
	"error": null
}
```

### GET `/session`

Response:

```js
{
	// Will be non-null if session does not exist
	"error": null,
	"email": "me@ethanent.me"
}
```

### GET `/logout`

Response:

*Deletes session cookie.*

## List Management

### GET `/list`

Response:

```js
{
	"items": [
		{
			"id": "d7d9a826dc100f728ddfc39e11245c",
			"desc": "Target Frosted Wheats Cereal",
			"qty": 2,
			"onWishlist": false
		},
		{
			"id": "4c57859d1a8a19572f8e3574c6e4bb",
			"desc": "Apple",
			"qty": 8,
			"onWishlist": true
		}
	]
}
```

### POST `/createItem`

Request:

```js
{
	"itemDesc": "Target Frosted Wheats Cereal",
	"itemQty": 2
}
```

Response:

```js
{
	"error": null
}
```

### POST `/completeItem`

Request:

```js
{
	"itemId": "d7d9a826dc100f728ddfc39e11245c"
}
```

Response:

```js
{
	"error": null
}
```

### POST `/setItemWishlisted`

Request:

```js
{
	"itemId": "d7d9a826dc100f728ddfc39e11245c",
	"wishlisted": false
}
```

Response:

```js
{
	"error": null
}
```

## Neighbor Features

Note: To complete a neighbor's request, use the `/completeItem` endpoint.

### GET `/neighborList`

Response:

```js
{
	"neighborRequests": [
		{
			"id": "4c57859d1a8a19572f8e3574c6e4bb",
			"desc": "Apple",
			"qty": 8,
			"owner": "Ethan Davis"
		},
		{
			"id": "d7d9a826dc100f728ddfc39e11245c",
			"desc": "Birthday Cake Mix",
			"qty": 1,
			"owner": "James Smith"
		}
	]
}
```

## Notification Features

### GET `/notifications`

Response:

```js
{
	"error": null,
	"notifications": [
		{
			"id": "b78a3f5ca090fbfca060e0e590d523",
			"title": "Request for 2x Apple Fulfilled",
			"body": "Your neighbor, Ethan Davis, has completed your request.",
			"created": "2020-06-20T22:50:50.464944-07:00"
		}
	]
}
```

### POST `/deleteNotification`

Request:

```js
{
	"notificationId": "b78a3f5ca090fbfca060e0e590d523"
}
```
Response:

```js
{
	"error": null
}
```