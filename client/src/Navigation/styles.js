import styled, { css } from "styled-components";

export const Nav = styled.div`
  border-bottom: 0.5px solid #ebebeb;
`;

export const NavList = styled.ul`
  list-style-type: none;
  overflow: hidden;
`;

export const NavButton = styled.a`
text-decoration: none !important
color: #8f8f8f
float: right;
font-weight: bold;
font-size: 16px;
padding: 1.7em 1em 1em

${props =>
  props.logo &&
  css`
    font-size: 25px;
    float: left;
    padding: 0.7em 0em 0.35em;
  `}

${props =>
  props.image &&
  css`
    padding: 1.4em 1em 0.75em;
  `}

${props =>
  props.home &&
  css`
    color: #3b3b3b;
  `}
`;

export const LogoWrapper = styled.div`
  display: inline;
`;
