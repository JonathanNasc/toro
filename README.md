# Toroburro

It's just a stupid script I use to update personal finances
## Set up

Create a non-versioned `settings.json` file. Here is a template:

```
{
    "sources":[
        {
            "name": "alphavantage",
            "apikey": "senha1234"
        }
    ],
    "assets":[
        {
            "source": "alphavantage",
            "codes": [
                "GOOGL",
                "AAPL",
                "AMZN",
                "PYPL",
                "TOBU"
            ]
        }
    ]
}

```

## Usage

- To build:

`go build main.go`

- To run:

`./go-toroburro`

## CSV output example

```
asset,price,date
GOOGL,12.2,2022-01-26
AAPL,10.3,2022-01-26
AMZN,1.99,2022-01-26
PYPL,20.0,2022-01-26
TOBU,3500.00,2022-01-26
```