import React, { Component } from 'react';
import { View, Text, TextInput, Alert, StyleSheet } from 'react-native';
import { Button } from 'react-native-elements';
import api from '../services/api';

import Info from './Info';

class Search extends Component {
  static navigationOptions = {
    headerTitle: 'Home'
  };

  constructor(props, context) {
    super(props, context);

    this.state = {
      text: '',
      loading: false
    };

    this.handleSearch = this.handleSearch.bind(this);
  }

  async handleSearch() {
    this.setState({ loading: true });

    let { text } = this.state;
    let names = text.replace(/\s+/g, '').toLowerCase().split(',');

    try {
      let streams = await api.getLiveStreamsByName(names);
      this.setState({ loading: false });

      if (streams.error) {
        Alert.alert(streams.message);
      } else {
        const { navigate } = this.props.navigation;

        if (!streams.length) {
          Alert.alert('No live stream(s) found');
        } else {
          navigate('Streams', { streams });
        }
      }
    } catch(error) {
      this.setState({ loading: false });
      Alert.alert('Oops! It looks like we encountered an error. Please try again.');
    }
  }

  render() {
    return (
      <View
        style={{
          flex: 1,
          justifyContent: 'space-around',
          alignItems: 'center'
        }}>
        <Info/>
        <View style={{ width: 300 }}>
          <TextInput
            onChangeText={(text) => this.setState({ text })}
            placeholder='Comma separated Twitch users...'
            style={{
              borderColor: '#6441A4',
              borderWidth: 2,
              height: 50,
              textAlign: 'center'
            }} />
            <Button
              loading={this.state.loading}
              onPress={this.handleSearch}
              title='Search'
              buttonStyle={{
                backgroundColor: '#6441A4',
                marginTop: 20
              }} />
        </View>
      </View>
    );
  }
}

export default Search;
