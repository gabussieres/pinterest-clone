import * as ActionTypes from "./actions";
import { FetchStatus } from "../constants/entityStatus";

const InitialState = {
  details: {},
  pins: [],
  loggedIn: false,
  userError: null,
  pinsError: null,
  userFetchStatus: FetchStatus.none,
  pinsFetchStatus: FetchStatus.none
};

export function user(state = InitialState, action) {
  switch (action.type) {
    case ActionTypes.FETCHING_USER: {
      return {
        ...state,
        details: {},
        userFetchStatus: FetchStatus.loading,
        userError: null
      };
    }
    case ActionTypes.FETCH_USER_SUCCESS: {
      return {
        ...state,
        details: action.results,
        loggedIn: true,
        userFetchStatus: FetchStatus.loaded,
        userError: null
      };
    }
    case ActionTypes.FETCH_USER_FAILURE: {
      return {
        ...state,
        details: {},
        userFetchStatus: FetchStatus.failed,
        userError: action.message
      };
    }
    case ActionTypes.FETCHING_USER_PINS: {
      return {
        ...state,
        pins: [],
        pinsFetchStatus: FetchStatus.loading,
        pinsError: null
      };
    }
    case ActionTypes.FETCH_USER_PINS_SUCCESS: {
      return {
        ...state,
        pins: action.results,
        pinsFetchStatus: FetchStatus.loaded,
        pinsError: null
      };
    }
    case ActionTypes.FETCH_USER_PINS_FAILURE: {
      return {
        ...state,
        pins: [],
        pinsFetchStatus: FetchStatus.failed,
        pinsError: action.message
      };
    }
    default: {
      return state;
    }
  }
}
