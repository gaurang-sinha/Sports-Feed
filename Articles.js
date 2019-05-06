import React, { Component } from "react";
import Feed from "../Components/feed";
class Article extends Component {
  render() {
    if (Array.isArray(this.props.data))
      return (
        <div>
          {this.props.data.map((item, index) => {
            if (item.Type === "article") return <Feed data={item} />;
          })}
        </div>
      );
    else {
      return null;
    }
  }
}

export default Article;
