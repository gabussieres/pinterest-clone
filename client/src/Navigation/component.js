import React from "react";
import GitHubLogin from "react-github-login";

import logo_img from "../images/Logo.png";
import { Nav, NavList, NavButton, LogoWrapper } from "./styles";

const defaultUser = "gabriel";

const onSuccess = response => {
  return props => {
    // Ideally, I would authenticate here
    // Instead we're going to default to defaultUser
    props.fetchUser(defaultUser);
    props.fetchUserPins(defaultUser);
  };
};
const onFailure = response => console.error(response);

const Navigation = props => (
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
          onSuccess={() => onSuccess()(props)}
          onFailure={onFailure}
          redirectUri=""
        />
      </NavButton>
      <NavButton href={`/${props.user.user.details.id}`}>
        {props.user.user.details.first_name}
      </NavButton>
      <NavButton home="true" href="/">
        Home
      </NavButton>
    </NavList>
  </Nav>
);

export default Navigation;
