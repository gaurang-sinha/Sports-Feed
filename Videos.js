import React, { Component } from "react";
import Feed from "../Components/feed";

class Videoes extends Component {
  state = {};
  render() {
    if (Array.isArray(this.props.data))
      return (
        <div>
          {this.props.data.map((item, index) => {
            if (item.Type === "video")
              return <Feed data={item} />;
          })}
        </div>
      );
    else {
      return null;
    }
  }
}

export default Videoes;
