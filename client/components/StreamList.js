import React, { Component } from 'react';
import { View, FlatList, WebView, Linking } from 'react-native';
import { List, ListItem, Text, Icon } from 'react-native-elements';

class StreamList extends Component {
  static navigationOptions = {
      headerTitle: 'Streams'
  };

  constructor(props, context) {
    super(props, context);

    this.state = { streamUrl: '' };
    this.renderStream = this.renderStream.bind(this);
  }

  renderStream(data) {
    const { item: stream } = data;

    return (
      <ListItem
        roundAvatar
        onPress={() => this.setState({ streamUrl: stream.channel.url })}
        avatar={{uri: stream.channel.logo}}
        key={stream._id}
        title={
            <Text style={{marginLeft: 10}}>
                <Text style={{fontWeight: 'bold'}}>{stream.channel.display_name}</Text>
                  <Text>{' - '}</Text>
                    <Text style={{fontWeight: 'bold'}}>{stream.game}</Text>
              </Text>
            }
            subtitle={stream.channel.status}
            />
    );
  }

  render() {
    const { state } = this.props.navigation;

    if (!state.params || !state.params.streams) {
      return (
        <View style={{
                flex: 1,
                justifyContent: 'center',
                alignItems: 'center'
              }}>
          <Text style={{fontSize: 24}}>No streams found.</Text>
        </View>
      );
    }

    const { streams } = state.params;
    const { streamUrl } = this.state;

    // open external native WebView to stream
    if  (streamUrl) {
      return (
        <WebView
          ref={(ref) => this.webView = ref}
          source={{ uri: streamUrl }}
          onNavigationStateChange={(event) => {
            if (event.url !== streamUrl) {
              this.webView.stopLoading();
              Linking.openURL(event.url);

              // null out streamUrl so we can re-render the StreamList
              setTimeout(() => {
                this.setState({ streamUrl: null });
              }, 50);
          }}}
          />
      );
    }
    return (
      <View>
        <List>
          <FlatList
            data={streams}
            renderItem={this.renderStream}
            keyExtractor={(item, i) => i}
            />
        </List>
      </View>
    );
  }
}

export default StreamList;
