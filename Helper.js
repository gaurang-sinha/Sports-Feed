export function getAllFeeds(){
    let url = "http://localhost:8000/all/";
    return fetch(url)
      .then(res => {
        return res.json();
      })
      .then(res => {
        return res;
      });
  };