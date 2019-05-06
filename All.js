import React, { Component } from "react"
import Feed from '../Components/feed'
class All extends Component {
  constructor() {
    super();
    this.state = {};
  }

render() {
    if (Array.isArray(this.props.data) || Array.isArray(this.props.searchData)){
        let data = this.props.searchData || this.props.data
        return (
            <div>
                {
                    data.map((item, index) => {
                        return(
                            <Feed data={item}/>
                        )
                    })
                }
            </div>
      
        );
    }

    else{
        return null
    }
  }
}

export default All;
