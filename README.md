# backend
> Backend service for CartShare

## Account Management

### POST `/login`

Request:

```js
{
	"username": "ethan",
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
	"username": "ethan",
	"password": "letmein",
	"address": "82 Bayview Ave, Mendocino, CA",
	"name": "Ethan Davis"
}
```

Response:

*Sets cookie for session.*

```js
{
	"error": null
}
```

## List Management

### GET `/list`

Request:

```js
{
	"items": [
		{
			"id": 1,
			"desc": "Target Frosted Wheats Cereal",
			"qty": 2,
			"onWishlist": false
		},
		{
			"id": 2,
			"desc": "Apple",
			"qty": 8,
			"onWishlist": true
		}
	]
}
```

Response:

```js
{
	"error": null
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
	"itemId": 2
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
	"itemId": 2,
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

Request:

```js
{
	"neighborRequests": [
		{
			"id": 2,
			"desc": "Apple",
			"qty": 8,
			"owner": "Ethan Davis"
		},
		{
			"id": 9,
			"desc": "Birthday Cake Mix",
			"qty": 1,
			"owner": "James Smith"
		}
	]
}
```
