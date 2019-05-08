# Tv Series Renamer in Golang 

## About 
This is a TV Series Renamer which uses TVDB Api to Rename the episodes according to Aired Name and also Organizes the collection into folders.
It is based on go lang and was created in an attempt to learn GO

## Obtaining Api Credentials
go to https://www.thetvdb.com/
Login and Obtain Username Apikey Userkey

`{"apikey":"Your_Api_key","username":"Your_Username","userkey":"Your_User_Key"}`

Replace all of them in the json string above mentioned and replace this line in token.go with your credentials

## How To Run

Clone or Download the repo and build it using 

`go build`

An Executable file will be created, run that with argument as the path of the folder which contains the Tv- Series Episodes

Eg ./tvdb /media/pr4k/New\ Volume/arrow/check

Done

## Sample Run

### Before
![Before](images/before.png)

### After

![After](images/after1.png)

![After](images/after2.png)

![After](images/after3.png)

## Note 
As it is in initial stage be ready to get too many things printed on screen and also few episodes left to rename but still it works.

## Developer

### - *Prakhar Kaushik*
