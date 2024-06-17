<h1 align="center">
  <br>
  <a href="https://github.com/AndrewVota/piper"><img src="https://via.placeholder.com/200" alt="Piper" width="200"></a>
  <br>
  Piper
  <br>
</h1>

<h4 align="center">A tool to pipe std.in or std.out to Discord, available as a CLI application or a Golang package.</h4>

<p align="center">
  <a href="#key-features">Key Features</a> •
  <a href="#how-to-use">How To Use</a> •
  <a href="#cli-options">CLI Options</a> •
  <a href="#config-file-settings">Config File Settings</a> •
  <a href="#golang-library-options">Golang Library Options</a> •
  <a href="#download">Download</a> •
</p>

![screenshot](https://via.placeholder.com/800x400)

## Key Features

* Pipe std.in or std.out to Discord
* CLI application and Golang package support
* Easy installation
* Cross platform

## How To Use

### CLI Application

To use Piper as a CLI application, install it using the following command:

```bash
curl -sSfL https://raw.githubusercontent.com/AndrewVota/piper/main/install.sh | sh
```

Then, you can run the application from your terminal:

```bash
# Pipe std.in to Discord
$ echo "Hello, Discord!" | piper

# Pipe std.out from a command to Discord
$ some_command | piper
```

### Golang Package

To use Piper as a Golang package, install it using:

```bash
go get github.com/andrewvota/piper
```

Then, import and use it in your Golang code:

```go
package main

import (
    "github.com/andrewvota/piper"
)

func main() {
    piper.SendToDiscord("Hello, Discord!")
}
```

## CLI Options

The Piper CLI supports the following options:

* `-t, --token` - Discord bot token
* `-c, --channel` - Discord channel ID
* `-m, --message` - Message to send
* `-f, --file` - Path to file to send

## Config File Settings

You can configure Piper using a config file. The default config file location is `~/.piper/config.yaml`. The following settings are available:

* `token` - Discord bot token
* `channel` - Discord channel ID
* `message` - Default message
* `file` - Default file path

Example config file:

```yaml
token: your_discord_bot_token
channel: your_discord_channel_id
message: Default message
file: /path/to/default/file
```

## Golang Library Options

The Piper Golang package provides the following functions:

* `SendToDiscord(message string)` - Sends a message to Discord
* `SendFileToDiscord(filePath string)` - Sends a file to Discord

Example usage:

```go
package main

import (
    "github.com/andrewvota/piper"
)

func main() {
    piper.SendToDiscord("Hello, Discord!")
    piper.SendFileToDiscord("/path/to/file.txt")
}
```

## Download

You can download the latest version of Piper from the [releases page](https://github.com/AndrewVota/piper/releases).

## License

MIT

---

> [vota.cc] (https://www.vota.cc) &nbsp;&middot;&nbsp;
> GitHub [@AndrewVota](https://github.com/AndrewVota) &nbsp;&middot;&nbsp;
> Twitter [@AndrewVota](https://twitter.com/AndrewVota)
