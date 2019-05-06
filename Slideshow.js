import React, { Component } from 'react';
import Feed from './feed';
class Slideshow extends Component {
    state = {  }
    render() { 
        if(Array.isArray(this.props.data))
        return (  
            <div>
                {this.props.data.map((item,index)=>{
                    if (item.Type === 'slideshow')
                    return <Feed data ={item}/>;
                })}
            </div>
        );
    }
}
 
export default Slideshow;