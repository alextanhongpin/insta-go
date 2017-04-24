# Instago

A modern Instagram created with Golang as the backend and Elm as the frontend.

The goal of the project is to understand how productive can it be to code an app using Golang. I simply choose Instagram for no particular reason. 

For the database, I'll be using Postgres. 


## Architecture

### What is the best way to store configs? 

According to the 12-Factor-App, configs should be stored in the environment variables. How does it apply in Golang?

### What is the best practices to use Postgres with Golang?

I'm still exploring this particular area. The current way is to just store a global sql context.

There are some issues that I encountered when working datetime conversion and postgres array in Golang.

### Designing models

C'mon, models should be fat. How can I make my models more useful?

### Deployment

Docker?

## Current progress

Watch me throw some css magic dust on this...
![The start!](assets/01-the-beginning.jpeg)