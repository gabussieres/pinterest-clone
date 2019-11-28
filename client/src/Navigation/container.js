import { connect } from "react-redux";
import { createStructuredSelector } from "reselect";

import Navigation from "./component.js";
import user from "./selectors.js";

import { fetchUser } from "./actions";

const selectors = { user };
const actions = { fetchUser };

export default connect(
  createStructuredSelector(selectors),
  actions
)(Navigation);
