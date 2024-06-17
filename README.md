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
  <a href="#download">Download</a>
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
# You can create a config file to forego arguments
# Default location to create it: $HOME/.piper.yaml
$ echo "Hello, Discord!" | piper --token "your_secret_token" --channelID "your_channel_id"
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
    pipe, err := piper.NewPipe("your_secret_token", "your_channel_id")
    if err != nil {
        fmt.Printf("error creating pipe: %s", err)
    }

    err = pipe.Start()
    if err != nil {
        fmt.Printf("error starting piper: %s", err)
    }
    defer pipe.Stop()

    // Rest of your program...
    fmt.Println("Hello, World.") // This will be sent to discord!
}
```

## CLI Options

The Piper CLI supports the following options:

* `--token` - Discord bot token
* `--channelID` - Discord channel ID
* `--config` - Config file (default is $HOME/.piper.yaml)

## Config File Settings

You can configure Piper using a config file. The default config file location is `~/.piper/config.yaml`. The following settings are available:

* `token` - Discord bot token
* `channel` - Discord channel ID

Example config file:

```yaml
token: your_secret_token
channel: your_channel_id
```

## Download

You can download the latest version of Piper from the [releases page](https://github.com/AndrewVota/piper/releases).

## License

MIT

---

> Website [Vota.cc](https://www.vota.cc) &nbsp;&middot;&nbsp;
> GitHub [@AndrewVota](https://github.com/AndrewVota) &nbsp;&middot;&nbsp;
> LinkedIn [@AndrewVota](https://linkedin/in/AndrewVota) &nbsp;&middot;
