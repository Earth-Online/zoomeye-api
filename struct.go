package zoomeye

// user login
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// zoomeye token
type Token struct {
	AccessToken string `json:"access_token"`
}

// zoomeye user
type User struct {
	Login
	Token
}

// account resources info
type ResourcesInfo struct {
	// Account plan e.g.  developer
	Plan      string `json:"plan"`
	//  Remaining resources amount
	Resources struct {
		Host string `json:"host-search"`
		Web  string `json:"web-search"`
	}
}

// Host search return
type HostSearchInfo struct {
	// 	total count of query results
	Total     int         `json:"total"`
	// A comma-separated list of properties to get summary information on query
	Facets    interface{} `json:"facets"`
	// result host  info
	Matches   []Matches   `json:"matches"`
	// The amount that your account can get
	Available int         `json:"available"`
}

// Web search return
type WebSearchInfo struct {
	// 	total count of query results
	Total     int          `json:"total"`
	// A comma-separated list of properties to get summary information on query
	Facets    interface{}  `json:"facets"`
	// result site  info
	Matches   []WebMatches `json:"matches"`
	// The amount that your account can get
	Available int          `json:"available"`
}

type Matches struct {
	// host ip
	Ip        string   `json:"ip"`
	// Access protocol information
	Protocol  Protocol `json:"protocol"`
	// Port info
	PortInfo  PortInfo `json:"portinfo"`
	// Geographic information
	GeoInfo   GeoInfo  `json:"geoinfo"`
	Timestamp string   `json:"timestamp"`
}

type WebMatches struct {
	//  May be a domain name or ip
	Site        string   `json:"site"`
	//  site domain
	Domains     []string `json:"domains"`
	// site ip
	Ip          []string `json:"ip"`
	// site descript
	Description string   `json:"description"`

	//  response  http header
	Headers string `json:"headers"`
	// site title
	Title   string `json:"title"`
	// Geographic information
	GeoInfo   GeoInfo `json:"geoinfo"`
	Timestamp string  `json:"timestamp"`

	// Web Composents
	Components []Component `json:"component"`
	// site programming language
	Language   []string    `json:"language"`
	// site system
	System     []string    `json:"system"`
	// site database
	Db         []string    `json:"db"`
	// site use programming framework
	Framework  []string    `json:"framework"`
	// site waf
	Waf        []string    `json:"waf"`
	// site app
	Webapp     []string    `json:"webapp"`
	// site http server
	Server     []Component `json:"server"`
}

type Protocol struct {
	Application string `json:"application"`
	Probe       string `json:"probe"`
	Transport   string `json:"transport"`
}

// Web Component info e.g PHP
type Component struct {
	// Component version
	Version string `json:"version"`
	// Component English name
	Name string `json:"name"`
	// Component China name
	Chinese string `json:"chinese"`
}

type GeoInfo struct {
	Asn          int    `json:"asn"`
	Isp          string `json:"isp"`
	Organization string `json:"organization"`
	Aso          string `json:"aso"`

	City         AddressInfo `json:"city"`
	Country      AddressInfo `json:"country"`
	Continent    AddressInfo `json:"continent"`
	Subdivisions AddressInfo `json:"subdivisions"`

	Location struct {
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lon"`
	}
}
type AddressInfo struct {
	GeonameId int    `json:"geoname_id"`
	Code      string `json:"code"`
	Name      struct {
		Zh string `json:"zh-cn"`
		En string `json:"en"`
	} `json:"name"`
}

type PortInfo struct {
	Port      int    `json:"port"`
	Product   string `json:"product"`
	Hostname  string `json:"hostname"`
	Service   string `json:"service"`
	Title     string `json:"title"`
	Os        string `json:"os"`
	Version   string `json:"version"`
	Info      string `json:"info"`
	Device    string `json:"device"`
	Extrainfo string `json:"extrainfo"`
	Rdns      string `json:"rdns"`
	Banner    string `json:"banner"`
}
