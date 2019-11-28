import React from "react";
import git_icon from "../images/GitHub-Mark-64px.png";
import logo_img from "../images/Logo.png";
import { Nav, NavList, NavButton, LogoWrapper } from "./styles";

export const Navigation = () => {
  return (
    <Nav>
      <NavList>
        <NavButton logo="true" href="/">
          <LogoWrapper>
            <img src={logo_img} alt="Logo" height="35" width="35" />
          </LogoWrapper>
        </NavButton>
        <NavButton
          image
          href="https://github.com/gabussieres/insight"
          target="_blank"
        >
          <img src={git_icon} alt="GitHub" height="30" width="30" />
        </NavButton>
        <NavButton href="/gabriel">Gabriel</NavButton>
        <NavButton href="/">Following</NavButton>
        <NavButton home="true" href="/">
          Home
        </NavButton>
      </NavList>
    </Nav>
  );
};

export default Navigation;
