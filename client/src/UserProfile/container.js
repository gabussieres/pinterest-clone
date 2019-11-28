import { connect } from "react-redux";
import { createStructuredSelector } from "reselect";

import UserProfile from "./component.js";
import user from "./selectors.js";

import { fetchUserProfile, fetchUserPins } from "./actions";

const selectors = { user };
const actions = { fetchUserProfile, fetchUserPins };

export default connect(
  createStructuredSelector(selectors),
  actions
)(UserProfile);
