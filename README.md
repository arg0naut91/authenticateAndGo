Authenticate and go!
====================

<img src="present/authGIF.gif" width="600" height="400" />

Purpose
-------

The aim of this project is to work on web authentication module with all
the major components integrated:

-   Register and login through the database (currently this is
    `PostgreSQL`);  
-   Register/Login through `OAuth2` (e.g. social networks - currently
    Facebook, Google and GitHub are supported);  
-   Persistent login (currently this works for the standard login);  
-   Password recovery (not yet implemented).

Disclaimer
----------

My personal goal with this project is to learn `Go` by doing. As such,
the project is very likely prone to bugs, and will be continuously
updated.

The final goal is to make a library, however in its current state the
project is far from that.

It also relies heavily on the other publicly available resources:

-   Login page comes from [this nice
    pen](https://codepen.io/FlorinPop17/pen/vPKWjd);  
-   Frontend landing page is a small adjustment of [a well-known
    preloader](https://codepen.io/pawelqcm/pen/ObwyNe?limit=all&page=12&q=loader);
-   Standard login backend procedures follow the code & methods nicely
    described
    [here](http://www.cihanozhan.com/building-login-and-register-application-with-golang/)
    and [here](https://www.sohamkamani.com/blog);  
-   Facebook and Google authentication follow the code from [this
    article on
    Medium](https://medium.com/@pliutau/getting-started-with-oauth2-in-go-2c9fae55d187);  
-   GitHub authentication relies heavily on [this great implementation
    of a forum](https://github.com/godiscourse/godiscourse).
