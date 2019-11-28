import React from "react";

import { AppWrapper } from "./styles";
import Navigation from "../Navigation/container";
import FeedContainer from "../Feed/container";

const Feed = () => (
  <AppWrapper className="App">
    <Navigation />
    <FeedContainer />
  </AppWrapper>
);

export default Feed;
