import React from "react";
import { FetchStatus } from "../constants/entityStatus";
import styled from "styled-components";

const PinListWrapper = styled.div`
  padding: 10px 10px 10px 17px;
`;

const Pin = styled.img`
  border-radius: 20px;
  margin: 9px 9px 9px 9px;
  width: 235px;
`;

const pin = p => (
  <a href={"/pin/" + p.id}>
    <Pin src={p.image_url} alt={p.title} />
  </a>
);

export const PinList = (pins, fetchStatus) => (
  <PinListWrapper>
    <p>{fetchStatus === FetchStatus.loaded ? pins.map(p => pin(p)) : null}</p>
  </PinListWrapper>
);
