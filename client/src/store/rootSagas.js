import { fork } from "redux-saga/effects";

import { watchFetchFeed } from "../Feed/sagas";
import { watchFetchUser, watchFetchUserPins } from "../User/sagas";
import { watchFetchPinDetails, watchDeletePin } from "../PinDetails/sagas";

export default function* root() {
  yield [
    fork(watchFetchFeed),
    fork(watchFetchUser),
    fork(watchFetchUserPins),
    fork(watchFetchPinDetails),
    fork(watchDeletePin)
  ];
}
