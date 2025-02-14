echo "file named a" > a
echo "file named !" > '!'
echo "file named \\" > '\'
echo "file named \"" > '"'

mkdir '`'

cat '!' > '`/!'

if [ "$MOVE_A" == "yes" ]; then
    mv a '`/'
elif [ "$MOVE_A" == "no" ]; then
    rm a
fi