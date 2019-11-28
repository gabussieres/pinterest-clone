import { connect } from "react-redux";
import { createStructuredSelector } from "reselect";

import PinDetails from "./component.js";
import pinDetails from "./selectors.js";

import { fetchPinDetails, deletePin } from "./actions";

const selectors = { pinDetails };
const actions = { fetchPinDetails, deletePin };

export default connect(
  createStructuredSelector(selectors),
  actions
)(PinDetails);
