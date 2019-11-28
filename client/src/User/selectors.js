import { createStructuredSelector } from "reselect";

const userSelector = state => state.user;
const pinsSelector = state => state.feed;

export default createStructuredSelector({
  user: userSelector,
  pins: pinsSelector
});
