numberOfFiles=$(find . \( ! -regex '.*/\..*' \) -type f,d | wc -l)

multipleFiles=$((numberOfFiles * 5))

printf "\t\vTotal files * 5: %s\v\n" $multipleFiles