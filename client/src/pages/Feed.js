import React from "react";
import { Row } from "react-bootstrap";

import { AppWrapper } from "./styles";
import Navigation from "../Navigation/component";
import FeedContainer from "../Feed/container";

const Feed = () => (
  <AppWrapper className="App">
    <Navigation />
    <Row>
      <FeedContainer />
    </Row>
  </AppWrapper>
);

export default Feed;
