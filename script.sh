sed -i .bak '11s/.*/const apikey = \"'$1'\"/' config.go
sed -i .bak '12s/.*/const username = \"'$2'\"/' config.go
sed -i .bak '13s/.*/const userkey = \"'$3'\"/' config.go
