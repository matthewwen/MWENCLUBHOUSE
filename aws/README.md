# AWS SDK with GoLang SDK

The mwenclubhouse runs on a GoLang server. There are two repo that handle the server. One Repo is called gateway, which is the main application where it does the API hosting (which is connected to my home). The other one is called the sdk, where I write methods in which the gateway will call those methods.

## Why split the SDK and Gateway
I split it so I can easily test the SDK to see if it is working or not.

## What happened to this code?
AWS is no longer used as a cloud provider; it cost too much money. The mwenclubhouse now uses Firebase as a database. 