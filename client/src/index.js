import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  useParams
} from "react-router-dom";
import styled from "styled-components";
import { PersistGate } from "redux-persist/lib/integration/react";
import { persistStore } from "redux-persist";

import configureStore from "./store/configureStore.js";
import Feed from "./pages/Feed";
import PinDetails from "./pages/PinDetails";
import UserProfile from "./pages/UserProfile";

const store = configureStore();
const persistor = persistStore(store);

function UserRoute() {
  let { user_id } = useParams();
  return UserProfile(user_id);
}

function PinDetailsRoute() {
  let { pin_id } = useParams();
  return PinDetails(pin_id);
}

const AppWrapper = styled.div`
  background-color: #fafafa;
  position: "absolute";
  height: "100% !important";
  width: "100% !important";
`;

function App() {
  return (
    <AppWrapper>
      <Switch>
        <Route path="/pin/:pin_id">
          <PinDetailsRoute />
        </Route>
        <Route path="/:user_id">
          <UserRoute />
        </Route>
        <Route path="/">
          <Feed />
        </Route>
      </Switch>
    </AppWrapper>
  );
}

ReactDOM.render(
  <Provider store={store}>
    <PersistGate loading={null} persistor={persistor}>
      <Router store={store}>
        <App />
      </Router>
    </PersistGate>
  </Provider>,
  document.getElementById("root")
);
