# Go-Discord-Message-Bot

Simple Discord message bot

## Behaviour

Bot reacts to certain text messages by giving a reply

## Installing

* You must have [Go v1.21](https://go.dev/doc/install) (or higher) installed on your system
* To enable this bot for your discord server you must have a bot template at [Discord Applications](https://discord.com/developers/applications)
* The bot template must have text and channel/message viewing permissions
* The bot must be authorized and be on your server

## Executing program

* Clone this repository to the location of your choosing
* Enter your Bot Token into *config.json* file
* Open your terminal
* Navigate to the saved location using ```cd folderName``` command, where *folderName* is the name of your path folder
* When in right location run:
```
go build
go run main.go
```

## Customization

You can add your own message queries by adding rules to the *messageHandler()* function in *bot.go* file
> For Example:
> ```
> if m.Content == "Desired message query" {
>    _, _ = s.ChannelMesageSend(m.ChannelID, "Desired answer")
> }
> ```

## Help

For more types of handlers or general information about Discord Bot behaviours, please refer to:
* https://github.com/bwmarrin/discordgo - Go Discord package
* https://discord.com/developers/docs/intro - Discord developer portal

Feel free to report any issues or suggest improvements

## Version History

* v.0.0.1:

    * Initial Release
