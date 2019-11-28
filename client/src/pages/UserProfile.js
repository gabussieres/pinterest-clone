import React from "react";
import { Row } from "react-bootstrap";

import { AppWrapper } from "./styles";
import Navigation from "../Navigation/component";
import UserContainer from "../User/container";

const UserProfile = userId => (
  <AppWrapper className="App">
    <Navigation />
    <Row>
      <UserContainer userId={userId} />
    </Row>
  </AppWrapper>
);

export default UserProfile;
