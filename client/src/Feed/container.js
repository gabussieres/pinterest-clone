import { connect } from "react-redux";
import { createStructuredSelector } from "reselect";

import Feed from "./component.js";
import feed from "./selectors.js";

import { fetchFeed } from "./actions";

const selectors = { feed };
const actions = { fetchFeed };

export default connect(createStructuredSelector(selectors), actions)(Feed);
