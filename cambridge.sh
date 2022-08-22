source ~/.zshenv
vocab
words=$(echo $(pbpaste)  | sed 's/,//g')
for word in $words[@]}; do
open -a "Google Chrome" https://dictionary.cambridge.org/dictionary/english/$word
done
