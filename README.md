# Password Generator Demo

A demo on how to use <https://github.com/goavega-software/passwordgenerator> written as golang REST API and vueJS front-end. The demo is hosted on Google App Engine and uses Echo and Echo Templates for the web. It took me nearly couple of hours to figure out the error with 
```
gcloud app deploy
```
throwing build errors like "cannot use multiple packages with -o". Figured out later the issue was with demo project being outside of the $GOPATH src folder, moving it inside fixed the issue.
