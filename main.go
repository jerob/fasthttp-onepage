package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
  "io/ioutil"
  "os"
)

var (
    str string
    contentType string
)

func main() {
  CONTENT, err := ioutil.ReadFile("file") // just pass the file name
  if err != nil {
      fmt.Print(err)
  }
  str = string(CONTENT) // convert content to a 'string'
  _ = str
  fmt.Print("CONTENT_TYPE = " + os.Getenv("CONTENT_TYPE"))

  if os.Getenv("CONTENT_TYPE")=="json" {
    contentType = "application/json; charset=utf8"
  } else if os.Getenv("CONTENT_TYPE")=="xml" {
    contentType = "application/xml; charset=utf8"
  } else if os.Getenv("CONTENT_TYPE")=="javacript" {
    contentType = "application/javascript; charset=utf8"
  } else if os.Getenv("CONTENT_TYPE")=="csv" {
    contentType = "text/csv; charset=utf8"
  } else if os.Getenv("CONTENT_TYPE")=="css" {
    contentType = "text/css; charset=utf8"
  } else if os.Getenv("CONTENT_TYPE")=="png" {
    contentType = "image/png"
  } else if os.Getenv("CONTENT_TYPE")=="gif" {
    contentType = "image/gif"
  } else if os.Getenv("CONTENT_TYPE")=="jpeg" {
    contentType = "image/jpeg"
  } else {
    contentType = "text/html; charset=utf8"
  }

  h := requestHandlerText

	h = fasthttp.CompressHandler(h)

	if err := fasthttp.ListenAndServe(":80", h); err != nil {
	}
}

func requestHandlerImage(ctx *fasthttp.RequestCtx) {
  // fmt.Fprintf(ctx, "??")
  // ctx.SetContentType("text/plain; charset=utf8")
  ctx.SetContentType(contentType)
  fmt.Fprintf(ctx, str)
}

func requestHandlerText(ctx *fasthttp.RequestCtx) {
  // fmt.Fprintf(ctx, "??")
  // ctx.SetContentType("text/plain; charset=utf8")
  ctx.SetContentType(contentType)
  fmt.Fprintf(ctx, str)
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, world!\n\n")

	fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
	fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
	fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
	fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
	fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
	fmt.Fprintf(ctx, "User-Agent is %q\n", ctx.UserAgent())
	fmt.Fprintf(ctx, "Connection has been established at %s\n", ctx.ConnTime())
	fmt.Fprintf(ctx, "Request has been started at %s\n", ctx.Time())
	fmt.Fprintf(ctx, "Serial request number for the current connection is %d\n", ctx.ConnRequestNum())
	fmt.Fprintf(ctx, "Your ip is %q\n\n", ctx.RemoteIP())

	fmt.Fprintf(ctx, "Raw request is:\n---CUT---\n%s\n---CUT---", &ctx.Request)

	ctx.SetContentType("text/plain; charset=utf8")

	// Set arbitrary headers
	ctx.Response.Header.Set("X-My-Header", "my-header-value")

	// Set cookies
	var c fasthttp.Cookie
	c.SetKey("cookie-name")
	c.SetValue("cookie-value")
	ctx.Response.Header.SetCookie(&c)
}
