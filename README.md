# Shoe Store using Go and Vue.js
## Why this tech stack?
As we advance in seniority, our responsibilty as Software Engineers is to keep up-to-date with new 
technologies/architectures/system designs and make informed decisions on what's best given a specific situation.

I decided to approach this challenge from that POV: Create a POC using technologies and architecture principles (that
are completely foreign to me) in an effort to make a case for or against it.

Go, for web applications, has been gaining in popularity over the last few years, whereas RoR's popularity has been
declining. The task was relatively small, but I chose to design it using principles from Domain Driven Design and Clean
Architecture. DDD is primarily used for microservices. In my opinion, it organizes the code in a consistent and 
informative way, which makes it great for scalability.
Because of the decoupled nature of the file structure, if we ever wanted to add different domains, or change DBs, 
it would be very easy to scale with little change to code or understandability.

When deciding on the frontend, since I am, admittedly, not the best at making pretty UIs and don't have much experience
coding frontends, it came down to choosing a framework that was relatively lightweight, new to the FE community, had
good documentation and community support, and easy for a backend developer to use. 
This is why I chose Vue.js.

## Features
* Real-time inventory updates
* Real-time sales
* Shoe transfer suggestions
* Visual alerts of low-stock and _almost_ low-stock shoe models
* Stores and Shoe Models automatically sorted by sales
* Best and Worst Sellers per Store
* Total sales per store
* Total sales per shoe model
* REST API

API routes can be found at /shoe-store-server/api/routes.go
I used Pusher for the real-time updates.

## Given more time
I spent a lot of time just learning the ins and outs of Go. It's a strongly typed language, that feels like a "simpler"
version of Java/C++. A lot of time was spent debugging and figuring out when (and when NOT) to use pointers, how to
implement Websockets, and choosing the right libraries to get the job done. This is a list of things I would have liked
to add to the project.

Backend:
* Implement the Repository pattern with at least one other DB to showcase the benefits of DDD
* Implement JWT for access to the service
* Implement the REST API as gRPC

Frontend:
* Sorting and Filtering on all tables

Features:
* Implement sign-in/sign-up
* Implement a user profile
* Implement user permissions
* Implement the "Get by time" where applicable
* Implement the ability to create/add/delete a Store/Shoe Model
* Allow Admin to edit the LOW_STOCK/HIGH_STOCK value
* Execute shoe transfer


Analytics:
* Implement some reporting to track:
  * sales per month (per shoe model)
  * shoe transfers per month
* Implement some reporting to monitor the health of the service

Infrastructure:
* Add docker!
* Unit tests!
* Create GitHub actions to run tests and linter before merging
* Add Redis for caching

## Requirements
* go 1.21
* vue 3
* sqlite

## Running the project