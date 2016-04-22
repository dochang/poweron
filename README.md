Poweron
=======

Run it after installing Debian.

Requirements
------------

  - sudo
  - ca-certificates (Optional but recommended)
  - curl, wget, git or any other downloader.

Usage
-----

    wget -O - https://github.com/dochang/poweron/raw/master/go | sh -s -- -i <inventory> [ other ansible arguments ... ]

