package main

import (
	"encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
	"os"
	"github.com/jinzhu/gorm"
     _ "github.com/jinzhu/gorm/dialects/mysql"
)


type Response struct {
          
    Feed []Feed `json:"feed"`
}


type Feed struct {
    
	Thumbnail string `json: "thumbnail"`
	Id float64 `json: "id"`
	Title string `json: "title"`
	Modified_date string `json: "modifies_date"`
	Permalink string `json: "permalink"`
	Published_date string `json: "published_date"`
	Read_count string `json: "read_count"`
	Comment_count string `json: "comment_count"`
	Live_traffic string `json: "live_traffic:"`
	Rank string `json: "rank"`
	Index string`json: "index"`
	Type string `json: "type"`
	Excerpt string `json:"excerpt"`
	Author Author `json:"author"`
	Algo_meta Algo_meta `json: "algo_meta"`
	Post_tag []string `json: "post_tag"`
	
}


type Author struct {
	Name string `json:"name"`
	
}

type Algo_meta struct {
	Recency string `json: "recency"`
	Sk_reads string `json: "sk_reads"`
	Section_breaking string `json: "section_breaking"`
	Section_live string `jsoon:"section_live"`
	Number_of_comments string `json: "number_of_comments"`
	Sk_live_traffic string `json: "sk_live_traffic"`
	Sport_rank string `json: "sport_rank"`
	Event_rank string `json: "event_rank"`
	Author_rank string `json: "author_rank"`
	Type_rank string `json: "type_rank"`
	Personal_score string `json: "personal_score"`

}

type Feeds struct {
	gorm.Model
	Thumbnail string `gorm:"size:800"`
	Modified_date string `gorm:"size:800"`
	Title string	`gorm:"size:800"` 
	Permalink string `gorm:"size:800"`
	Published_date string `gorm:"size:800"`
	Read_count string	`gorm:"size:800"`
	Comment_count string `gorm:"size:800"`
	Live_traffic string	`gorm:"size:800"`
	Rank 		string	`gorm:"size:800"`
	Index 	string		`gorm:"size:800"`
	Type string 	`gorm:"size:800"`
	Excerpt string 	`gorm:"size:800"`
	Author Author 	`gorm:"size:800"`
	Algo Algo_meta 	`gorm:"size:800"`
}

type Authors struct {
	gorm.Model
	Name string `gorm:"size:100"`
}

type Algo_metas struct {
	gorm.Model
	Recency string `gorm:"size:100"`
	Sk_reads string `gorm:"size:100"`
	Section_breaking string `gorm:"size:100"`
	Section_live string `gorm:"size:100"`
	Number_of_comments string `gorm:"size:100"`
	Sk_live_traffic string `gorm:"size:100"`
	Sport_rank string `gorm:"size:100"`
	Event_rank string `gorm:"size:100"`
	Author_rank string `gorm:"size:100"`
	Type_rank string `gorm:"size:100"`
	Personal_score string `gorm:"size:100"`
}

func populateDB(){
	db, err :=gorm.Open("mysql", "root:Gori@1234@/BrainlyCrawler?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
	  panic("failed to connect database")
	}

	defer db.Close()
	db.CreateTable(&Feeds{})
	

	response, err := http.Get("http://data.sportskeeda.com/feed.json")
	
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
	}
	
	responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	for i :=0;i<len(responseObject.Feed);i++{
		feed:=Feeds{
			Thumbnail: responseObject.Feed[i].Thumbnail,
			Modified_date: responseObject.Feed[i].Modified_date,
			Title: responseObject.Feed[i].Title, 	
			Permalink: responseObject.Feed[i].Permalink,  
			Published_date: responseObject.Feed[i].Published_date,  
			Read_count: responseObject.Feed[i].Read_count,  
			Comment_count: responseObject.Feed[i].Comment_count,  
			Live_traffic : responseObject.Feed[i].Live_traffic, 
			Rank : responseObject.Feed[i].Rank, 
			Index : responseObject.Feed[i].Index, 
			Type : responseObject.Feed[i].Type,
			Excerpt : responseObject.Feed[i].Excerpt,
			

		}
		author:=Authors{
			Name: responseObject.Feed[i].Author.Name,
		}
		algo_metas := Algo_metas {
			Recency:  responseObject.Feed[i].Algo_meta.Recency,
			Sk_reads:  responseObject.Feed[i].Algo_meta.Sk_reads,
			Section_breaking: responseObject.Feed[i].Algo_meta.Section_breaking,
			Section_live:  responseObject.Feed[i].Algo_meta.Section_live,
			Number_of_comments:  responseObject.Feed[i].Algo_meta.Number_of_comments,
			Sk_live_traffic : responseObject.Feed[i].Algo_meta.Sk_live_traffic,
			Sport_rank : responseObject.Feed[i].Algo_meta.Sport_rank, 
			Event_rank : responseObject.Feed[i].Algo_meta.Event_rank,
			Author_rank : responseObject.Feed[i].Algo_meta.Author_rank, 
			Type_rank : responseObject.Feed[i].Algo_meta.Type_rank,
			Personal_score : responseObject.Feed[i].Algo_meta.Personal_score,
		}
		
		
		db.Create(&feed)
		db.Create(&author)
		db.Create(&algo_metas)
	}
		
	
}
}
	


func handler(w http.ResponseWriter, r *http.Request) {
	
	db, err :=gorm.Open("mysql", "root:Gori@1234@/BrainlyCrawler?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
	  panic("failed to connect database")
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	defer db.Close()
	feeds:=[]Feed{}
	db.Find(&feeds)
	
	authors:=[]Author{}
	db.Find(&authors)
	algo_metas:=[]Algo_metas{}
	db.Find(&algo_metas)

	type res map[string]interface{}
	var obje [] res
	for i :=0;i<len(feeds);i++{
		
		
		m:=res{
			"Thumbnail": feeds[i].Thumbnail ,
			"Modifies_date": feeds[i].Modified_date , 
			"Title": feeds[i].Title ,
			"Permalink" :feeds[i].Permalink ,
			"Published_date" : feeds[i].Published_date ,
			"Read_count" : feeds[i].Read_count ,
			"Comment_count" : feeds[i].Comment_count ,
			"Live_traffic" : feeds[i].Live_traffic ,
			"Rank" : feeds[i].Rank ,
			"Index" : feeds[i].Index ,
			"Type" : feeds[i].Type ,
			"Excerpt" : feeds[i].Excerpt ,
			"Name" : authors[i].Name ,
			"Recency" : algo_metas[i].Recency ,
			"Sk_reads" : algo_metas[i].Sk_reads ,
			"Section_breaking" : algo_metas[i].Section_breaking ,
			"Section_live" : algo_metas[i].Section_live , 
			"Number_of_comments" : algo_metas[i].Number_of_comments ,
			"Sk_live_traffic" : algo_metas[i].Sk_live_traffic ,
			"Sport_rank" : algo_metas[i].Sport_rank ,
			"Event_rank" : algo_metas[i].Event_rank ,
			"Author_rank" : algo_metas[i].Author_rank ,
			"Type_rank" : algo_metas[i].Type_rank ,
			"Personal_score" : algo_metas[i].Personal_score ,

		}
		

		obje = append(obje,m)

	}
	json.NewEncoder(w).Encode(obje)
}


func main() {
	populateDB()
	
	http.HandleFunc("/all/", handler)
	http.ListenAndServe(":8000", nil) 
}