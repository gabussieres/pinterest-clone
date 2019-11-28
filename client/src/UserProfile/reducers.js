import * as ActionTypes from "./actions";
import { FetchStatus } from "../constants/entityStatus";

const InitialState = {
  user: {},
  error: null,
  fetchStatus: FetchStatus.none
};

export function userProfile(state = InitialState, action) {
  switch (action.type) {
    case ActionTypes.FETCHING_USER_PROFILE: {
      return {
        ...state,
        user: {},
        fetchStatus: FetchStatus.loading,
        error: null
      };
    }
    case ActionTypes.FETCH_USER_PROFILE_SUCCESS: {
      return {
        ...state,
        user: action.results,
        fetchStatus: FetchStatus.loaded,
        error: null
      };
    }
    case ActionTypes.FETCH_USER_PROFILE_FAILURE: {
      return {
        ...state,
        user: {},
        fetchStatus: FetchStatus.failed,
        error: action.message
      };
    }
    default: {
      return state;
    }
  }
}
