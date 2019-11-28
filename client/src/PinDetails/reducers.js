import * as ActionTypes from "./actions";
import { FetchStatus } from "../constants/entityStatus";

const InitialState = {
  details: {},
  error: null,
  fetchStatus: FetchStatus.none
};

export function pinDetails(state = InitialState, action) {
  switch (action.type) {
    case ActionTypes.FETCHING_PIN_DETAILS: {
      return {
        ...state,
        details: {},
        fetchStatus: FetchStatus.loading,
        error: null
      };
    }
    case ActionTypes.FETCH_PIN_DETAILS_SUCCESS: {
      return {
        ...state,
        details: action.results,
        fetchStatus: FetchStatus.loaded,
        error: null
      };
    }
    case ActionTypes.FETCH_PIN_DETAILS_FAILURE: {
      return {
        ...state,
        details: {},
        fetchStatus: FetchStatus.failed,
        error: action.message
      };
    }
    default: {
      return state;
    }
  }
}
