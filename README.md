# GOBİT Challange

This demo created for gobit camp.

```console
docker-compose build
```

```console
docker-compose up
```

## System Design

```mermaid
  graph TD;
      A-->B;
      A-->C;
      B-->D;
      C-->D;
```

exchangerateapi --> er-api-consumer --> rabbitmq --> er-rabbit-consumer --> postresql --> er-api --> public

### er-api 

#### GET /api/exchange/{parite}

List last database record from given parite

```web
localhost:3000/api/exchange/TRY
```

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

#### GET /api/exchange/

List all database record

```web
localhost:3000/api/exchange
```

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
			"USD": 0,
			"EUR": 0.9259,
			"TRY": 14.7667
		},
		{
			"ID": 3,
			"Time": "2022-04-25T19:20:07.988473Z",
			"USD": 0,
			"EUR": 0.9259,
			"TRY": 14.7667
		},
		{
			"ID": 4,
			"Time": "2022-04-25T19:21:08.028916Z",
			"USD": 0,
			"EUR": 0.9259,
			"TRY": 14.7667
		}
	],
	"message": "Exchanges Found",
	"status": "success"
}
```

