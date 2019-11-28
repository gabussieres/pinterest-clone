import React, { Component } from "react";
import { Row, Col } from "react-bootstrap";

import { FetchStatus } from "../constants/entityStatus";
import { LeftWrapper, RightWrapper, ImageWrapper } from "./styles";
import { PinList } from "../PinList/component";

class UserProfile extends Component {
  constructor(props) {
    super(props);
    props.fetchUserProfile(props.userId);
    props.fetchUserPins(props.userId);
  }

  render() {
    const { user, error, fetchStatus } = this.props.user.userProfile;
    const { name, followers, following, image_url } = user;
    if (error != null) {
      return null;
    }
    return (
      <div>
        {fetchStatus === FetchStatus.loaded ? (
          <Row>
            <div>
              <Col md={8}>
                <LeftWrapper>
                  <Row>
                    <h1>{name}</h1>
                  </Row>
                  <Row>
                    {followers} followers · {following} following
                  </Row>
                </LeftWrapper>
              </Col>
              <RightWrapper>
                <Col md={3}>
                  <ImageWrapper src={image_url} alt={name}></ImageWrapper>
                </Col>
              </RightWrapper>
            </div>
          </Row>
        ) : null}
        {PinList(this.props.user.pins.pins, fetchStatus)}
      </div>
    );
  }
}

export default UserProfile;
