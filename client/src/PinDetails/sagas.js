import { takeEvery, put } from "redux-saga/effects";
import Api from "../utils/Api";

import * as actions from "./actions";

export function* fetchPinDetails(action) {
  try {
    yield put(actions.fetchingPinDetails());
    const result = yield Api.fetchPinDetails(action.pin_id);
    console.log(result);

    yield put(actions.fetchPinDetailsSuccess(result));
  } catch (e) {
    yield put(actions.fetchPinDetailsFailure(e));
  }
}

export function* watchFetchPinDetails() {
  yield takeEvery(actions.FETCH_PIN_DETAILS, fetchPinDetails);
}

export function* deletePin(action) {
  try {
    yield put(actions.deletingPin());
    const result = yield Api.deletePin(action.user_id, action.pin_id);
    console.log(result);

    yield put(actions.deletePinSuccess(result));
  } catch (e) {
    yield put(actions.deletePinFailure(e));
  }
}

export function* watchDeletePin() {
  yield takeEvery(actions.DELETE_PIN, deletePin);
}
