import { createStructuredSelector } from "reselect";

const userSelector = state => state.user;
const pinsSelector = state => state.feed;
const userProfileSelector = state => state.userProfile;

export default createStructuredSelector({
  user: userSelector,
  pins: pinsSelector,
  userProfile: userProfileSelector
});
