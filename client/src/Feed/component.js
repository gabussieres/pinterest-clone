import { Component } from "react";
import debounce from "lodash.debounce";

import { FetchStatus } from "../constants/entityStatus";
import { PinList } from "../PinList/component";

class Feed extends Component {
  constructor(props) {
    super(props);
    props.resetFeed();
    props.fetchFeed(20);

    // This infinite scroll implementation was guided by this blog post
    // https://alligator.io/react/react-infinite-scroll/
    window.onscroll = debounce(() => {
      const { error, fetchStatus } = this.props.feed.feed;
      if (error || fetchStatus === FetchStatus.failed) return;

      if (window.innerHeight + window.scrollY >= document.body.scrollHeight) {
        props.fetchFeed(20);
      }
    }, 100);
  }

  render() {
    const { pins, error, fetchStatus } = this.props.feed.feed;

    if (error != null) {
      return null;
    }

    return PinList(pins, fetchStatus);
  }
}

export default Feed;
