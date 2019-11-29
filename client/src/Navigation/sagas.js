import { takeEvery, put } from "redux-saga/effects";
import Api from "../utils/Api";

import * as actions from "./actions";

export function* fetchUser(action) {
  try {
    yield put(actions.fetchingUser());
    const result = yield Api.fetchUser(action.user_id);

    yield put(actions.fetchUserSuccess(result));
  } catch (e) {
    yield put(actions.fetchUserFailure(e));
  }
}

export function* watchFetchUser() {
  yield takeEvery(actions.FETCH_USER, fetchUser);
}

export function* fetchUserPins(action) {
  try {
    yield put(actions.fetchingUserPins());
    const result = yield Api.fetchUserPins(action.user_id);

    yield put(actions.fetchUserPinsSuccess(result));
  } catch (e) {
    yield put(actions.fetchUserPinsFailure(e));
  }
}

export function* watchFetchUserPins() {
  yield takeEvery(actions.FETCH_USER, fetchUserPins);
}
