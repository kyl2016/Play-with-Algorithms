# Given a text file file.txt, transpose its content.
# 
# You may assume that each row has the same number of columns and each field is separated by the ' ' character.
# 
# For example, if file.txt has the following content:
# 
# name age
# alice 21
# ryan 30
#
# Output the following:
# 
# name alice ryan
# age 21 30

# input is here
rm -rf file.txt
echo "name age location" >> file.txt
echo "alice 21 USA" >> file.txt
echo "ryan 30 CH"  >> file.txt
echo "input is:"
cat file.txt
echo ""

# transpose
echo "output is:"
awk '
    {
        for (i=1; i<=NF; i++) {
            if (line[i] == "") { line[i] = $i } 
            else { line[i] = line[i]" "$i }
        }
    } END {
        for (i=1; i<=NF; i++) { print line[i] }
    }
' file.txt

# clean up
rm -rf file.txt
