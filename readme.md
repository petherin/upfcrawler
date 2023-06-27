# UPF Crawler

Uses github.com/tebeka/selenium and Chrome Driver to open online supermarket site and crawl for all product pages.

Depends on Chrome Driver which is included in the `bin` folder. If it fails to run on a Mac, `cd bin` and run `xattr -d com.apple.quarantine chromedriver`.

To install Chrome Driver onto your machine, get it from https://chromedriver.chromium.org/downloads

Move it to: `sudo mv chromedriver /usr/local/bin`

Allow Mac to run it by `cd`'ing to `/usr/local/bin` and running: `xattr -d com.apple.quarantine chromedriver`