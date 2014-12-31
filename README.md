# Go HTTP Router Test

Go HTTP Router Test tests router behavior and documents idiomatic code for
solving common problems with common HTTP routers (multiplexers, muxers, etc.)

If you see anything that needs fixing feel free to create an issue or
pull request. I'm especially looking to fix:

* Anything that is wrong or not idiomatic for the router in question. I don't
    use most of these (I only need one!), and I don't want to portray any
    unfairly, so if you see something, please please create an issue with
    a link to the documentation I missed or whatever.
* Any tests that are testing the wrong thing. I'm testing for the sanest
    behavior I know, but if you have wisdom to share, feel free to create
    an issue.

I'm expecting this to be a rationale heavy project. That's kind of its purpose.

## The Story

For some reason the standard library's http package and I have never got
along very well. Maybe it's me, or maybe I just don't know how to use it,
but whatever the case, the many alternatives suggest I'm not the only one.

Performance hasn't been my concern, but I do enjoy high performance.
Fortunately, [Julien Schmidt's benchmark][b] has been triggering advances
there.

[b]: https://github.com/julienschmidt/go-http-routing-benchmark

The concerns I did have led me to to try something besides the standard
library, and at first I was quite excited. But as I put in more time with
an alternative, I began to discover problems with functionality and to notice
exactly how much difference from the standard library I was really getting.
As [Blake Mizerany so eloquently points out][f], sometimes it's not much.

[f]: https://www.youtube.com/watch?v=yi5A3cK1LNA

So I'm starting this project to work out precisely what functionality I expect
from an HTTP router and to compare code required to achieve it with the
different routers.

# Cast of characters

* [Standard Library](http://golang.org/pkg/net/http/)
* [Goji](https://github.com/zenazn/goji)
* [Gorilla Mux](http://www.gorillatoolkit.org/pkg/mux)

# The Tests

## Method Not Allowed

```
GET / HTTP/1.1

HTTP/1.1 404 Not Found

```

[Wat][w]?

[w]: https://www.destroyallsoftware.com/talks/wat

Okay, only accepting POST on the root is unlikely, but it makes for an eloquent
example.

Implemented for:

* Standard Library (passed)
* Goji (failed)
* Gorilla Mux (can pass, but only with non-idiomatic code)

Installing handlers for all known but unsupported HTTP verbs doesn't count.
It isn't future proof, and it means installing them all (probably by writing
a function to do it en masse), and that's just silly.

## HEAD

Even if the server generates the whole page for HEAD requests, it should not
try to put it on the wire. It should especially not reply with 405.

Unimplemented

## Custom Not Found Page

Servers should support custom 404 pages. For instance, if I follow a link to a
project on GitHub that has gone away, I should be able to look for the myriad
of forks that probably exist without having to go to GitHub's index page to get
a search form.

Unimplemented

## Custom Not Found Page with File Server

If file serving functionality is offered, it should allow custom not found
pages, too, and [without sacrificing use of sendfile and such][s].

[s]: http://avtok.com/2014/11/05/interface-upgrades.html

Unimplemented

