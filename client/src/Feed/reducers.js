import * as ActionTypes from "./actions";
import { FetchStatus } from "../constants/entityStatus";

const InitialState = {
  pins: [],
  error: null,
  fetchStatus: FetchStatus.none
};

export function feed(state = InitialState, action) {
  switch (action.type) {
    case ActionTypes.FETCHING_FEED: {
      return {
        ...state,
        fetchStatus: FetchStatus.loading,
        error: null
      };
    }
    case ActionTypes.FETCH_FEED_SUCCESS: {
      return {
        ...state,
        pins: state.pins.concat(action.results),
        fetchStatus: FetchStatus.loaded,
        error: null
      };
    }
    case ActionTypes.FETCH_FEED_FAILURE: {
      return {
        ...state,
        fetchStatus: FetchStatus.failed,
        error: action.message
      };
    }
    case ActionTypes.FETCHING_USER_PROFILE_PINS: {
      return {
        ...state,
        pins: [],
        fetchStatus: FetchStatus.loading,
        error: null
      };
    }
    case ActionTypes.FETCH_USER_PROFILE_PINS_SUCCESS: {
      return {
        ...state,
        pins: action.results,
        fetchStatus: FetchStatus.loaded,
        error: null
      };
    }
    case ActionTypes.FETCH_USER_PROFILE_PINS_FAILURE: {
      return {
        ...state,
        pins: [],
        fetchStatus: FetchStatus.failed,
        error: action.message
      };
    }
    case ActionTypes.RESET_FEED: {
      return {
        ...state,
        pins: [],
        fetchStatus: FetchStatus.none,
        error: null
      };
    }
    default: {
      return state;
    }
  }
}
