import { combineReducers } from "redux";
import { routerReducer as routing } from "react-router-redux";

import { feed } from "../Feed/reducers";
import { user } from "../User/reducers";
import { pinDetails } from "../PinDetails/reducers";

const rootReducer = combineReducers({
  feed,
  user,
  pinDetails,
  routing
});

export default rootReducer;
