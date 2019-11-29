import { connect } from "react-redux";
import { createStructuredSelector } from "reselect";

import UserProfile from "./component.js";
import user from "./selectors.js";

import { fetchUserProfile, fetchUserProfilePins } from "./actions";

const selectors = { user };
const actions = { fetchUserProfile, fetchUserProfilePins };

export default connect(
  createStructuredSelector(selectors),
  actions
)(UserProfile);
