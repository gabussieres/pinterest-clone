import React from "react";
import GitHubLogin from "react-github-login";

import logo_img from "../images/Logo.png";
import { Nav, NavList, NavButton, LogoWrapper } from "./styles";

const onSuccess = response => {
  console.log(response);
  return props => {
    // Ideally, I would authenticate here
    props.fetchUser("gabriel");
  };
};
const onFailure = response => console.error(response);

const Navigation = () => (
  <Nav>
    <NavList>
      <NavButton logo="true" href="/">
        <LogoWrapper>
          <img src={logo_img} alt="Logo" height="35" width="35" />
        </LogoWrapper>
      </NavButton>
      <NavButton>
        <GitHubLogin
          clientId="ce9f5d621c5709bf97e5"
          onSuccess={() => onSuccess()(this.props)}
          onFailure={onFailure}
          redirectUri=""
        />
      </NavButton>
      <NavButton href="/gabriel">Gabriel</NavButton>
      <NavButton home="true" href="/">
        Home
      </NavButton>
    </NavList>
  </Nav>
);

export default Navigation;
