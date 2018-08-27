# How would you print just the 10th line of a file?
# 
# For example, assume that file.txt has the following content:
# 
# Line 1
# Line 2
# Line 3
# Line 4
# Line 5
# Line 6
# Line 7
# Line 8
# Line 9
# Line 10
# 
# Your script should output the tenth line, which is:
# Line 10

# input 
# INPUT=$'line 1\nline 2\nline 3\nline 4\nline 5\nline 6\nline 7\nline 8\nline 9\nline 10'
INPUT=$'line 1\nline 2\nline 3\nline 4\nline 5\nline 6\nline 7\nline 8\nline 9\nline 10'
echo "$INPUT" > file.txt
echo "input is:"
cat file.txt
echo ""

# output using sed
echo "output using sed: "
sed -n "10p" file.txt
echo ""

# output with awk
echo "output using awk: "
awk 'NR == 10 {print; exit}' file.txt
echo ""

# output using wc and tail
echo "output using wc and tail:"
if [[ $(wc -l < file.txt) -ge 10 ]]
then
    head -n 10 file.txt | tail -n 1
fi

# now clean up
rm -rf file.txt










