# Introduction

This is the bot that bridges the alert signal from TradingView to Alpaca broker for trading behaviors

### TradingView: 
* This is the data source and analysis tool of stock market. And it is used as the operation signal here with PineScript generated alert to call our server's webhook for triggering the trades
* link: https://www.tradingview.com/chart

### Ngrok: 
* This is the security tunnel to expose your local interface to the public Internet
* link: https://ngrok.com/use-cases/api-gateway

### Alpaca: 
* This is the broker for actual stock trading with decent APIs available
* link: https://app.alpaca.markets/paper/dashboard/overview

### Configuration
* Setup your Ngrok configues to make your local Ngrok client able to establish the security tunnel to the remote Ngrok server
```
export NGROK_USRNAME=<your own username>
export NGROK_PASSWD=<your own password>
export NGROK_AUTHTOKEN=<your own token>
```
* Setup the secret pairs for your Alpaca broker
```
export ALPACA_LIVE_KEY=<access key for your live account>
export ALPACA_LIVE_SEC=<secret key for your live account>
export ALPACA_PAPER_KEY=<access key for your paper account>
export ALPACA_PAPER_SEC=<secret key for your paper account>
```
* Setup the token which is specified in the webhook payload for operation signal authentication to avoid malicious attack. Normally, you can use `echo -n $username:$password | md5` to generate this token
```
export WEBHOOK_TOKEN=<your webhook token used for call access control>
```
