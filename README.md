# Shenas

**Shenas** is a telegram bot in Persian that lets a person set some questions and ask their friends to answer those questions in order to how much they know him/her.

## Requirements

- [Git](https://git-scm.com)

- [Go](https://go.dev)

- [MySQL](https://www.mysql.com)

- [Go-MySQL-Driver](https://github.com/go-sql-driver/mysql)

- [Telegram-Bot-API](https://github.com/go-telegram-bot-api/telegram-bot-api/v5)

## How to use

Before anything, you should clone the shenas repository into your computer:

````
git clone https://github.com/barahouei/shenas.git
````

When the code is on your computer, you need to first import `shenas` database from `database` folder into MySQL and then set the following environment variables in your machine:

| Environment Variable      | Description                                    |
| :------------------------ | :--------------------------------------------- |
| `TELEGRAM_APITOKEN`       | Bot API token that you got from BotFather      |
| `BOT_USERNAME`            | Bot Username                                   |
| `DB_NAME`                 | Database Name - `shenas`                       |
| `DB_USERNAME`             | MySQL Username                                 |
| `DB_PASSWORD`             | MySQL Password                                 |

After that you can go to `shenas` folder and simply run the program with the following command:

````
go run .
````

Or build an executive file to run with the following command:

````
go build
````

## Licence

[MIT](https://choosealicense.com/licenses/mit/)