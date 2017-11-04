### Coin ticker

It's a small utility tool which notifies you inside of your terminal about changes in cryptocurrencies changes.

#### Bag Value
Tracking value of your bags - not completed.
Create cointicker_config.yaml file in your home directory.
Example you will find in `example_config.yaml`

Add in `coins`, coin identifier you want to track.

#### Building
To build use makefile.

`make build`

or

`make install`

for testing run

`make test`

Data are fetched from [coinmarketcap api](https://coinmarketcap.com/api/).
