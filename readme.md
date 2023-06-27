# UPF Crawler

Uses github.com/tebeka/selenium and Chrome Driver to open online supermarket site and crawl for all product pages.

Written in Go, requires Go installation.

Run with `go run cmd/cli/main.go`.

Depends on Chrome Driver which is included in the `bin` folder. If it fails to run on a Mac, `cd bin` and run `xattr -d com.apple.quarantine chromedriver`.

To install Chrome Driver onto your machine, get it from https://chromedriver.chromium.org/downloads

Move it to: `sudo mv chromedriver /usr/local/bin`

Allow Mac to run it by `cd`'ing to `/usr/local/bin` and running: `xattr -d com.apple.quarantine chromedriver`.

## TODO

* Run in Docker.
* Once product pages gathered, save to file and re-use so we can skip the crawl stage next time we run.
* Refine product pages visited.
* Classify foods as UPF or non-UPF.
* Generate file like a CSV to save food classifications.
* Do more stores.
* Add unit tests.
* Add Makefile.