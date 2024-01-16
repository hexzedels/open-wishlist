# Open Wishlist

Want to focus on privacy? Try this simple self-hosted telegram bot for managing your wishlists.

## Features

* Create multiple wishlists
* Switch between wishlists
* Add items to selected wishlist
* Show all items from wishlist

## How to run

1. Clone repo
2. Create your own bot at [@BotFather](https://t.me/BotFather)
3. Copy token
4. Run bot and database with docke compose
   ```shell
   sudo -E TOKEN=<your_bot_token> docker compose up -d
   ```

## How to debug

Run this command to read container logs

```shell
sudo docker logs telegram_bot
```

or for database container:

```shell
sudo docker logs postgres-open-wishlist
```

By default bot is only logging error. To turn on debug logs please run command below when starting bot

```shell
sudo -E TOKEN=<your_bot_token> DEBUG=true docker compose up -d
```
