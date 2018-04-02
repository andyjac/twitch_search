# Twitch Search

Hello and welcome to __Twitch Search__.

This app allows you to see if your favorite Twitch streamers are online and jump directly to their live stream. Just follow the instructions below to setup the app on your local machine.

Thanks!

## Setup Instructions

### Dependencies

The following pieces of software are required to run this project:

- [yarn](https://yarnpkg.com/en/docs/install) - `nodejs` dependency manager
- [glide](https://yarnpkg.com/en/docs/install) - `golang` dependency manager
- [create-react-native-app](https://github.com/react-community/create-react-native-app) - `react-native` boilerplace
- [expo](https://expo.io/learn) - Run `react-native` apps on your personal device (__optional__)

### Install

From the root of the project, run `make install` in order to install client and server packages.

That's it!

## Running the App

### Server

To start the server, run `make server` from the root directory.

### Client

In a separate terminal, run `make client` to initiate the react-native app. Then follow the on-screen instructions for viewing the app on your device.

> NOTE: If you want to run the client on your personal device with `expo` you *may* need to change the api url that the client points to from `127.0.0.1` to your machine's private ip. If this is the case, open the file `client/config.js` and edit the `api.host` property to point to your private ip. Then restart the client app by running `make client` from the root directory.
