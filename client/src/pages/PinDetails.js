import React from "react";

import { AppWrapper } from "./styles";
import Navigation from "../Navigation/container";
import PinContainer from "../PinDetails/container";

const Pin = pinId => (
  <AppWrapper className="App">
    <Navigation />
    <PinContainer pinId={pinId} />
  </AppWrapper>
);

export default Pin;
