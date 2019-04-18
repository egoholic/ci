# CI is a standalone continuous integration solution.

CI agent -> CI client -> CI server

## CI agent
Container level daemon that runs tests, and fitnesse functions. Collects build data and sends it to CI server if you need it.

## CI client
Host system level daemon and a CLI app that allows to build and manage Docker containes with CI daemon.

## CI server
Ecosystem level daemon and web app that shows build data, could automatically deploy.
