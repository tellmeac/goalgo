# Описание back-end API

query unix time. returns json body (needPoint: Boolean | null)

1. GET /updates?from=1701781200

```json
{
  "lastStamp": 1701954000,
  "stamps": [
    {
      "x": 1701781200,
      "y": {
        "open": 1.0,
        "high": 1.0,
        "low": 1.0,
        "close": 1.0
      },
      "topLine": 1.0,
      "downLine": 1.0,
      "blueLine": 1.0,
      "needPoint": true
    }
  ]
}
```

2. GET /chart?from=1701781200&to=1701954000

```json
{
  "stamps": [
    {
      "x": 1701781200,
      "y": {
        "open": 1.0,
        "high": 1.0,
        "low": 1.0,
        "close": 1.0
      },
      "topLine": 1.0,
      "downLine": 1.0,
      "blueLine": 1.0,
      "needPoint": true
    }
  ]
}
```

