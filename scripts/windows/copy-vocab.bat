:copy-vocab
@ECHO OFF

Set "Type=%~1"
Set "Quantity=%~2"

if exist C:\\Users\\ozcol\\Desktop\\repos\\vocabulary-csv-parser\\*.zip (
 echo "extract zip file..." C:\\Users\\ozcol\\Desktop\\repos\\vocabulary-csv-parser\\scripts\\windows\\generate-csv.sh
)


if exist C:\\Users\\ozcol\\Desktop\\repos\\vocabulary-csv-parser\\*.csv (
echo "csv file found"
) else (
  echo "csv file not found"
  goto :eof
)

go run C:\\Users\\ozcol\\Desktop\\repos\\vocabulary-csv-parser\\main.go -t %Type% -q %Quantity% | clip

if %Type%==1 (
echo quizlet: %Quantity% words copied to clipboard
) else (
 echo vocabulary.com: %Quantity% words copied to clipboard
)
