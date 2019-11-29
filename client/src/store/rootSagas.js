import { fork, all } from "redux-saga/effects";

import { watchFetchFeed } from "../Feed/sagas";
import {
  watchFetchUserProfile,
  watchFetchUserProfilePins
} from "../UserProfile/sagas";
import { watchFetchPinDetails, watchDeletePin } from "../PinDetails/sagas";
import { watchFetchUser, watchFetchUserPins } from "../Navigation/sagas";

export default function* root() {
  yield all([
    fork(watchFetchFeed),
    fork(watchFetchUserProfile),
    fork(watchFetchUserPins),
    fork(watchFetchPinDetails),
    fork(watchDeletePin),
    fork(watchFetchUser),
    fork(watchFetchUserProfilePins)
  ]);
}
