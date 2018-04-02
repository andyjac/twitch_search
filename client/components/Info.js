import React, { Component } from 'react';
import { View, Text } from 'react-native';

class Info extends Component {
  constructor(props, context) {
    super(props, context);
  }

  render() {
    return (
      <View style={{width: 350}}>
        <Text style={{fontSize: 24, textAlign: 'left'}}>
          <Text>Hello and welcome to</Text>
          <Text style={{fontWeight: 'bold'}}>{' Twitch Search'}</Text>
          <Text>{'!'}{'\n'}{'\n'}</Text>
          <Text>
            Enter a comma separated list of
            Twitch streamers below and click
          </Text>
          <Text style={{fontWeight: 'bold'}}>{' Search'}</Text>
          <Text>.</Text>
        </Text>
      </View>
    );
  }
}

export default Info;
