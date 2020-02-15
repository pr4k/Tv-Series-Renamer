# Tv Series Renamer in Golang 
![build](https://api.travis-ci.org/pr4k/Tv-Series-Renamer.svg?branch=master)
![Issues](https://img.shields.io/github/issues/pr4k/Tv-Series-Renamer)
![Forks](https://img.shields.io/github/forks/pr4k/Tv-Series-Renamer)
![Stars](https://img.shields.io/github/stars/pr4k/Tv-Series-Renamer)
![license](https://img.shields.io/github/license/pr4k/Tv-Series-Renamer)

## About 
This is a TV Series Renamer which uses TVDB Api to Rename the episodes according to Aired Name and also Organizes the collection into folders.
It is based on go lang and was created in an attempt to learn GO

## Obtaining API Credentials
go to https://www.thetvdb.com/
Login and Obtain Username Apikey Userkey


Replace all of them in the in config.go 
or use `./script.sh $apiKey $username $userkey `
it will automatically replace the keys with your keys.

## How To Run

You can get it easily by using `go get pr4k/Tv-Series-Renamer`

Or simply Clone / Download the repo and build it using 

`go build`

An Executable file will be created, run that with argument as the path of the folder which contains the Tv- Series Episodes

Eg `./Tv-Series-Renamer -i /media/pr4k/New\ Volume/arrow/check`

Done

## Sample Run

### Before
![Before](images/before.png)

### After

![After](images/after1.png)

![After](images/after2.png)

![After](images/after3.png)

## Note 
As it is in initial stage be ready to get some bugs here and there.

## Wanna Give It a Try

Go to releases and try one for yourself :)

## Developer

### - *Prakhar Kaushik*
