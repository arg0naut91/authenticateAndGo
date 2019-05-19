Authenticate and Go!
====================

<img src="markdown/authGIF.gif" width="750" height="500" />

Purpose
-------

The aim of this project is to work on web authentication module with all
the major components integrated:

-   Register and login through the database (currently this is
    `PostgreSQL`) with full encryption;  
-   Register/Login through `OAuth2` (e.g. social networks - currently
    Facebook, Google and GitHub are supported);  
-   Persistent login (currently this works for the standard login);  
-   Password recovery (not yet implemented).

Remarks on (local) setup
------------------------

To test and/or further develop all functionalities, you would need:

-   `PostgreSQL` installed, the password stored in environment variable
    `PGSQL` and a table named `users` (with columns `username`, `email`,
    `password` of type `text`) in a database called `mydb`. You’d also
    need to adjust the configuration of `db.go` in `dbAuth` if your
    username is not `postgres` and port not `5432`;  
-   Applications created on Google, Facebook and GitHub with environment
    variables declared in `socialAuth` sub-module of the project (such
    as `GOOGLE_CLIENT_ID`, `GOOGLE_CLIENT_SECRET` etc.).

Disclaimer
----------

My personal goal with this project is to learn `Go` by doing. As such,
the project is very likely prone to bugs, and will be continuously
updated.

The final goal is to make a library, however in its current state the
project is far from that.

It also relies heavily on other publicly available resources:

-   Login page comes from [this
    pen](https://codepen.io/FlorinPop17/pen/vPKWjd);  
-   Frontend landing page is a small adjustment of [a well-known
    preloader](https://codepen.io/pawelqcm/pen/ObwyNe?limit=all&page=12&q=loader);
-   Standard login backend procedures follow the code & methods
    described
    [here](http://www.cihanozhan.com/building-login-and-register-application-with-golang/)
    and [here](https://www.sohamkamani.com/blog);  
-   Facebook and Google authentication follow the code from [this
    article on
    Medium](https://medium.com/@pliutau/getting-started-with-oauth2-in-go-2c9fae55d187);  
-   GitHub authentication relies heavily on [this implementation of
    GitHub authentication](https://github.com/godiscourse/godiscourse)
    in the context of online forum development.
