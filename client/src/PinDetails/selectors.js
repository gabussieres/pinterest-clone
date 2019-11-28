import { createStructuredSelector } from "reselect";

const pinSelector = state => state.pinDetails;

export default createStructuredSelector({
  pin: pinSelector
});
