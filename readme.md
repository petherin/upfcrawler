# UPF Crawler

Use Selenium and Chrome Driver to open site and click links to get HTML from JavaScript.

Get Selenium: `go get github.com/tebeka/selenium`

Install Chrome Driver from https://chromedriver.chromium.org/downloads

Move it to: `sudo mv chromedriver /usr/local/bin`

Allow Mac to run it with: `xattr -d com.apple.quarantine chromedriver`

Should get this working in Docker.

Need to click each aisle link to see a list of products, click on each one, get ingredients.