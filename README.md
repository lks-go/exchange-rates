# exchange-rates

The util converts currency rates using https://exchangeratesapi.io

## Usage

For example,  you can go to the root of the repository and execute next command

```
go run cmd/app/main.go 25.50 USD RUB
```

After that you'll get USD converted to RUB and multiplied to 25.50.
In other words you will see how many rubbles you would buy for $25.50.

## Configuration

You have to set you own API key of your exchangerates account
to the config.yml which is in the root. Without that you will not have access
to get rates.

```
exchangeratesapi:
  accessKey: your_key
```