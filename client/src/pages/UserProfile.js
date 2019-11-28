import React from "react";

import { AppWrapper } from "./styles";
import Navigation from "../Navigation/container";
import UserContainer from "../UserProfile/container";

const UserProfile = userId => (
  <AppWrapper className="App">
    <Navigation />
    <UserContainer userId={userId} />
  </AppWrapper>
);

export default UserProfile;
