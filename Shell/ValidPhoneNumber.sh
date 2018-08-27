# Given a text file file.txt that contains list of phone numbers (one per line), write a one liner bash script to print all valid phone numbers.
# 
# You may assume that a valid phone number must appear in one of the following two formats: (xxx) xxx-xxxx or xxx-xxx-xxxx. (x means a digit)
# 
# You may also assume each line in the text file must not contain leading or trailing white spaces.
# 
# For example, assume that file.txt has the following content:
# 
# 987-123-4567
# 123 456 7890
# (123) 456-7890
# Your script should output the following valid phone numbers:
# 987-123-4567
# (123) 456-7890

# get the input
rm -rf file.txt
echo "987-123-4567" >> file.txt
echo "123 456 7890" >> file.txt
echo "(123) 456-7890" >> file.txt
echo "input is:"
cat file.txt
echo ""

# output using awk
echo "output using awk:"
awk '/^([0-9][0-9][0-9]-|\([0-9][0-9][0-9]\) )[0-9][0-9][0-9]-[0-9][0-9][0-9][0-9]$/ {print $0}' file.txt
echo ""

# output using sed
echo "output using sed:"
sed -nr '/^([0-9]{3}-|\([0-9]{3}\) )[0-9]{3}-[0-9]{4}$/p' file.txt
echo ""

# output using grep
echo "output using grep:"
grep -P '^(\d{3}-|\(\d{3}\) )\d{3}-\d{4}$' file.txt
echo ""

# clean up
rm -rf file.txt


