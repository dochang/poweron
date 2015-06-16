Poweron
=======

Run it after installing Debian.

Requirements
------------

  - sudo
  - ca-certificates (Optional but recommended)

Usage
-----

    wget -O - https://github.com/dochang/poweron/raw/master/go | sh

Or if a different apt source uri is needed:

    wget -O - https://github.com/dochang/poweron/raw/master/go | sh -s -- --extra-vars 'aptsource_uri=http://example.com/debian'

