an experiment in a go backend for jam buds

### why?

i kind of hate go, so this is a good question! 

### what

layout as of now

* api - this is like "http stuff"? so it's the routing table, middleware, and route handlers, which parse params, send them to services, and respond with the results of services. it also currently holds the "grab bag" of request state (between request context + global state of the `API` struct), I'm not sure if that should be closed over instead and passed more cautiously

* models - structs representing database query results. usually these map to tables, but occasionally they may include computed data or joined results

* resources - structs representing the json returned from the API. sometimes these are just exactly the same as models (which are embedded in them), but in cases like e.g. the `User` model I manually serialize the bit I need into a resource (which is verbose, but eh)

* services - these are the controllers or logic or basically the code that _does the stuff_. in some ways, this is the most "unnecessary" layer, and all of this could just be through the route handlers. I like having the separation, though, as it keeps the handlers from getting overstuffed and makes it possible to test the logic without going through HTTP. plus if i change my mind and replace the http framework (not unlikely) this theoretically makes it easier to swap out

* store - data access code. this is also a weird thing where there's a store state that everything hangs off that contains the DB. I'm not sure how I'd make this work with service-managed transactions if I ever wanted to, since all the methods currently just get a new connection from the DB, but that's probably an okay compromise?