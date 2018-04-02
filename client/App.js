import React, { Component } from 'react';
import { StackNavigator } from 'react-navigation';

import Search from './components/Search';
import StreamList from './components/StreamList';

const RootStack = StackNavigator(
  {
    Home: {
      screen: Search
    },
    Streams: {
      screen: StreamList
    }
  },
  {
    initialRouteName: 'Home',
    navigationOptions: {
      headerStyle: {
        backgroundColor: '#6441A4'
      },
      headerTintColor: '#fff',
      headerTitleStyle: {
        fontWeight: 'bold'
      }
    }
  }
);

class App extends Component {
  render() {
    return (
      <RootStack/>
    );
  }
}

export default App;
