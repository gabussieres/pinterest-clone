import { combineReducers } from "redux";
import { routerReducer as routing } from "react-router-redux";
import { persistReducer } from "redux-persist";
import storage from "redux-persist/lib/storage";
import autoMergeLevel2 from "redux-persist/lib/stateReconciler/autoMergeLevel2";

import { user } from "../Navigation/reducers";
import { feed } from "../Feed/reducers";
import { userProfile } from "../UserProfile/reducers";
import { pinDetails } from "../PinDetails/reducers";

const persistConfig = {
  key: "root",
  storage: storage,
  stateReconciler: autoMergeLevel2 // see "Merge Process" section for details.
};

const rootReducer = combineReducers({
  user,
  feed,
  userProfile,
  pinDetails,
  routing
});

const pReducer = persistReducer(persistConfig, rootReducer);

export default pReducer;
