# eve-ng-cli-manager

[![Go Report Card](https://goreportcard.com/badge/github.com/inexio/eve-ng-cli-manger)](https://goreportcard.com/report/github.com/inexio/eve-ng-cli-manger)
[![GitHub license](https://img.shields.io/badge/license-BSD-blue.svg)](https://github.com/inexio/eve-ng-cli-manger/blob/master/LICENSE)
[![GitHub code style](https://img.shields.io/badge/code%20style-uber--go-brightgreen)](https://github.com/uber-go/guide/blob/master/style.md)

## Description

A tool for using the [Eve-NG](https://www.eve-ng.net/)  [REST API](https://www.eve-ng.net/index.php/documentation/howtos/how-to-eve-ng-api/) via the command-line written in go using the cobra CLI framework.

The tool uses our [eve-ng-restapi-go-client](https://github.com/inexio/eve-ng-restapi-go-client).

## Code Style

This project was written according to the **[uber-go](https://github.com/uber-go/guide/blob/master/style.md)** coding style.

## Features

### Key Eve-NG Feature Support

This client allows you to create/delete:

- Labs
- Folders
- Networks
- Nodes
- Users

and also enables you to:

- Move Labs and Folders
- Edit Labs and Users
- Get SystemStatus values
- Start/Stop/Export/Wipe Nodes

## Requirements

Requires a running instance of [Eve-NG](https://www.eve-ng.net/).

To check if your setup works, follow the steps provided in the **'Tests'** section of this document.

## Installation

```
go get github.com/inexio/eve-ng-cli-manager
```

or

```
git clone https://github.com/inexio/eve-ng-cli-manager.git
```

## Setup

After downloading you have to decide wether you use the tool via source or if you want to use its binary.

For using the tool via source you can run:

```
go run eve-ng-cli-manager/eve-ng/main.go  
```

or you can compile it to a binary:

```
cd eve-ng-cli-manager/eve-ng
go install
cd $GOBIN
eve-ng
```

After installing the tool you have to either declare a path to your config file or set certain environment variables for the tool to work.

These can be set as follows:

#### Config File

##### Using source

In **cmd/root.go**, in the **initConfig()** function you can see the following lines of code:

```
viper.AddConfigPath("config/")
viper.SetConfigType("yaml")
viper.SetConfigName("eve-cli-config")
```

ConfigPath is relative to the package location.
ConfigType and name can also be changed to match your desired type of config.

##### Using the binary

If you are using the binary you have to add a folder named **config** in the same directory the binary is located at.

In this **config** folder you have to add a file named **eve-cli-config.yaml**. This file has to contain all data required in the eve-cli config.

##### Using the '--config' flag

If you want to use an alternative config file for only one command, you can use the --config flag with the command. 

Here's an example of this:

```
eve-ng get folders "/" --config "~/go/src/eve-ng-cli-manager/eve-ng/config/configtwo.yaml"
```

Note, that this only works for one command and doesn't permanently change the config file.

##### Using the env-var EVE_CLI_CONFIG

You can also set an environment variable to read in the config file given in that variable.

To set this variable use:

```
export <YOUR_ENV_PREFIX>_CONFIG=/absolut/path/to/your/config
```

This way the config will always be read in instead of how '--config' has to be set every time you want to execute a command.

#### Environment Variables

Also in the **root.go** file, in the **initConfig()** function you will find:

```
//Set env var prefix to only match certain vars
viper.SetEnvPrefix("EVE_CLI")
```

`SetEnvPrefix` can be changed to whatever prefix you prefer to have in your environment vars.

The needed environment variables can then be added as follows:

For the management endpoint:

```
export EVE_CLI_BASEURL="<your mgmt baseUrl>"
export EVE_CLI_USERNAME="<your username>"
export EVE_CLI_PASSWORD="<your password>"
```

optional env-vars:

```
export EVE_CLI_LABNAME="<name of your desired standard lab>"
export EVE_CLI_LABPATH="path to your desired standard lab"
```

## Usage

The following section will show you how to create a lab and do various operations with it.

```go
eve-ng create lab --path "/" --name "TestLab" --author "Admin" --version 1 --description "A Test Lab" --body "A longer version of your description"
```

Should return:

```
Successfully created lab.
```

You can get information about a component via the get command:

```
eve-ng get lab --lab "/TestLab"
```

Which should return:

```
Lab
  Id: a486aa2d-6036-4f3a-9162-82050d80e1ec
  Name: TestLabab
  Version: 1
  Author: admin
  Body: A longer version of your description
  Description: A Test Lab
  Filename: TestLabab.unl
```

To add a network use the create command:

```
eve-ng create network --name "TestNet" --left 10 --top 10 --type nat0 --lab "/TestLab"
```

Which should return:

```
Successfully added network to lab.
ID: 1
```

To retreive data about the network you just created use the get command once again:

```
eve-ng get network 1 --lab "/TestLab"
```

 Which will return:

```
Network
  Name: TestNet
  Type: nat0
  Top: 10
  Left: 10
  Style: Solid
  Linkstyle: Straight
  Color:
  Label:
  Visibility: 1
```

Now you can create a node using:

```
eve-ng create node --name TestNode --ram 512 --top 100 --left 100 --icon Router.png --image "asav-952-204" --template asav --type qemu --lab "/TestLab"
```

Which returns:

```
Successfully added node to lab.
ID: 1
```

You can now once again use the get command to retreive information about the node:

```
eve-ng get node 1 --lab "/TestLab"
```

This will return:

```
Node
  Uuid: 93af4a9d-a7ac-4514-879d-a36738f52a0d
  Name: TestNode
  Type: qemu
  Status: 0
  Template: asav
  Cpu: 1
  Ram: 512
  Image: asav-952-204
  Console: telnet
  Ethernet: 4
  Delay: 0
  Icon: AristaSW.png
  Url: html5/#/client/unknowntoken
  Top: 100
  Left: 100
  Config: 0
  Firstmac: 50:04:00:01:00:00
  Configlist:
```



## Tests

In order to test if your setup is operational you can use the following command:

```
go run main.go get lab-files "/"
```

or

```
eve-ng get lab-files "/"
```

The output should be a list of all laboratory files located in the root location:

```
LabFiles(0)
    /
```

or if you have any labs configured it should look like this:

```
LabFiles(2) 

  LabFile
    File: TestLab.unl
    Path: ///TestLab.unl

  LabFile
    File: TestLab2.unl
    Path: ///TestLab2.unl
```

## Getting Help

If there are any problems or something does not work as intended, open an [issue](https://github.com/inexio/eve-ng-cli-manager/issues/new/choose) on GitHub.

If you have problems with the usage of the application itself you can use the built in **--help** flag to get some useful information about every command.

## Contribution

Contributions to the project are welcome.

We are looking forward to your bug reports, suggestions and fixes.

If you want to make any contributions make sure your go report does match up with our projects score of **A+**.

When you contribute make sure your code is conform to the **uber-go** coding style.

Happy Coding!
