package urlhandler

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
)

func TestSimple(t * testing.T) {
	var a string = "Hello"
	var b string = "Hello"
	assert.Equal(t, a, b, "Words should be the same")
}

func testLinkedList(t * testing.T, i int, b * UrlDirection, result []string) {
	assert.Equal(t, result[i], b.Url, "Should be the Same")
	if (i == (len(result) - 1)) {
		assert.Nil(t, b.Next, "Should be NULL")
	} else {
		assert.NotNil(t, b.Next, "Should NOT be NULL")
	}
}

func TestParseUrl(t * testing.T) {
	var b * UrlDirection = ParseUrl("/")
	result := []string{"/"}

	for i := 0; i < len(result); i++ {
		testLinkedList(t, i, b, result)
		b = b.Next
	}
}

func TestMultipleUrl(t * testing.T) {
	var b * UrlDirection = ParseUrl("/pageview/direction/index.html")
	result := []string{"/", "pageview/", "direction/", "index.html/"}
	temp := GetUrlStr(b)

	for i := 0; i < len(result); i++ {
		testLinkedList(t, i, b, result)
		b = b.Next
	}

	assert.Equal(t, "/pageview/direction/index.html", temp, "String Should Be the Same without the /")
}

func TestInvalidStart(t * testing.T) {
	var b * UrlDirection = ParseUrl("pageview/direction/index.html")
	assert.Equal(t, b == nil, true, "Invalid Start of URL")
}

func TestInvalidMiddle(t * testing.T) {
	var b * UrlDirection = ParseUrl("/pageview//direction/index.html")
	assert.Equal(t, b == nil, true, "Invalid Middle of URL")
}

func TestErrorNilHandler(t * testing.T) {
	var head * UrlHandler
	b, _ := GetHandlerUrl(head, "/")
	assert.Nil(t, b, "Should Be Null")
}

func TestErrorEmptyHandler(t * testing.T) {
	var head * UrlHandler = new(UrlHandler)
	b, n := GetHandlerUrl(head, "/");
	assert.NotNil(t, b, "Should Not Be Null")
	assert.Equal(t, b, head,  "Should Be the Same Since / is not defined")
	assert.NotNil(t, n, "/ still in stack")
	assert.Nil(t, b.H, "No Function Should Be set @ initial")
}

func TestCreateHandler(t * testing.T) {
	head := new(UrlHandler)

	SetHandler(head, "/", func(w http.ResponseWriter, r * http.Request, d * UrlDirection) bool {
		fmt.Println("Hello There from Handler Function");
		return true
	})

	callHandler, node := GetHandlerUrl(head, "/")
	assert.NotNil(t, callHandler, "Should Not be Null")
	assert.Nil(t, node, "All Elements of Linked List Should Be Popped")

	var w http.ResponseWriter
	// var r * http.Request = new(http.Request)
	callHandler.H(w, nil, node)

	callHandler, node = GetHandlerUrl(head, "/NotThere")
	assert.NotNil(t, callHandler, "Should Not Be Null because Path / exists")
	assert.NotNil(t, node, "Not NULL because NotThere should still be in stack")
	callHandler.H(w, nil, node)
}

func TestMultPath(t * testing.T) {
	head := new(UrlHandler)

	SetHandler(head, "/", func(w http.ResponseWriter, r * http.Request, d * UrlDirection) bool {
		fmt.Println("Hello There from Handler Function pt 2")
		return true
	})

	SetHandler(head, "/hello", func(w http.ResponseWriter, r * http.Request, d * UrlDirection) bool {
		fmt.Println("Hello There from Handler Function pt 2 from /hello")
		return true
	})

	SetHandler(head, "/there", func(w http.ResponseWriter, r * http.Request, d * UrlDirection) bool {
		fmt.Println("Hello There from Handler Function pt 2 from /there")
		return true
	})

	SetHandler(head, "/hello/there", func(w http.ResponseWriter, r * http.Request, d * UrlDirection) bool {
		fmt.Println("Hello There from Handler Function pt 2 from /hello/there")
		return true
	})

	callHandlerParent, _ := GetHandlerUrl(head, "/")
	assert.NotNil(t, callHandlerParent, "Parent Should Not be Null")

	var w http.ResponseWriter
	callHandlerParent.H(w, nil, nil)

	callHandlerHello, _ := GetHandlerUrl(head, "/hello")
	assert.NotNil(t, callHandlerHello, "/hello Should Not be Null")
	callHandlerHello.H(w, nil, nil)

	callHandlerThere, _ := GetHandlerUrl(head, "/there")
	assert.NotNil(t, callHandlerThere, "/there Should Not be Null")
	callHandlerThere.H(w, nil, nil)

	callHandlerHelloThere, _ := GetHandlerUrl(head, "/hello/there")
	assert.NotNil(t, callHandlerHelloThere, "/hello/there Should Not be Null")
	callHandlerHelloThere.H(w, nil, nil)

	callHandlerNobody, _ := GetHandlerUrl(head, "/hello/nobody")
	assert.NotNil(t, callHandlerNobody, "/hello/nobody should not be Null")
	assert.Equal(t, callHandlerNobody, callHandlerHello, "These Should be the Same Call Back")
	callHandlerNobody.H(w, nil, nil)

}
