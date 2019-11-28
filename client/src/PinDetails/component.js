import React, { Component } from "react";
import { Row, Col, Button } from "react-bootstrap";

import { FetchStatus } from "../constants/entityStatus";
import {
  Wrapper,
  RightWrapper,
  LeftWrapper,
  Image,
  SourceUrl,
  Title
} from "./styles";

class PinDetails extends Component {
  constructor(props) {
    super(props);
    props.fetchPinDetails(props.pinId);
  }

  render() {
    const { details, error, fetchStatus } = this.props.pinDetails.pin;
    if (error != null) {
      return null;
    }
    const { title, description, image_url, source_url, id } = details;
    const handleClick = pinId => {
      this.props.deletePin("gabriel", pinId);
    };
    return fetchStatus === FetchStatus.loaded ? (
      <Wrapper className="App">
        <div>
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
                <Button onClick={() => handleClick(id)}>Delete Pin</Button>
              </RightWrapper>
            </Col>
          </Row>
        </div>
      </Wrapper>
    ) : null;
  }
}

export default PinDetails;
