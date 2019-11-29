export const FETCH_USER = "FETCH_USER";
export const FETCHING_USER = "FETCHING_USER";
export const FETCH_USER_SUCCESS = "FETCH_USER_SUCCESS";
export const FETCH_USER_FAILURE = "FETCH_USER_FAILURE";

export const FETCH_USER_PINS = "FETCH_USER_PINS";
export const FETCHING_USER_PINS = "FETCHING_USER_PINS";
export const FETCH_USER_PINS_SUCCESS = "FETCH_USER_PINS_SUCCESS";
export const FETCH_USER_PINS_FAILURE = "FETCH_USER_PINS_FAILURE";

export function fetchUser(user_id = "") {
  return { type: FETCH_USER, user_id };
}

export function fetchingUser() {
  return { type: FETCHING_USER };
}

export function fetchUserSuccess(results) {
  return {
    type: FETCH_USER_SUCCESS,
    results
  };
}

export function fetchUserFailure(message) {
  return {
    type: FETCH_USER_FAILURE,
    message
  };
}

export function fetchUserPins(user_id = "") {
  return { type: FETCH_USER_PINS, user_id };
}

export function fetchingUserPins() {
  return { type: FETCHING_USER_PINS };
}

export function fetchUserPinsSuccess(results) {
  return {
    type: FETCH_USER_PINS_SUCCESS,
    results
  };
}

export function fetchUserPinsFailure(message) {
  return {
    type: FETCH_USER_PINS_FAILURE,
    message
  };
}
