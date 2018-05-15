# Shavronne
Simple discord bot for currency/uniques prices search, based on [poe.ninja](http://poe.ninja) api

Currently league locked by constant - need to implement league selection by user, but i'm too lazy 
# Building and launch
```bash
go get github.com/Mi7teR/shavronne
cd $GOPATH/src/github.com/Mi7teR/shavronne
dep ensure
go build
./shavronne -discordToken=abczexampleToken
```

