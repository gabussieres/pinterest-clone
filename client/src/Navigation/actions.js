export const FETCH_USER = "FETCH_USER";
export const FETCHING_USER = "FETCHING_USER";
export const FETCH_USER_SUCCESS = "FETCH_USER_SUCCESS";
export const FETCH_USER_FAILURE = "FETCH_USER_FAILURE";

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
