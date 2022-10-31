
### Hi there 👋, I'm Ridwan Rais!
This is my personal project that I've been working on for a week; a backend API for a YouTube clone website.

 I'm using Swagger 2.0 to create the API documentation as well as to generate the request and response structs, as well as the middleware for the handler. 

For the database, I use MongoDB because I've been using PostgreSQL and MySQL so far and want to try something different. I'm a little bit surprised because it's easier than expected to implement it!

Last but not least, I implement Clean Architecture paradigm to make the API more scalable and easier to maintain.

You can see the API documentation by visiting [this link](https://app.swaggerhub.com/apis/ridwanrais2/youtube-clone-api/1)

Frontend Repo: COMING SOON!

Deployed Web App: COMING SOON!

 ## 🧑‍💻  Tech Stack
-   Golang
-   Swagger 2.0 (used go-swagger as the api generator)
-   MongoDB
-   JSON Web Tokens

## 🤖  Set Environment Variables
Copy & paste the `.env.example` file into the same directory. 
Afterward, rename the new file to `.env` and insert the appopriate values into the variables
-   Create a MongoDB database (either locally or remotely) and put the mongo uri string into MONGODB_URI.
-  Copy and paste the base url of your frontend app into CLIENT_URL. If you wish to let your API to accept requests from any source, change the string to "*".
-   ACCESS_TOKEN_KEY can be any string.

## 📦  Setup

To run this app locally,


<pre><code>
$ git clone https://github.com/ridwanrais/youtube-api-golang.git
$ cd youtube-api-golang
$ go mod tidy
$ make run
</code></pre>

> Don't forget to set environment variables after cloning the repo

The app wil start on  [http://localhost:8099](http://localhost:5000/)
