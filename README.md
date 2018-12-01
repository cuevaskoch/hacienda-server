# hacienda-server

Hacienda is a suite of bespoke home automation software built for the Cuevas-Koch household.  Most of Hacienda's functionality is highly specialized to our needs and unlikely to be of use to anyone else, but I'm providing the source code online on the off-chance that you will find it useful.

## Features

* **Manage Pi-Hole:** I use [Pi-Hole](http://pi-hole.net/) to block Internet advertisements at a network level.  My wife uses advertising-supported referral programs to get cash back while doing online shopping.  Hacienda provides an easy way for her to temporarily disable ad blocking.

  She currently uses an iOS Shortcut to make an HTTP PUT request that will disable ad blocking for 5 minutes, but it's rather laborious to set up.  `hacienda-client` is coming soon to a repository near you.

## Installing

I have the server running in raspbian on a Raspberry Pi Zero W.  If you want to run it on your own raspi, you can run:

```bash
$ make build-debian-package
```

Which will create `./artifacts/hacienda.deb` for you for easy installation.
