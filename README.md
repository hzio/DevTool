# DevTool
An all-round developer toolset.

## USAGE:
   devtool [global options] command [command options] [arguments...]

## VERSION:
   1.0

## DESCRIPTION:
   The toolset includes encryption, encoding, formatting, generation, searching, etc...

## COMMANDS:
  help, h  Shows a list of commands or help for one command
  
   - ENCODE:
     - **urlencode**  Urlencode the specified string
     - **urldecode**  Urldecode the specified string
   - ENCRYPT:
     - **base64**  Encrypt with base64 algorithm
     - **md5**     Encrypt with md5 algorithm
   - FETCH:
     - **ip**  Fetch the local IP address
   - FORMAT:
     - **timeformat, tf**  Format timestamp into date-time pattern
     - **jsonformat, jf**  Fortmat json string with indent
   - GENERATE:
     - **time, t, now**   Generate current date-time string
     - **date, d**        Generate current date string
     - **timestamp, ts**  Generate current timestamp
     - **passwd**         Generate a random password of 16 bits
     - **uuid**         Generate a UUID
   - SEARCH:
     - **baidu, b**         Search the specified keyword with Baidu search engine
     - **google, g**        Search the specified keyword with Google search engine
     - **baidumap, map**    Locate the specified place with Baidu map search engine
     - **googlemap, gmap**  Locate the specified place with Google map search engine
     - **wiki, wk**         Define the specified word with Wikipedia
   - SERVER:
     - **serve, s**  Start a static file server at port 9090

## GLOBAL OPTIONS:
   - --reverse, -r            Reverse current operation (default: false)
   - --help, -h               show help (default: false)
   - --init-completion value  generate completion code. Value must be 'bash' or 'zsh'
   - --version, -v            print the version (default: false)


## INSTALLATION
- macOS `brew tap mrhunterzhao/devtool && brew install devtool`