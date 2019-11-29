import React, { Component } from "react";
import { Row, Col, Button } from "react-bootstrap";

import { FetchStatus } from "../constants/entityStatus";
import {
  Wrapper,
  RightWrapper,
  LeftWrapper,
  Image,
  SourceUrl,
  Title,
  SignInPrompt
} from "./styles";

class PinDetails extends Component {
  constructor(props) {
    super(props);
    if (this.props.pinDetails.user.loggedIn) {
      props.fetchPinDetails(props.pinId);
    }
  }

  render() {
    const { details, error, fetchStatus } = this.props.pinDetails.pin;

    if (error != null) {
      return null;
    }

    const { title, description, image_url, source_url, id } = details;
    const handleClick = pinId => {
      this.props.deletePin(this.props.pinDetails.user.id, pinId);
    };

    return this.props.pinDetails.user.loggedIn ? (
      fetchStatus === FetchStatus.loaded ? (
        <Wrapper className="App">
          <Row>
            <Col lg={5}>
              <LeftWrapper>
                <Image src={image_url} alt={title} width="475px" />
              </LeftWrapper>
            </Col>
            <Col lg={7}>
              <RightWrapper>
                <SourceUrl href={source_url}>{source_url}</SourceUrl>
                <Title>{title}</Title>
                <p>{description}</p>
                {this.props.pinDetails.user.pins.map(p => p.id).includes(id) ? (
                  <Button onClick={() => handleClick(id)}>Delete Pin</Button>
                ) : null}
              </RightWrapper>
            </Col>
          </Row>
        </Wrapper>
      ) : null
    ) : (
      <SignInPrompt>
        <h1>Please sign in!</h1>
      </SignInPrompt>
    );
  }
}

export default PinDetails;
