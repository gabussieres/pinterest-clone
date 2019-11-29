export const RESET_FEED = "RESET_FEED";

export const FETCH_FEED = "FETCH_FEED";
export const FETCHING_FEED = "FETCHING_FEED";
export const FETCH_FEED_SUCCESS = "FETCH_FEED_SUCCESS";
export const FETCH_FEED_FAILURE = "FETCH_FEED_FAILURE";

export const FETCH_USER_PROFILE_PINS = "FETCH_USER_PROFILE_PINS";
export const FETCHING_USER_PROFILE_PINS = "FETCHING_USER_PROFILE_PINS";
export const FETCH_USER_PROFILE_PINS_SUCCESS =
  "FETCH_USER_PROFILE_PINS_SUCCESS";
export const FETCH_USER_PROFILE_PINS_FAILURE =
  "FETCH_USER_PROFILE_PINS_FAILURE";

export function resetFeed() {
  return { type: RESET_FEED };
}

export function fetchFeed(pin_amount = 0) {
  return { type: FETCH_FEED, pin_amount };
}

export function fetchingFeed() {
  return { type: FETCHING_FEED };
}

export function fetchFeedSuccess(results) {
  return {
    type: FETCH_FEED_SUCCESS,
    results
  };
}

export function fetchFeedFailure(message) {
  return {
    type: FETCH_FEED_FAILURE,
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
