#curl https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json > json
#grep -A2 -P '"id": 70,' json > json2
#echo json2

##echo 
##grep 'name' json | echo [13:20] > json-bourne
###jq '.name' > json3
###cat json3


#simple solution only for the answer
curl -s https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq '.[52].name' 
