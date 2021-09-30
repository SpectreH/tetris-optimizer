printf "Bad example 1: \n"
go run . error-tests/example00.txt
sleep 3
printf "\nBad example 2: \n"
go run . error-tests/example01.txt
sleep 3
printf "\nBad example 3: \n"
go run . error-tests/example02.txt
sleep 3
printf "\nBad example 4: \n"
go run . error-tests/example03.txt
sleep 3
printf "\nBad example 5: \n"
go run . error-tests/example04.txt
printf "\nGood Example 1: \n"
go run . tests/sample00.txt
sleep 3
printf "\nGood Example 2: \n"
go run . tests/sample01.txt
sleep 3
printf "\nGood Example 3: \n"
go run . tests/sample02.txt
sleep 3
printf "\nGood Example 4: \n"
go run . tests/sample03.txt
sleep 3
printf "\nGood Example 5 (need to wait longer, because hard example): \n"
go run . tests/sample04.txt