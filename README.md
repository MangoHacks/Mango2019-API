![MangoHacks 2019 Bandage](https://img.shields.io/badge/MangoHacks-2019-ed821e.svg?logo=data%3Aimage%2Fpng%3Bbase64%2CiVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAQAAABKfvVzAAADWElEQVQ4y22SfWxTZRTGHza21mF1OhlgR0YEhOBHFDI1fgYBP9CRYIbG6HTLFIJOJG6GzG2ZCpolxlmzNBvJlllpBrJsA7UZTi1jKaXOcrl2N0tXa725er17eXN36evleikxr38gyUb7%2FHee8zsnJ08OME%2F3A%2FDfEfucJNk%2F7IJyWthzqHQMR10Dq1ehHlnyorFAqKa%2FWpeMuBzUTrMZ%2B7J%2BMvhwqoFO%2FHB79Fr8BRTnx962%2FjbORl%2F0Ln3IWX%2F94Np4M1OZTpNEkj%2Fm2DAXfxAcY88as%2FKwr5xhGwCgHsDoI8aUIcuC%2BBlw99yBAI6U6hFyxrv86Bx3AzjGN5vn%2F%2BVkUqxpKei62tgIjvgbpiE%2BY2PHvEMbUZiX%2BjRzmQqMxmv5VduP3hL6Exna5%2FRkRSFDvNeaiTXI7frkidUAgFXgEJ83jVilCWeO7N530BF9oG%2BFLiabAADdaC%2BiAXqq54b%2B7KxxCziSHzDJe7PiIeMA3oQG4UlmSHUca5BLGcRfNf8Il6deoeqV%2FYvo1%2FTskdLjyK0Uxqt0bXBt9Ali4ClwiC%2BZafE1joostBGfFIV3dhcnqtlfwZXi02QWYXyz0pgko13F%2Fhy730PndeG6jptSbcZ0V0myhigYcOk%2BRiKP6rgrC%2B9El0PaLpZtLyAB%2Bi3ylP1aBIl6i8mNWPByFu6EjfAmg8QqR9YxTd7LoX6l9EGPku86XL4s%2FAEE0LtUO0nPdbqT7aYSXtNzI51I7QFLJ5o5yq%2FBi6Ggb5ly2EoLVcH7zPOajyN4J40Lj4GxZBPHknl4GyoLxjaSU9bFqXdDbvP7DFePFeZHn9PEQ2XQfiHH3nJ%2BBABYhlpsXehZEtqa%2BtJMM03c6Vms%2BuxZqVv%2FPbI50awGdjggNplpcbc7rwXACQy7lS%2F0aStjMvXw4Ppet9pvWcmmd1xaP51UflY6OOArIUPmBanB45IQRGAdOUdC8QPHK7YVjT5Oz1gXpdZWZxhjt5GgzRO7OSBieIUWsDPaiFDlL9u1qO%2FW10vaSiNbFJ%2FJzD%2BFuprCD3EPVAwsn%2FKGtiQAFybQs1g6YM7YGfYb%2BdEeIqP6tG1bac0fWr9wQe3%2FQbRgn2OXY%2B%2BVYj%2Bq88crpA5VorbJ9UskkTwY3tTqzPXs%2FwFTeMBV0Y%2Fb3gAAAABJRU5ErkJggg%3D%3D)
[![GoDoc](https://godoc.org/github.com/MangoHacks/Mango2019-API?status.svg)](https://godoc.org/github.com/MangoHacks/Mango2019-API)
![Travis Build Badge](https://travis-ci.org/MangoHacks/Mango2019-API.svg?branch=master)

# Mango2019-API

This is the backend of 2019's MangoHacks, herein, all database functionality and most data manipulation
events will occur.

The API is built with scalability in mind and is able to be scaled horizontally with ease.

## Running/Testing The Server

1.  `$ git clone` the repository into your `$GOPATH`.
2.  `$ go build` the binary within the directory the repository was cloned to.
3.  `$ ./Mango2019-API` to run the binary.
