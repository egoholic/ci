# CI is a standalone continuous integration solution.

## Overview

CI agent -> CI client -> CI server

### CI agent
Container level daemon that runs tests, and fitnesse functions. Collects build data and sends it to CI server if you need it.

### CI client
Host system level daemon and a CLI app that allows to build and manage Docker containes with CI daemon.

### CI server
Ecosystem level daemon and web app that shows build data, could automatically deploy.


## Architecture
### Ubiquitous Language

**Pipeline** is an executing thing. Pipeline is a process of going throught all the stages and executing their commands.

**Pipeline Record** is a bunch of logs related to pipelines. Any pipeline state change logs into pipeline record.

**PipelineFactory** is a kinda factory pattern implementation that is responsible for taking *pipeline file*, creating a pipeline by its definitions and running it.

**Pipeline Records Repository** is a repository to deal with pipeline records.

**Build** is a product of a pipeline that succeed.

**Build Record** is a build-related information.

**Build Records Repository** is a repository to to deal with build records.

**Stage** is a bunch of commands that have the common purpose.

**Command** is a single instruction. Commands are grouped in stages by purpose.