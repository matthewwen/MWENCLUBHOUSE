# Purpose
This package created hash maps to handle https request (GET, PUT, PUSH). It could covert the url into a linked list. (Ex. /hello/there is "/" -> "hello/" -> "there/ -> nil) and the element inside linked list "/" and "hello\" would be the defined index. No matter what, the url path would always touch a handler; trailing characters would be passed in as an argument as a linked list.

# Archieved? 
After much debate, I decided to leave project for gorilla/mux, which already been proven to become an industry standard for companies like Microsoft. 

# Definitions
```
type Handler func(http.ResponseWriter, * http.Request)
```

# In Main Function
```
var handler Handler = nil
s := &http.Server {
	Addr: ":8080",
	Handler: handler
	ReadTimeout: 10 * time.Second,
	WriteTimeout: 10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}

# Handler
func (h Handler) ServeHTTP(w http.ResponseWriter, r * http.Request) {
	// callback should never be null, node is passed through
	var callback * urlhandler.UrlHandler
	var node * urlhandler.UrlDirection
	if str.Compare(r.Method, "GET") == 0 {
		callback, node = urlhandler.GetHandlerUrl(GetHandler, r.URL.Path) 
	} else if str.Compare(r.Method, "PUT") == 0 {
		callback, node = urlhandler.GetHandlerUrl(PutHandler, r.URL.Path) 
	} else if str.Compare(r.Method, "POST") == 0 {
		callback, node = urlhandler.GetHandlerUrl(PostHandler, r.URL.Path)
	}

	var found bool = callback != nil && callback.H != nil

	if found {
		found = callback.H(w, r, node)
	}

	if (!found) {
		http.Error(w, "Request Not Found", 404)
	}
}
```

# Declaration of Url
```
PostHandler = new(urlhandler.UrlHandler)

urlhandler.SetHandler(PostHandler, "/apiwebsite/signin", func(w http.ResponseWriter, r * http.Request, d * urlhandler.UrlDirection) bool {
	fmt.Println("Signin In Website")
	fmt.Println(r);
	return false
})

urlhandler.SetHandler(PostHandler, "/apiwebsite/signup", func(w http.ResponseWriter, r * http.Request, d * urlhandler.UrlDirection) bool {
	fmt.Println("Signin Up Website")
	fmt.Println(r);
	return false
})
```
