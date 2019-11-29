import { createStructuredSelector } from "reselect";

const userSelector = state => state.user;
const pinSelector = state => state.pinDetails;

export default createStructuredSelector({
  user: userSelector,
  pin: pinSelector
});
