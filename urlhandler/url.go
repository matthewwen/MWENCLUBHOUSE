package urlhandler

import (
	"fmt"
	"net/http"
)

type Handler func(http.ResponseWriter, * http.Request, * UrlDirection) bool

type UrlDirection struct {
	Url string
	Next * UrlDirection
}

type UrlHandler struct {
	M map[string] * UrlHandler;
	H Handler;
}

func createUrlDirection(head ** UrlDirection, tail ** UrlDirection, str string) {
	obj := new(UrlDirection)
	obj.Url  = str
	obj.Next = nil

	if (*head == nil) {
		*head = obj
	} else {
		(*tail).Next = obj
	}

	*tail = obj;
}

func ParseUrl(url string) * UrlDirection {
	var isValid bool = url[0] == '/'
	
	var head *UrlDirection
	var tail *UrlDirection

	var startidx int = 0;
	var i int
	for i = 0; i < len(url) && isValid; i++ {
		if url[i] == '/' {
			if ( i - startidx == 0 ) {
				isValid = startidx == 0
			}

			if isValid {
				createUrlDirection(&head, &tail, url[startidx: i + 1])
			}
			startidx = i + 1
		}
	}

	if !isValid {
		startidx = 0
		i = 0
		head = nil
	}

	if (i != startidx) {
		createUrlDirection(&head, &tail, url[startidx: i] + "/")
	}
	return head 
}

func SetHandler(head * UrlHandler, url string, handler Handler) {
	var list *UrlDirection
	if (head != nil) {
		list = ParseUrl(url)
	} else {
		fmt.Println("Head should be Allocaed Before Using in SetHandler Method")
	}

	for ; list != nil; list = list.Next {
		helperfunction := func(head *UrlHandler, dir string) * UrlHandler {
			isValid := head != nil

			// creates a hash map if it does not exist
			if isValid && head.M == nil{
				head.M = make(map[string]*UrlHandler)
				isValid = head.M != nil
			}

			// gets the handler from hash map
			var h * UrlHandler
			if isValid {
				h = head.M[dir] 
			}

			// creates handler if not exist
			if h == nil {
				h = new(UrlHandler)
			}
			head.M[dir] = h

			return h
		}
		head = helperfunction(head, list.Url)
	}

	if (head != nil) {
		head.H = handler;
	}
}

func GetHandlerUrl(head * UrlHandler, url string) (* UrlHandler, * UrlDirection) {
	var list * UrlDirection;
	if (head != nil) {
		list = ParseUrl(url)
	}

	var h * UrlHandler   = head;
	var l * UrlDirection = list;
	for ; list != nil && head != nil; list = list.Next {

		// head set to path of next, nil if not found or no hash map
		if (head.M != nil) {
			head = head.M[list.Url]
		} else {
			head = nil
		}

		// update if path exists or there is hash map
		if (head != nil) {
			h = head;
			l = list.Next;
		}
	}

	return h , l 
}

func GetUrlStr(d * UrlDirection) string {
	var s string

	for ; d != nil; d = d.Next {
		s = s + d.Url
	}

	if (len(s) > 1) {
		s = s[:len(s) - 1]
	}

	return s
}