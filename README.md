# Blog-Platform
Building a distributed backend system for blogging platform using microservice artictecture.

## Why ?

To learn and build how large scalable and distributed platforms are build.

### What we are building?

We will build four microservices, hosted as four separate applications. Article, User, Tag, Comment are the four microservices that will be implemeted in Golang

### 1. Articles microservice
Each article consists of text, a title or headline, an author, and timestamps for the article’s creation and the last time the article was modified.

### 2. Tags microservice
Each article can be have one or more tags associated with it. Since this API is exposed separately from the Articles API, individual articles are referred to by URL.

### 3. Comments microservice
Users can post comments on each article. As with the tags microservice, individual articles are referred to by URL. Each comment has an author and a date.

### 4. Users microservice
Each user has a display name shown to other users, an email address (used as a username when logging in), and a hashed password.


## Authors

* **Bhuvnesh Maheshwari** - *bhuvnesh13396@gmail.com* - [Github Link](https://github.com/bhuvnesh13396)
* **Nirav Panchal**


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
