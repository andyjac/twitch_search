# Twitch Search

Hello and welcome to __Twitch Search__.

This app allows you to see if your favorite Twitch streamers are online and jump directly to their live stream. Just follow the instructions below to setup the app on your local machine.

Thanks!

## Setup Instructions

### Dependencies

The following pieces of software are required to run this project:

- [yarn](https://yarnpkg.com/en/docs/install) - `nodejs` dependency manager
- [glide](https://glide.sh/) - `golang` dependency manager (__optional__ - see below)
- [create-react-native-app](https://github.com/react-community/create-react-native-app) - `react-native` boilerplace
- [expo](https://expo.io/learn) - Run `react-native` apps on your personal device (__optional__)

### Install

Clone the repo with the following command:

```sh
$ git clone git@github.com:andyjac/twitch_search.git $GOPATH/src/twitch_search
```

If you don't wish to install `glide` you can alternatively run the following commands to install the required `golang` packages:

```sh
$ go get -u github.com/gin-gonic/gin
$ go get -u github.com/spf13/viper
```

Then, from the root of the project, run:

```sh
$ make install
```

That's it!

## Running the App

### Server

To start the server, run the following command from the root directory:

```
$ make server
```

### Client

In a separate terminal, run the following command:

```sh
$ make client
```

Then follow the onscreen instructions to run the app on and emulator or your personal device.

> NOTE: If you want to run the client on your personal device with `expo` you *may* need to change the api url that the client points to from `127.0.0.1` to your machine's private ip. If this is the case, open the file `client/config.js` and edit the `api.host` property to point to your private ip. Then restart the client app by running `make client` from the root directory.
