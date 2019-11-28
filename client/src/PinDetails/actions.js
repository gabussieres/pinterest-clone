export const FETCH_PIN_DETAILS = "FETCH_PIN_DETAILS";
export const FETCHING_PIN_DETAILS = "FETCHING_PIN_DETAILS";
export const FETCH_PIN_DETAILS_SUCCESS = "FETCH_PIN_DETAILS_SUCCESS";
export const FETCH_PIN_DETAILS_FAILURE = "FETCH_PIN_DETAILS_FAILURE";

export const DELETE_PIN = "DELETE_PIN";
export const DELETING_PIN = "DELETING_PIN";
export const DELETE_PIN_SUCCESS = "DELETE_PIN_SUCCESS";
export const DELETE_PIN_FAILURE = "DELETE_PIN_FAILURE";

export function fetchPinDetails(pin_id = "") {
  return { type: FETCH_PIN_DETAILS, pin_id };
}

export function fetchingPinDetails() {
  return { type: FETCHING_PIN_DETAILS };
}

export function fetchPinDetailsSuccess(results) {
  return {
    type: FETCH_PIN_DETAILS_SUCCESS,
    results
  };
}

export function fetchPinDetailsFailure(message) {
  return {
    type: FETCH_PIN_DETAILS_FAILURE,
    message
  };
}

export function deletePin(user_id = "", pin_id = "") {
  return { type: DELETE_PIN, user_id, pin_id };
}

export function deletingPin() {
  return { type: DELETING_PIN };
}

export function deletePinSuccess(results) {
  return {
    type: DELETE_PIN_SUCCESS,
    results
  };
}

export function deletePinFailure(message) {
  return {
    type: DELETE_PIN_FAILURE,
    message
  };
}
