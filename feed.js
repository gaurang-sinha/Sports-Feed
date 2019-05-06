import React, { Component } from "react";
class Feed extends Component {
  render() {
    return (
      <div
        style={{ border: "2px solid", display: "flex", flexDirection: "row" }}
      >
        <div style={{ flex: "3" }}>
          <div>{this.props.data.Title}</div>
          <div style={{ marginTop: "40px" }}>{this.props.data.Name}</div>
        </div>

        <div style={{ flex: "1" }}>
          <img
            src={this.props.data.Thumbnail}
            style={{ width: "100%", height: "100%" }}
          />
        </div>
      </div>
    );
  }
}

export default Feed;
