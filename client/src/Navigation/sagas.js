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
