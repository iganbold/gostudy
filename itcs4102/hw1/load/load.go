package load

import (
        "log"
        "fmt"
        "encoding/xml"
        "encoding/json"
	    "os"
	    "errors"
	    "net/http"
)

const rssJson = "load/rssList.json"

//url json data structure
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
}

//rss document structure
type (
	item struct {
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
	}

	channel struct {
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		Item           []item   `xml:"item"`
	}

	rssDoc struct {
		Channel channel  `xml:"channel"`
	}
)

//Starting point
func Run() {
    
    //getting Feeds from a file
    feeds, err := getFeeds()
    
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Fetching the news from the given web sites...")
    for _, feed := range feeds {
        err := load(feed)
        if err != nil {
            fmt.Println(err)
            return
        }
    }
}


func load(feed *Feed) (error) {
    
    doc, err := pullDoc(feed)
    
    if err != nil {
        return err
    }
    
    //Prints the items of the rss document
    for _,item := range doc.Channel.Item {
        fmt.Println("\033[33m",doc.Channel.Title," : \033[32m",item.Title)
    }
    
    return nil
}

//Pull rss feeds from given web site
func pullDoc(feed *Feed) (*rssDoc , error) {
    
    if feed.URI == "" {
        return nil, errors.New("Error: NO URI");
    }
    
    //sending http get request and getting response
    response, err := http.Get(feed.URI)
    
    if err != nil {
        return nil, err
    }
    
    //closing response after return
    defer response.Body.Close()
    
    //checking http response status code 
    if response.StatusCode != 200 {
        return nil, errors.New("Error: HTTTP RESPONSE ");
    }
    
    var doc rssDoc
    
    //Unmarshalling XML to rssDoc
    err = xml.NewDecoder(response.Body).Decode(&doc)
    return &doc, err
} 

//Read json based url list from rssList.json file
func getFeeds() ([]*Feed,error){
    //opening a file
	file, err := os.Open(rssJson)
	
	if err != nil {
		return nil , err
	}
	
	//close a file when after return
	defer file.Close()
	
	var feeds []*Feed
	
	//Unmarshalling json to Feed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds,err
}