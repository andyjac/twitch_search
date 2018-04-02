import Config from '../config';

class Api {
  constructor() {
    const { url, port, version } = Config.api;

    this.baseUrl = `${url}:${port}/api/${version}`;
  }

  async getUsersByName(names) {
    // create comma separated list of usernames
    let query = names.join(',');
    let users = await fetch(`${this.baseUrl}/users?name=${query}`);
    let json = await users.json();

    return json;
  }

  async getLiveStreamByUserId(id) {
    let stream = await fetch(`${this.baseUrl}/users/${id}/stream`);
    let json = stream.json();

    return json;
  }

  async getLiveStreamsByName(names) {
    let streams = [];
    let users = await this.getUsersByName(names);

    // bail out if we get an error fetching users
    if (users.error) {
      return users;
    }
    // for each user, try to fetch their live stream
    for (let user of users.result) {
      let stream = await this.getLiveStreamByUserId(user._id);

      if (!stream.error) {
        streams = streams.concat(stream.result);
      }
    }

    return streams;
  }
}

export default new Api();
