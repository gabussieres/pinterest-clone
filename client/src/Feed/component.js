import React, { Component } from "react";
import { PinList } from "../PinList/component";

class Feed extends Component {
  constructor(props) {
    super(props);
    props.fetchFeed(15);
  }

  render() {
    const { pins, error, fetchStatus } = this.props.feed.feed;
    if (error != null) {
      return null;
    }
    return <div>{PinList(pins, fetchStatus)}</div>;
  }
}

export default Feed;
