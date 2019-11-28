import { createStructuredSelector } from "reselect";

const feedSelector = state => state.feed;

export default createStructuredSelector({
  feed: feedSelector
});
