count="$1"

[ "$count" -gt 100 ] && count=100

for ((repetitions = 1; repetitions <= count; repetitions++)); do
echo "$repetitions times I've printed rasmusjoenurm"
done