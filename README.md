# GOBİT Challange

This demo created for GOBIT Camp.

## System Design

```
exchangerateapi --> er-api-consumer --> rabbitmq --> er-rabbit-consumer --> postresql --> er-api --> public
```
## Configure

İnsert your ExchangeRate-API Key in .env file to ER_API_CONSUMER_KEY

```
ER_API_CONSUMER_KEY=xxxxxxxxxxxxxx
```

All other environment variables are in .env file. You can edit them to your taste.

## Build and Run

```console
docker-compose build
```

```console
docker-compose up
```

## Api Design 

### List last database record from given parite

! ATTENTION - Time is recorded and displayed in the database using the UTC time zone.

```
/api/exchange/{parite}
```

```web
GET  -  localhost:3000/api/exchange/TRY
```
response:

```json
{
	"data": {
		"ID": 4,
		"Time": "2022-04-25T18:35:14.655Z",
		"USD": 14.78,
		"EUR": 15.892473118279568,
		"TRY": 1
	},
	"message": "Exchange Found",
	"status": "success"
}
```
### List all database record

```
 /api/exchange
```

```web
GET  -  localhost:3000/api/exchange
```
response:

```json
{
	"data": [
		{
			"ID": 1,
			"Time": "2022-04-25T18:35:14.655Z",
			"USD": 1,
			"EUR": 0.93,
			"TRY": 14.78
		},
		{
			"ID": 2,
			"Time": "2022-04-25T19:18:19.271282Z",
			"USD": 1,
			"EUR": 0.9259,
			"TRY": 14.7667
		},
		{
			"ID": 3,
			"Time": "2022-04-25T19:20:07.988473Z",
			"USD": 1,
			"EUR": 0.9259,
			"TRY": 14.7667
		},
		{
			"ID": 4,
			"Time": "2022-04-25T19:21:08.028916Z",
			"USD": 1,
			"EUR": 0.9259,
			"TRY": 14.7667
		}
	],
	"message": "Exchanges Found",
	"status": "success"
}
```

