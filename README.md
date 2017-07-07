## Follow shorten links
Small CLI utility written in Go to follow the redirects of shorten links


#### Command-line usage

|Flag|Default|Description|
|----|-------|-----------|
|u|```Empty string```|Shorten URL to follow|
|f|```false```|Show all redirects in output, otherwise only the last URL is printed|
|r|```true```|Remove query strings from every URL|


#### IMPORTANT

- This DOESN'T provide any form of privacy, there are still requests made that can be logged.
- It does give some sort of security due to the fact that the actual content is not loaded and executed


#### License
GNU General Public License v3.0
