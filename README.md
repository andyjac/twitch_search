# Streaming App

Hello and welcome to my very simple "streaming" app.

## Setup Instructions

### Dependencies

The following pieces of software are required to run this project:

- [mysql](https://dev.mysql.com/downloads/mysql/)
- [yarn](https://yarnpkg.com/en/docs/install)
- [glide](https://yarnpkg.com/en/docs/install)
- [create-react-native-app](https://github.com/react-community/create-react-native-app)

### Database

This project uses a MySQL database so you will want to make sure you have the [mysql server](https://dev.mysql.com/downloads/mysql/) installed and running on your local machine.

After you have it installed, run the following commands to create a database and user for the project:

```sh
mysql> create database streaming_app_dev;
mysql> create user 'go_user'@'localhost' identified by 'password';
mysql> grant all privileges on streaming_app_dev.* to 'go_user'@'localhost';
```

After your database and user are setup, run `make setup` from the root of the project to install packages as well as migrate and seed the database.

That's it!

## Running the App

In order to start the server, run `make server` from the root directory. Then, in a separate terminal, run `make client` to initiate the react-native app. Then follow the on-screen instructions for viewing the app on your device.
