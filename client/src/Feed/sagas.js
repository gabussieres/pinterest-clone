import { takeEvery, put } from "redux-saga/effects";
import Api from "../utils/Api";

import * as actions from "./actions";

export function* fetchFeed(action) {
  try {
    yield put(actions.fetchingFeed());
    const result = yield Api.fetchFeed(action.pin_amount);

    yield put(actions.fetchFeedSuccess(result));
  } catch (e) {
    yield put(actions.fetchFeedFailure(e));
  }
}

export function* watchFetchFeed() {
  yield takeEvery(actions.FETCH_FEED, fetchFeed);
}
