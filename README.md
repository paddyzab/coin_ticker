### Coin ticker

It's a small utility tool which notifies you inside of your terminal about changes in cryptocurrencies changes.

#### Bag Value
Tracking value of your bags.
Create cointicker_config.yaml file in your home directory.
Example you will find in `example_config.yaml`


Add in `coins`, coin identifier and quantities you want to track.

System will calculate aggregated value of all your bags, based on the value from coinmarketcap (not impl).

#### Building
To build use makefile.

`make build`

or

`make install`

Data are fetched from [coinmarketcap api](https://coinmarketcap.com/api/).
