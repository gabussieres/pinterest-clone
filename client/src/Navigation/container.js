import { connect } from "react-redux";
import { createStructuredSelector } from "reselect";

import Navigation from "./component.js";
import user from "./selectors.js";

import { fetchUser, fetchUserPins } from "./actions";

const selectors = { user };
const actions = { fetchUser, fetchUserPins };

export default connect(
  createStructuredSelector(selectors),
  actions
)(Navigation);
