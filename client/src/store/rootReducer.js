import { combineReducers } from "redux";
import { routerReducer as routing } from "react-router-redux";

import { user } from "../Navigation/reducers";
import { feed } from "../Feed/reducers";
import { userProfile } from "../UserProfile/reducers";
import { pinDetails } from "../PinDetails/reducers";

const rootReducer = combineReducers({
  user,
  feed,
  userProfile,
  pinDetails,
  routing
});

export default rootReducer;
