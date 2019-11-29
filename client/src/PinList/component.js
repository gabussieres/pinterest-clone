import React from "react";
import { FetchStatus } from "../constants/entityStatus";
import styled from "styled-components";

const Pin = styled.img`
  border-radius: 20px;
  margin: 16px 9px 0px 9px;
  width: 234px;
`;

const pin = p => (
  <a key={Math.random()} href={"/pin/" + p.id}>
    <Pin src={p.image_url} alt={p.title} />
  </a>
);

export const PinList = (pins, fetchStatus) =>
  fetchStatus === FetchStatus.loaded || FetchStatus.loading
    ? pins.map(p => pin(p))
    : null;
