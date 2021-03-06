import { takeEvery, put } from "redux-saga/effects";
import Api from "../utils/Api";

import * as actions from "./actions";

export function* fetchUserProfile(action) {
  try {
    yield put(actions.fetchingUserProfile());
    const result = yield Api.fetchUser(action.user_id);

    yield put(actions.fetchUserProfileSuccess(result));
  } catch (e) {
    yield put(actions.fetchUserProfileFailure(e));
  }
}

export function* watchFetchUserProfile() {
  yield takeEvery(actions.FETCH_USER_PROFILE, fetchUserProfile);
}

export function* fetchUserProfilePins(action) {
  try {
    yield put(actions.fetchingUserProfilePins());
    const result = yield Api.fetchUserPins(action.user_id);

    yield put(actions.fetchUserProfilePinsSuccess(result));
  } catch (e) {
    yield put(actions.fetchUserProfilePinsFailure(e));
  }
}

export function* watchFetchUserProfilePins() {
  yield takeEvery(actions.FETCH_USER_PROFILE_PINS, fetchUserProfilePins);
}
