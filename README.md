# yaBot (яБот)

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

yaBot (яБот) is a command-line tool written in Go for solving the v2 invisible reCaptcha only with requests.

## Table of Contents
- [Features](https://github.com/grozdniyandy/yaBot#features)
- [Usage](https://github.com/grozdniyandy/yaBot#usage)
- [Installation](https://github.com/grozdniyandy/yaBot#installation)
- [Input URL](https://github.com/grozdniyandy/yaBot#input-url)
- [Dependencies](https://github.com/grozdniyandy/yaBot#dependencies)
- [License](https://github.com/grozdniyandy/yaBot#license)
- [Author](https://github.com/grozdniyandy/yaBot#author)
- [Contributing](https://github.com/grozdniyandy/yaBot#contributing)

## Features
- reCaptcha v2 Inviisble solving

## Usage
1. **Clone or Download:** Clone this repository or download the code to your local machine.
2. **Run the tool:** Run the tool using the following command:
   ```
   go run main.go
   ```
3. **Give Anchor URL as input:**
   ```
   http://127.0.0.1:9090/cap?url={URL ENCODED URL HERE}
   ```
   
## Installation
You can either check the "Usage" and "Dependencies" or download already compiled code from "releases".

## Input URL
Input URL should be urlencoded
```
http%3A%2F%2F127.0.0.1%3A9090%2Fcap%3Furl%3Dhttps%3A%2F%2Fwww.google.com%2Frecaptcha%2Fapi2%2Fanchor%3Far%3D1%26k%3DKEY%26co%3DRANDOM.%26hl%3Den%26v%3DRANDOM%26size%3Dinvisible%26sa%3Dlogin%26cb%3DCB
```

## Dependencies
Install goquery by
```
go get github.com/PuerkitoBio/goquery
```

## License
This code is released under the [MIT License](LICENSE).

## Author
yaBot is developed by GrozdniyAndy of [XSS.is](https://xss.is).

## Contributing
Feel free to contribute, report issues, or suggest improvements by creating pull requests or issues in the GitHub repository. Enjoy using this simple captcha bypasser!
