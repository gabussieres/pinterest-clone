import React from "react";
import { Row } from "react-bootstrap";

import { AppWrapper } from "./styles";
import Navigation from "../Navigation/component";
import PinContainer from "../PinDetails/container";

const Pin = pinId => (
  <AppWrapper className="App">
    <Navigation />
    <Row>
      <PinContainer pinId={pinId} />
    </Row>
  </AppWrapper>
);

export default Pin;
