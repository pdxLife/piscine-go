#ls -Atr | tr '\n' ','

ls -lcrt --format=comma -p | tr -d ' '


#ls -c -m --indicator-style=comma | tr -d ' '