copyVocab(){

if ls ~/go/src/notion-english-csv-parser/*.zip 1> /dev/null 2>&1; then

 echo "extract zip file..."

 ~/go/src/notion-english-csv-parser/script.sh

fi



if ls ~/go/src/notion-english-csv-parser/*.csv 1> /dev/null 2>&1;then

echo "csv file found"

else

echo "csv file not found"

return

fi



go run ~/go/src/notion-english-csv-parser/main.go -t $1 -q $2 | pbcopy;





if [ $1 -eq 1 ]

then

 echo "quizlet: $2 words copied to clipboard"

else

 echo "vocabulary.com: $2 words copied to clipboard"

fi

}