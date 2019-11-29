export const FETCH_USER_PROFILE = "FETCH_USER_PROFILE";
export const FETCHING_USER_PROFILE = "FETCHING_USER_PROFILE";
export const FETCH_USER_PROFILE_SUCCESS = "FETCH_USER_PROFILE_SUCCESS";
export const FETCH_USER_PROFILE_FAILURE = "FETCH_USER_PROFILE_FAILURE";

export const FETCH_USER_PROFILE_PINS = "FETCH_USER_PROFILE_PINS";
export const FETCHING_USER_PROFILE_PINS = "FETCHING_USER_PROFILE_PINS";
export const FETCH_USER_PROFILE_PINS_SUCCESS =
  "FETCH_USER_PROFILE_PINS_SUCCESS";
export const FETCH_USER_PROFILE_PINS_FAILURE =
  "FETCH_USER_PROFILE_PINS_FAILURE";

export function fetchUserProfile(user_id = "") {
  return { type: FETCH_USER_PROFILE, user_id };
}

export function fetchingUserProfile() {
  return { type: FETCHING_USER_PROFILE };
}

export function fetchUserProfileSuccess(results) {
  return {
    type: FETCH_USER_PROFILE_SUCCESS,
    results
  };
}

export function fetchUserProfileFailure(message) {
  return {
    type: FETCH_USER_PROFILE_FAILURE,
    message
  };
}

export function fetchUserProfilePins(user_id = "") {
  return { type: FETCH_USER_PROFILE_PINS, user_id };
}

export function fetchingUserProfilePins() {
  return { type: FETCHING_USER_PROFILE_PINS };
}

export function fetchUserProfilePinsSuccess(results) {
  return {
    type: FETCH_USER_PROFILE_PINS_SUCCESS,
    results
  };
}

export function fetchUserProfilePinsFailure(message) {
  return {
    type: FETCH_USER_PROFILE_PINS_FAILURE,
    message
  };
}
