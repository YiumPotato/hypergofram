
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

## Installation:

- Clone the repo!\
`git clone github.com/YiumPotato/HyperGoFram@latest`

- [Install TEMPL](https://templ.guide/quick-start/installation/)\
`go install github.com/a-h/templ/cmd/templ@latest`

- [Install Air](https://github.com/cosmtrek/air#installation) for live reloading of the app, it will rebuild the .TEMPL files on file change, perfect for go development quite good in development\
`go install github.com/cosmtrek/air@latest`

- run `air` and you're good to go, it will build & run automagically.

### TODO checklist

- [ ] Add Docker installation method
- [ ] Create CLI to add or delete certain things
- - [ ] Such as GCP Identifier Platform backend authentication
- - [ ] firebase auth
- [X] todos of todos

