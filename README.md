availabilitybot
===============

[![](https://img.shields.io/docker/automated/atsnngs/availabilitybot.svg)](https://hub.docker.com/r/atsnngs/availabilitybot/)
[![Build Status](https://travis-ci.org/ngs/availabilitybot.svg?branch=master)](https://travis-ci.org/ngs/availabilitybot)
[![Coverage Status](https://coveralls.io/repos/github/ngs/availabilitybot/badge.svg?branch=master)](https://coveralls.io/github/ngs/availabilitybot?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/ngs/availabilitybot)](https://goreportcard.com/report/github.com/ngs/availabilitybot)

Tweets Parts Availability in Apple Stores in Japan.

Environment variables
---------------------

Grab your Twitter credentials from https://apps.twitter.com/

```sh
TWITTER_CONSUMER_KEY
TWITTER_CONSUMER_SECRET
TWITTER_ACCESS_TOKEN
TWITTER_ACCESS_SECRET
```

Quickstart
----------

```sh
docker pull atsnngs/availabilitybot
docker run --env-file .env \
  --rm -v $(pwd)/bot:/home/bot \
  -it atsnngs/availabilitybot \
  MMEF2J/A
```

Crontab
-------

```sh
* * * * * docker run --rm --env-file /home/ngs/availabilitybot/.env -v /home/ngs/availabilitybot:/home/bot -it atsnngs/availabilitybot MMEF2J/A
```

Build
-----

```sh
go get -v -t -d .
go build .
./availabilitybot MMEF2J/A
```
