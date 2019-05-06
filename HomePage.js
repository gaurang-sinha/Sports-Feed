import React, { Component } from "react";
// import Articles from "../Components/Articles";
import Videos from "../Components/Videos";
// import Slideshow from "../Components/Slideshow";
import All from "../Components/All";
import Article from "../Components/Articles";
import { getAllFeeds } from "../Apis/Helper";
// import Videoes from "../Components/Videos";
import SlideShow from "../Components/Slideshow";
import Feed from "../Components/feed";

class HomePage extends Component {
  state = {
    tab: "all",
    search: ""
  };

  handleFeed(pgName) {
    this.setState({ tab: pgName });
  }

  componentDidMount() {
    getAllFeeds().then(res => {
      this.setState(
        {
          allfeeds: res
        },
        () => console.log(this.state.allfeeds)
      );
    });
  }
  handleSearch = event => {
    if (!this.state.allfeeds) return;
    let search = event.target.value;
    this.setState({
      searchFeeds: this.state.allfeeds.filter(item => {
        return item.Title.toLowerCase().includes(search);
      })
    });
  };

  render() {
    return (
      <div>
        <div className="topnav" style={{ display: "flex" }}>
          <input
            type="text"
            placeholder="Search.."
            onChange={this.handleSearch.bind(this)}
          />
        </div>

        <div>
          <button className="tablink" onClick={() => this.handleFeed("all")}>
            All
          </button>
          <button
            className="tablink"
            onClick={() => this.handleFeed("article")}
          >
            Articles
          </button>
          <button className="tablink" onClick={() => this.handleFeed("videos")}>
            Videos
          </button>
          <button className="tablink" onClick={() => this.handleFeed("slide")}>
            SlideShow
          </button>
        </div>
        <div>
          {this.state.tab === "article" && (
            <Article data={this.state.allfeeds} />
          )}
          {this.state.tab === "all" && (
            <All
              searchData={this.state.searchFeeds}
              data={this.state.allfeeds}
            />
          )}
          {this.state.tab === "videos" && <Videos data={this.state.allfeeds} />}
          {this.state.tab === "slide" && (
            <SlideShow data={this.state.allfeeds} />
          )}
          <div />
        </div>
      </div>
    );
  }
}

export default HomePage;
