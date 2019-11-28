export const FETCH_FEED = "FETCH_FEED";
export const FETCHING_FEED = "FETCHING_FEED";
export const FETCH_FEED_SUCCESS = "FETCH_FEED_SUCCESS";
export const FETCH_FEED_FAILURE = "FETCH_FEED_FAILURE";

export const FETCH_USER_PINS = "FETCH_USER_PINS";
export const FETCHING_USER_PINS = "FETCHING_USER_PINS";
export const FETCH_USER_PINS_SUCCESS = "FETCH_USER_PINS_SUCCESS";
export const FETCH_USER_PINS_FAILURE = "FETCH_USER_PINS_FAILURE";

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
