
# HyperGoFram
> fullstack golang development experience

Develop SSR golang apps quickly with this minimalist yet powerful framework!\
No more huge bloating JS for your websites. Bonus, uses HTMX's battery included Web Transitions API for proper web SPA sex appeal. 

## Features
- TEMPL as the templating engine
- Gin as the Web Framework
- HTMX (HTML extended with AJAX) inside TEMPL components
- opinionated folder structure
- instant reload with Air (build and run)
- pre instant SQLC generation of db queries into go code
- Go Migrate to generate db schema
- simple Makefile to debug / configure DB

## Installation:

### Docker installation

You can either use `docker compose up`, which will boot up 2 containers:
- air_dev: image made from the Dockerfile consisting of hot reloading Air with TEMPL and sqlc
- postgres: psql running on env settings set in the compose file

To deal with db migrations that way I recommend to start the postgres container first, and use go migrate to migrate the db. (`docker compose start postgres`)


### Manually

- Clone the repo!\
`git clone github.com/YiumPotato/HyperGoFram@latest`

- [Install TEMPL](https://templ.guide/quick-start/installation/)\
`go install github.com/a-h/templ/cmd/templ@latest`

- [Install Air](https://github.com/cosmtrek/air#installation) for live reloading of the app, it will rebuild the .TEMPL files on file change, perfect for go development quite good in development\
`go install github.com/cosmtrek/air@latest`

- run `air` and you're good to go, it will build & run automagically.

### Makefile
Use the Makefile to: (`make createdb` for example)
- dropdb 
- createdb
- migrateup
- migratedown

## TODO checklist

- [X] Add Docker installation method
- [X] Added DB configuration
- [ ] Go migrate db 
- [ ] Create CLI to add or delete certain things
- - [ ] Such as GCP Identifier Platform backend authentication
- - [ ] firebase auth
- [X] todos of todos

